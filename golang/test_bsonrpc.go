package main

import (
    oocrpc "msgpack-oocrpc/rpc"
    "errors"
)


type Args struct {
	A, B int
}

type Reply struct {
	C int
}

type Arith int

func (t *Arith) Add(args *Args, reply *Reply) error {
	reply.C = args.A + args.B
	return nil
}

func (t *Arith) Mul(args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	return nil
}

func (t *Arith) Div(args *Args, reply *Reply) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	reply.C = args.A / args.B
	return nil
}


func (t *Arith) NError(args *Args, reply *Reply) error {
	return errors.New("normalerror")
}

func (t *Arith)SimpleValue(arg *int,reply *bool) error {
    if *arg == 2 {
        *reply = true
    }
    return nil
}

func main() {
    newServer := oocrpc.NewServer("localhost",9090)
    newServer.Register(new(Arith))
    newServer.Serv()
}
