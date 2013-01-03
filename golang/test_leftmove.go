
package main

import (
    "fmt"
    "encoding/binary"
    "bytes"
)

func main(){
    var i int64 =  258
    fmt.Println(byte(i))
    a := new(bytes.Buffer)
    err := binary.Write(a,binary.LittleEndian,i)
    if err != nil {
        fmt.Printf("%v\n",err)
    }
    fmt.Printf("%#vdffdfdf\n",a.Bytes())
}
