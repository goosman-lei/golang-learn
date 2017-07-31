package server

import (
    "net"
    "regexp"
    "fmt"
    "bufio"
    "strings"
    "strconv"
)

type Client struct {
    hostname string
    port int
    conn net.Conn
    reader *bufio.Reader
    writer *bufio.Writer
    svr *Server
    closed bool
}

var reParseCommand, reParseAddr *regexp.Regexp

func init() {
    re, err := regexp.Compile("^([\\w-]+)(?::\\s*(.*))?$")
    if err != nil {
        panic(err.Error())
    }
    reParseCommand = re

    re, err = regexp.Compile("^(\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}):(\\d{1,5})$")
    if err != nil {
        panic(err.Error())
    }
    reParseAddr = re
}

func NewClient(s *Server, conn net.Conn) (c *Client, err error) {
    matches := reParseAddr.FindStringSubmatch(conn.RemoteAddr().String())
    if len(matches) != 3 {
        err = fmt.Errorf("address format failed: %s\n", conn.RemoteAddr().String())
        return
    }
    port, err := strconv.ParseInt(matches[2], 10, 64)
    if err != nil {
        err = fmt.Errorf("address format failed: %s\n", conn.RemoteAddr().String())
        return
    }

    c = &Client{
        conn: conn,
        svr: s,
        hostname: matches[1],
        port: int(port),
        closed: false,
    }

    c.reader = bufio.NewReader(conn)
    c.writer = bufio.NewWriter(conn)

    return
}

func (c *Client) Run() {
    for !c.closed {
        command, err := c.ReadCommand()
        if err != nil {
            c.Close(err.Error())
            break
        }
        command.Run()
    }
}

func (c *Client) Close(message string) {
    if c.closed {
        return
    }
    c.closed = true
    c.conn.Close()
    c.svr.RemoveClient(c)
    fmt.Printf("Client[%s:%d] closed: %s\n", c.hostname, c.port, message)
}

func (c *Client) ReadCommand() (command *Command, err error){
    input, err := c.reader.ReadString('\n')
    if err != nil {
        if c.closed {
            err = fmt.Errorf("client closed")
        } else {
            err = fmt.Errorf("read input failed: %s\n", err.Error())
        }
        return
    }
    input = strings.TrimRight(input, "\r\n")

    //fmt.Printf("[%s:%d] message: %s\n", c.hostname, c.port, input)

    matches := reParseCommand.FindStringSubmatch(input)
    if len(matches) != 3 {
        err = fmt.Errorf("input format failed: %s\n", input)
        return
    }

    command = NewCommand(c, matches[1], matches[2])
    if command == nil {
        command = NewCommand(c, CMD_UNKNOWN, matches[0])
    }

    return
}

func (c *Client) Write(resp string) (nWrite int, err error) {
    for eWrite := len(resp); nWrite < eWrite; {
        n, e := c.writer.WriteString(resp[nWrite:eWrite])
        if e != nil {
            err = e
            return
        }
        nWrite += n
    }
    c.writer.WriteString("\n")
    c.writer.Flush()
    return
}
