
package main


import (
    "fmt"
    "github.com/edsrzf/tnetstring-go"
)

type TestTnetstring struct {
    Int         int
    String      string
    String2     string
}

func main() {

    var bb string

    testdata := TestTnetstring{Int:12345,String:"aaaa"}
    testdata2 := TestTnetstring{Int:12345,String:"aaaa",String2:""}
    bb,_ = tnetstring.Marshal(testdata)
    bb2,_ :=tnetstring.Marshal(testdata2)
    fmt.Println(bb)
    fmt.Println(bb2)

    var cc TestTnetstring

    tnetstring.Unmarshal(bb2,&cc)
    fmt.Printf("%#v\n",cc)
}
