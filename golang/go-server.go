package main

import (
    "net"
    "fmt"
    "flag"
)

var maxReadOnceSize = 10;

func connectionHandler(connection net.Conn) {
    
    defer connection.Close()
    var ibuf []byte = make([]byte, maxReadOnceSize+1)
    _, err := connection.Read(ibuf[0:maxReadOnceSize])
    ibuf[maxReadOnceSize] = 0

    write, err := connection.Write(ibuf)

    print(write)
    if err != nil {
        print(write)
        print("Write: ",err.String())
    }

    connection.Close()
}

func initServer(hostAndPort string) (l *net.TCPListener) {
    serverAddr,err := net.ResolveTCPAddr("tcp",hostAndPort)

    if err != nil {
        panic(hostAndPort)
    }

    listener, err := net.ListenTCP("tcp",serverAddr)

    if(err != nil){
        panic(err.String())
    }

    println("Listening to:  ",listener.Addr().String())

    return listener
}

func main() {
    flag.Parse()

    if(flag.NArg() != 2) {
        panic("usage: host port")
    }

    hostAndPort := fmt.Sprintf("%s:%s",flag.Arg(0),flag.Arg(1))

    listener := initServer(hostAndPort)

    for {
        connection,err := listener.Accept()

        if err != nil {
            panic(err.String())
        }

        go connectionHandler(connection)
    }
}

