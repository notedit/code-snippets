package main

import (
    "fmt"
)

type Atype struct{
    A string
    B string
}

func (a *Atype)Read(){

}

func main(){

    atype := new(Atype)
    fmt.Printf("%#v",atype.Read)
}
