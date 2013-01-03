package main

import (
    "github.com/ugorji/go-msgpack"
    "net"
    "net/rpc"
    "fmt"
)

type Args struct {
    A,B int
}

func main(){
    conn,err := net.Dial("tcp","localhost:8080")
    if err != nil {
        fmt.Println(err)
    }
    codec := msgpack.NewRPCCodec(conn)
    client := rpc.NewClientWithCodec(codec)
    arg := Args{7,8}
    var reply  int
    err = client.Call("Arith.Mul",arg,&reply)
    if err != nil {
        fmt.Println(err)
    }
    err = client.Call("Arith.Divide",arg,&reply)
    if err != nil {
        fmt.Println(err)
    }

}
