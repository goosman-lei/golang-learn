package main

import (
    "server"
    "os"
    "os/signal"
    "fmt"
)

var s *server.Server

func main() {
    sigSetup()

    s = server.NewServer("0.0.0.0", 8088)
    err := s.Startup()
    if err != nil {
        fmt.Printf("Server shutdown: %s\n", err.Error())
    } else {
        fmt.Printf("Server shutdown normal\n")
    }
}

func sigSetup() {
    sigCh := make(chan os.Signal, 1)
    signal.Notify(sigCh, os.Interrupt)

    go func() {
        for {
            if sig, ok := <-sigCh; ok {
                if !sigHandler(sig) {
                    break
                }
            }
        }
    }()
}
func sigHandler(sig os.Signal) bool {
    switch sig {
        case os.Interrupt:
            if s != nil && s.CanShutdown() {
                s.Shutdown("SIG_INTERRUPT")
            }
        default:
            fmt.Printf("uncatch signal: %s\n", sig.String())
    }
    return true
}
