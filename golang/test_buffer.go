
package main

import (
    "fmt"
    "bytes"
)


func main() {

    byts := []byte("admin.$cmd\x00")

    fmt.Printf("%v",byts)

    if bytes.HasPrefix(byts,[]byte("admin.$cmd")) {
        fmt.Println("yes of cause")
    }
    buf := bytes.NewBuffer(byts)

    buf.Next(4)
    fmt.Printf("%v\n",buf.Len())

    byts = []byte("")

    fmt.Printf("%v\n",byts)

    buf = bytes.NewBuffer([]byte("admin.$cmd\x00"))

    aaa,_ := buf.ReadBytes('\x00')

    fmt.Printf("%#v\n",aaa)

}
