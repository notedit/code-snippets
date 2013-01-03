
package main

import (
    "bytes"
    "fmt"
)

func main() {
    
    byt := []byte("dfdfdfdfdfdfdfdfdfd")

    buf := bytes.NewBuffer(byt)

    newb := make([]byte, 3)

    a,err := buf.Read(newb)

    b,err := buf.Read(newb)

    fmt.Println(a,b,err)
}
