package main

import (
    "fmt"
)

const (
    VersionTag = 131
)

func writeUint8(u uint8) {
    fmt.Println(u)
}

func main() {
    writeUint8(VersionTag)
}
