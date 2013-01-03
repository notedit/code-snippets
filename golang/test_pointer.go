package main 

import (
    "fmt"
)

type Res struct {
    Reply string
}


func parseRes(flag int) Res {
    res := &Res{Reply:"teststete"}
    if flag == 1 {
        return *res
    } else {
        return Res{Reply:"aaaa"}
    }
    return Res{Reply:"end"}
}

func main(){
    var r Res
    r = parseRes(1)
    fmt.Printf("%v\n",r)
    r = parseRes(2)
    fmt.Printf("%v\n",r)
}
