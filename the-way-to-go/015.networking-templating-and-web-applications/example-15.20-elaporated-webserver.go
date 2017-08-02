package main

import (
    "bytes"
    "expvar"
    "flag"
    "fmt"
    "net/http"
    "io"
    "log"
    "os"
    "strconv"
    "syscall"
)

/*
route mapping:
/
    log request path
        curl http://localhost:8088/abcdefg
/go/hello
    print "Hello World!" && publish hello-requests count
        curl http://localhost:8088/go/hello
        curl http://localhost:8088/debug/vars | jq . | grep hello-requests
/counter
    view publish counter
        curl http://localhost:8088/debug/vars | jq . | grep counter
    get current counter
        curl http://localhost:8088/counter
    reset counter to 100
        curl -d100 http://localhost:8088/counter
/go/<path>
    fileserver: retrieve file content
        curl http://localhost:8088/go/proxy.php // notice: here is text content
/flags
    retrieve flags used when server startup
        curl http://localhost:8088/flags
/args
    retrieve command args when server startup
        curl http://localhost:8088/args
/chan
    get counter from a chan and print it
        curl http://localhost:8088/chan
/date
    get current date and time
        curl http://localhost:8088/date
 */

// hello world, the web server
var helloRequests = expvar.NewInt("hello-requests")

// flags:
var webroot = flag.String("root", "/home/work/www/www.tec-inf.com", "web root directory")
// simple flag server
var booleanflag = flag.Bool("boolean", true, "another flag for testing")

// Simple counter server. POSTING to it will set the value
type Counter struct {
    n int
}

// a channel
type Chan chan int

func main() {
    flag.Parse()

    http.Handle("/", http.HandlerFunc(Logger))
    http.Handle("/go/hello", http.HandlerFunc(HelloServer))
    // counter is published as a variable director
    ctr := new(Counter)
    expvar.Publish("counter", ctr)
    http.Handle("/counter", ctr)
    http.Handle("/go/", http.StripPrefix("/go/", http.FileServer(http.Dir(*webroot))))
    http.Handle("/flags", http.HandlerFunc(FlagServer))
    http.Handle("/args", http.HandlerFunc(ArgServer))
    http.Handle("/chan", ChanCreate())
    http.Handle("/date", http.HandlerFunc(DateServer))
    err := http.ListenAndServe(":8088", nil)
    if err != nil {
        log.Panicln("ListenAndServe:", err)
    }
}

func Logger(w http.ResponseWriter, r *http.Request) {
    log.Print(r.URL.String())
    w.WriteHeader(http.StatusNotFound)
    w.Write([]byte("oops"))
}

func FlagServer(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
    fmt.Fprintf(w, "Flags:\n")
    flag.VisitAll(func(f *flag.Flag) {
        if f.Value.String() != f.DefValue {
            fmt.Fprintf(w, "%s = %s [default = %s]\n", f.Name, f.Value.String(), f.DefValue)
        } else {
            fmt.Fprintf(w, "%s = %s\n", f.Name, f.Value.String())
        }
    })
}

func ArgServer(w http.ResponseWriter, r *http.Request) {
    for i, s := range os.Args {
        fmt.Fprintf(w, "Args[%d]: %s\n", i, s)
    }
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    helloRequests.Add(1)
    io.WriteString(w, "Hello World!\n")
}

func (c *Counter) String() string {
    return fmt.Sprintf("%d", c.n)
}

func (c *Counter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
        case http.MethodGet:
            c.n ++
        case http.MethodPost:
            buf := new(bytes.Buffer)
            io.Copy(buf, r.Body)
            body := buf.String()
            if n, err := strconv.Atoi(body); err != nil {
                fmt.Fprintf(w, "bad POST: %v\nbody: [%v]\n", err, body)
            } else {
                c.n = n
                fmt.Fprintf(w, "Counter reset\n")
            }
    }
    fmt.Fprintf(w, "counter = %d\n", c.n)
}

func ChanCreate() Chan {
    c := make(Chan)
    go func(c Chan) {
        for x := 0; ; x ++ {
            c <- x
        }
    }(c)
    return c
}

func (ch Chan) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, fmt.Sprintf("channel send #%d\n", <-ch))
}

func DateServer(rw http.ResponseWriter, req *http.Request) {
    rw.Header().Set("Content-Type", "text/plain; charset=UTF-8")
    r, w, err := os.Pipe()
    if err != nil {
        fmt.Fprintf(rw, "pipe: %s\n", err)
        return
    }

    p, err := os.StartProcess("/bin/date", []string{"date"}, &os.ProcAttr{Files: []*os.File{nil, w, w}})
    defer r.Close()
    w.Close()
    if err != nil {
        fmt.Fprintf(rw, "fork/exec: %s\n", err)
        return
    }
    defer p.Release()
    io.Copy(rw, r)
    wait, err := p.Wait()
    if err != nil {
        fmt.Fprintf(rw, "wait error: %s\n", err)
        return
    }
    if !wait.Exited() || wait.Sys().(syscall.WaitStatus) != 0 {
        fmt.Fprintf(rw, "wait status error: %v\n", wait)
        return
    }
}