
package main

import (
    "fmt"
    "net"
)

func main() {

    var timeout int64 
    timeout = 2000000000
    addr,err := net.ResolveTCPAddr("tcp","localhost:10000")
    listener,err := net.ListenTCP("tcp",addr)
    if err != nil {
        fmt.Println(err)
    }
    listener.SetTimeout(timeout)
    fmt.Println(timeout)

    for {
        fmt.Println("enter for loop")
        conn,err := listener.AcceptTCP()
        if err != nil {
            fmt.Sprintf("%#v\n",conn)
        }
        fmt.Println("dfdf\n")
        fmt.Sprintf("%#v%#v\n",conn,err)
    }
}
