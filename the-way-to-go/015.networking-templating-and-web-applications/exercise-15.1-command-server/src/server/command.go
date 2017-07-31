package server

import (
    "fmt"
)

type Command struct {
    client *Client
    name string
    arg string
    handler func(c *Command)
}

const (
    CMD_UNKNOWN = "unknown"
    CMD_ECHO = "echo"
    CMD_LIST = "list"
    CMD_EXIT = "exit"
    CMD_SHUTDOWN = "shutdown"
)

var commandHandlers map[string]func(c *Command)

func init() {
    commandHandlers = map[string]func(c *Command) {
        CMD_UNKNOWN: command_handler_unknown,
        CMD_ECHO: command_handler_echo,
        CMD_LIST: command_handler_list,
        CMD_EXIT: command_handler_exit,
        CMD_SHUTDOWN: command_handler_shutdown,
    }
}

func NewCommand(client *Client, name, arg string) *Command {
    handler, exists := commandHandlers[name]
    if exists {
        return &Command{client: client, name: name, arg: arg, handler: handler}
    }
    return nil
}

func (c *Command) Run() {
    c.handler(c)
}

func command_handler_echo(c *Command) {
    c.client.Write(c.arg)
}

func command_handler_unknown(c *Command) {
    fmt.Printf("unknown message: %s\n", c.arg)
}

func command_handler_list(c *Command) {
    fmt.Printf("client list:\n")
    for _, client := range c.client.svr.clients {
        fmt.Printf("\t%s:%d\n", client.hostname, client.port)
    }
}

func command_handler_exit(c *Command) {
    c.client.Close("client exit")
}

func command_handler_shutdown(c *Command) {
    c.client.svr.Shutdown("client shutdown")
}
