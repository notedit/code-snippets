package main

import (
    "crypto/sha1"
    "fmt"
    "io"
)

func main(){
    h :=  sha1.New()
    io.WriteString(h,"acftxxtmtu"+"pass05")
    sha := h.Sum(nil)
    fmt.Printf("%#x\n",string(sha))
}
