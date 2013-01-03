package main

import (
    "fmt"
    "bytes"
)

func main(){
    var a int32 = -120
    fmt.Printf("%b\n",uint32(a))

    b := bytes.NewBufferString("lifeistooshorttowait")
    fmt.Println(b.Bytes())
    fmt.Println(uint64(1)<<63)
}
