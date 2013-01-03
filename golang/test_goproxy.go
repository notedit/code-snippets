package main

import (
    "net"
    "fmt"
    "os"
)

func main() {
    if len(os.Args) != 3 {
        fmt.Println("usage: netfwd local remote")
        os.Exit(2)
    }
    localAddr := os.Args[1]
    remoteAddr := os.Args[2]

    local,err := net.Listen("tcp",localAddr)
    if err != nil {
        fmt.Println("listen error:",err)
        os.Exit(2)
    }
    for {
        conn,err := local.Accept()
        if err != nil {
            fmt.Println("accept error:",err)
            os.Exit(2)
        }
        go forward(conn,remoteAddr)
    }
}

func forward(local net.Conn,remoteAddr string) {
    remote,err := net.Dial("tcp",remoteAddr)
    if err != nil {
        fmt.Fprintf(os.Stderr,"remote dial failed: %v\n",err)
        return
    }
    //go io.Copy(local,remote)
    //go io.Copy(remote,local)
    go proxy(local,remote)
    go proxy(remote,local)
}

func proxy(first,second net.Conn) {
    in := make([]byte,20480)
    _,err := first.Read(in)
    if err != nil {
        return
    }
    fmt.Fprintf(os.Stdout,"**:%v\n",string(in))
    _,err = second.Write(in)
    if err != nil {
        return
    }
}
