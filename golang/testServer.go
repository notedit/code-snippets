package main

import (
    "recca/rpc"
)

type Args struct {
    A,B int
}

type Reply struct {
    C int
}

type Test struct{}

func (t *Test)Add(args Args,reply *Reply) error {
    reply.C = args.A + args.B
    return nil
}

func main() {
    newServer := rpc.NewServer("localhost",1989)
    newServer.Register(new(Test))
    newServer.Serv()
}
