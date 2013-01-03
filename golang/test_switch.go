
package main

import (
    "fmt"
)

func main() {
    aa := "aa"
    switch  {
        case (bb := "bb"); bb == aa:
            fmt.Println("dfdfdfdfd")
        default:
            fmt.Println("wo cha")
    }
}
