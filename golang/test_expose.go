
package main

import (
    "fmt"
)

type foo struct {
    fn interface{}
}

func (f foo)Execute( fn  func(int,int)) {
    fn(1, 2)
}

func main() {
    f := new(foo)
    fn := func(a int,b int){
        fmt.Println(a,b,"\n")
    }
    f.Execute(fn)
}
