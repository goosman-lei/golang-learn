package main
/*
/test1
    * handled by SimpleServer
    * output hello world in the browser
/test2
    * handled by FormServer
    * if request method is GET
        * response form, which contains the html for a simple input form with text box and submit button
    * if request method is POST
        * retrieve content of the textbox which name is inp, and written it to the browser page
 */

import (
    "net/http"
    "io"
    "fmt"
)

const (
    html_form = "<html><head><title>Form Demo</title></head><body><form method=\"POST\" action=\"\"><input type=\"text\" name=\"inp\"/><br /><input type=\"submit\"/></form></body></html>"
    html_helloworld = "<html><head><title>HelloWorld</title></head><body><h1>Hello World</h1></body></html>"
    html_fmt_form_resp = "<html><head><title>HelloWorld</title></head><body><h1>%s</h1></body></html>"
)

func main() {
    http.HandleFunc("/test1", SimpleServer)
    http.HandleFunc("/test2", FormServer)
    err := http.ListenAndServe("0.0.0.0:8088", nil)
    if err != nil {
        fmt.Printf("web-server startup failed: %s\n", err.Error())
    }
}

func SimpleServer(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    io.WriteString(w, html_helloworld)
}

func FormServer(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    switch req.Method {
        case http.MethodGet:
            io.WriteString(w, html_form)
        case http.MethodPost:
            _, err := io.WriteString(w, fmt.Sprintf(html_fmt_form_resp, req.FormValue("inp")))
            if err != nil {
                fmt.Printf("Write response content failed: %s\n", err.Error())
            }
        default:
            w.WriteHeader(http.StatusMethodNotAllowed)
    }
}