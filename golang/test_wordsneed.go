package main 

import (
    "fmt"
    "math"
)

func wordsNeed(i uint) uint {
    if i == math.MaxUint32 {
        return math.MaxUint32>>5
    } else if (i == 0) {
        return 1
    }
    return (i+31) >> 5
}

func main() {
    a := uint(32)
    b := uint(64)
    c := uint(128)

    nd := wordsNeed(a)
    fmt.Println(nd)
    nd = wordsNeed(b)
    fmt.Println(nd)
    nd = wordsNeed(c)
    fmt.Println(nd)
}
