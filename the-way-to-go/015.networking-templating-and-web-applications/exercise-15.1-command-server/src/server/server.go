package server

import (
    "net"
    "fmt"
)

type Server struct {
    hostname string
    port int
    status int
    listener net.Listener
    waitCh chan string      // main goroutine wait exit msg of server
    clients []*Client
}

const (
    SVR_SHUTDOWN = iota
    SVR_STARTING
    SVR_RUNNING
    SVR_SHUTDOWNING
)

const DEFAULT_CLIENT_BUFFER_SIZE = 128

func NewServer(hostname string, port int) *Server {
    return &Server{
        hostname: hostname,
        port: port,
        status: SVR_SHUTDOWN,
        waitCh: make(chan string),
        clients: make([]*Client, 0, DEFAULT_CLIENT_BUFFER_SIZE),
    }
}
func (s *Server) CanStartup() bool {
    return s.status == SVR_SHUTDOWN
}
func (s *Server) Startup() error {
    if !s.CanStartup() {
        return fmt.Errorf("server can not startup, its status now is: %s", s.status)
    }
    // listening
    s.status = SVR_STARTING
    listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.hostname, s.port))
    if err != nil {
        return fmt.Errorf("Server.Startup() failed at Listen[%s:%d]: %s\n", s.hostname, s.port, err.Error())
    }
    s.listener = listener

    // accept client connect request
    go s.run()
    s.status = SVR_RUNNING
    fmt.Printf("Server running success at: %s:%d\n", s.hostname, s.port)

    return s.wait()
}

func (s *Server) wait() (err error) {
    // wait exit msg
    if msg, ok := <-s.waitCh; ok {
        if msg == "ok" {
            return
        }
        return fmt.Errorf("%s", msg)
    }
    return fmt.Errorf("wait channel close unexpected")
}

func (s *Server) CanShutdown() bool {
    return s.status != SVR_SHUTDOWNING && s.status != SVR_SHUTDOWN
}
func (s *Server) Shutdown(reason string) {
    if !s.CanShutdown() {
        return
    }
    s.status = SVR_SHUTDOWNING

    clients := make([]*Client, len(s.clients))
    copy(clients, s.clients)
    for _, client := range clients {
        client.Close("server shutdown")
    }
    // maybe, here need a sync, wait all client success closed
    s.status = SVR_SHUTDOWN
    s.waitCh <- fmt.Sprintf("[shutdown: %s]", reason)
}
func (s *Server) Fail(fmtStr string, args ...interface{}) {
    s.waitCh <- fmt.Sprintf("[fail: " + fmtStr + "]", args...)
    s.status = SVR_SHUTDOWN
}

func (s *Server) run() {
    for {
        conn, err := s.listener.Accept()
        if err != nil {
            s.waitCh <- fmt.Sprintf("Server Accept Failed: %s", err.Error())
            break
        }

        go s.NewClientAndRun(conn)
    }
}

func (s *Server) pushClient(c *Client) {
    s.clients = append(s.clients, c)
}
func (s *Server) RemoveClient(c *Client) {
    for i, tc := range s.clients {
        if tc == c {
            s.clients = append(s.clients[:i], s.clients[i+1:]...)
            break
        }
    }
}

func (s *Server) NewClientAndRun(conn net.Conn) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Client panic: %s\n", r)
        }
    }()
    c, err := NewClient(s, conn)
    if err != nil {
        panic(err)
    }
    s.pushClient(c)
    c.Run()
}
