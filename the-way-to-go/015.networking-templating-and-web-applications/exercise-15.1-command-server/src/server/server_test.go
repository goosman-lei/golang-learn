package server

import (
    "testing"
)

func TestNewServer(t *testing.T) {
    host := "localhost"
    port := 8088
    s := NewServer(host, port)
    if s.hostname != "localhost" || s.port != 8088 {
        t.Logf("invalid [host:port]: [%s:%d] wanted: [%s:%d]", s.hostname, s.port, host, port)
        t.Fail()
    }
}
