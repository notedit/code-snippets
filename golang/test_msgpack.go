package main

import (
    "msgpack"
    "bytes"
    "fmt"
)

func main() {

    m := map[string]string {"dfdfd": "dfdfd","aaaa": "bbb"}
    buf := &bytes.Buffer{}
    if _,err := msgpack.Pack(buf,m); err != nil {
        panic("oops")
    }
    value,_,err := msgpack.Unpack(buf)
    if err != nil {
        panic("oops")
    }

    fmt.Println("%#v",value)
}
