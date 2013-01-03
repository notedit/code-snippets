package main

import (
    "io"
    "net/http"
    "strconv"
)

func HelloServer(c http.ResponseWriter, req *http.Request) {
    body := "<b>hello world</b>"
    c.Header().Add("Content-Type", "text/html")
    c.Header().Add("Content-Length", strconv.Itoa(len(body)))
    io.WriteString(c, body)
}

func main(){
    http.Handle("/",http.HandlerFunc(HelloServer))
    http.ListenAndServe(":9090",nil)
}
