
package main

import (
    "fmt"
)

func main() {

    s := [...]string{"dfdf","dfdfd","aaaaa"}
    fmt.Printf("%v,%v",s[0:len(s)-1])
}
