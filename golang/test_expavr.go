package main

import (
    "expvar"
    "net/http"
)

var exportedValue = expvar.NewString("hello")

func main(){
    exportedValue.Set("HELLO")
    http.ListenAndServe(":9090",nil)
}
