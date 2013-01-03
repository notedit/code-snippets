
package main

import (
    "encoding/binary"
    "fmt"
)


func main() {
    
    var test int16 = -2 

    b := make([]byte,4)
    binary.BigEndian.PutUint16(b,uint16(test))
    fmt.Printf("%v\n",b)
}
