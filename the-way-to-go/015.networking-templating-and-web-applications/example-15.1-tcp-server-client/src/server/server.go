package server

import (
    "net"
    "fmt"
    "log"
    "os"
    "io"
)

type Command struct {
    keyword string
    handler func(net.Conn)
}

type Server struct {
    status int
    listener net.Listener
    clients []net.Conn
}

const (
    SERVER_STARTING = iota
    SERVER_RUNING
    SERVER_SHUTDOWNING
    SERVER_SHUTDOWN
)

var logger *log.Logger

var commands []Command

var svr *Server

func init() {
    svr = new(Server)
    logger = log.New(os.Stderr, "", log.LstdFlags)
    commands = []Command{
        Command{"who", server_handler_who},
        Command{"sh", server_handler_sh},
    }
}

func (s *Server) Shutdown() {
    s.status = SERVER_SHUTDOWNING
    for _, clientPtr := range svr.clients {
        clientPtr.Close()
    }
    s.listener.Close()
    s.status = SERVER_SHUTDOWN
}

func StartServer(ip string, port int) {
    svr.status = SERVER_STARTING
    svr.clients = make([]net.Conn, 0, 128)

    server, err := net.Listen("tcp4", fmt.Sprintf("%s:%d", ip, port))
    if err != nil {
        logger.Printf("server listen at %s failed: %s", server.Addr(), err.Error())
        return
    }
    defer func() {
        logger.Printf("server[%s] shutdown", server.Addr())
        server.Close()
    }()
    logger.Printf("server startup at %s", server.Addr())
    svr.listener = server
    svr.status = SERVER_RUNING


    for {
        client, err := server.Accept()
        if err != nil {
            if err.Error() == "use of closed network connection" {
                logger.Printf("Server Shutdown With client command\n")
                break
            }
            logger.Printf("server[%s] accept new client connect failed: %s", server.Addr(), err.Error())
            break
        }
        go clientHandler(client)
    }
    return
}

func clientHandler(client net.Conn) {
    svr.clients = append(svr.clients, client)
    remoteAddr := client.RemoteAddr()
    defer func() {
        if r := recover(); r != nil {
            logger.Printf("client[%s] panic: %s\n", remoteAddr, r)
        }
        logger.Printf("close connect of client[%s]\n", remoteAddr)
        for idx, c := range svr.clients {
            if c == client {
                svr.clients = append(svr.clients[:idx], svr.clients[idx+1:]...)
            }
        }
        client.Close()
    }()
    logger.Printf("new client from %s\n", remoteAddr)

    buffer := make([]byte, 512)
    LOOP_NEW_MSG:
    for {
        nRead, err := client.Read(buffer)
        if err == io.EOF {
            logger.Printf("client[%s] closed connect\n", remoteAddr)
            break
        } else if err != nil {
            panic(err.Error())
        }

        for _, command := range commands {
            if string(buffer[:nRead]) == "guoguo: " + command.keyword {
                command.handler(client)
                continue LOOP_NEW_MSG
            }
        }

        logger.Printf("Receive %d bytes from[%s]: %s\n", nRead, client.RemoteAddr(), buffer[:nRead])
    }
}

func server_handler_who(client net.Conn) {
    logger.Printf("[%s] ask client list:\n", client.RemoteAddr())
    for _, c := range svr.clients {
        logger.Printf("%s\n", c.RemoteAddr())
    }
}

func server_handler_sh(client net.Conn) {
    svr.Shutdown()
}