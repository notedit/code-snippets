package main

import (
    "fmt"
)

type SomeType struct {
    A int
    B string
}

func main() {
    st := SomeType{}
    fmt.Printf("%#v\n",st.A)
    fmt.Printf("%#v\n",st.B)
}
