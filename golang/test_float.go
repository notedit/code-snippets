
package main

import (
    "fmt"
    "unsafe"
    "encoding/binary"
)


func main(){
    var a float64 = 12.3
    var b uint64 = *(*uint64)(unsafe.Pointer(&a))

    fmt.Println(a)
    c := make([]byte,8)
    binary.BigEndian.PutUint64(c,b)
    fmt.Println(c)
}
