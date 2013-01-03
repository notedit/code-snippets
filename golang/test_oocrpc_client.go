package main

import (
    "fmt"
    oocrpc "msgpack-oocrpc/rpc"
)

type Args struct {
    A, B int
}

type Reply struct {
    C int
}

func main() {
    client := oocrpc.New("localhost:9091")
    fmt.Println("string....")
    args := &Args{7,8}
    reply := new(Reply)
    err := client.Call("Arith.Add",args,reply)
    if err != nil {
        fmt.Println(err.Error())
    }
}
