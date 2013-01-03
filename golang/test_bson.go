
package main


import (
    "fmt"
    "launchpad.net/mgo/bson"
)

type Person struct {
    AAA string
}

func main() {

    person := &Person{AAA:"AAA"}
    bb,_ := bson.Marshal(struct{VAL *Person}{person})

    fmt.Println(bb)
    
    person2 := &Person{}
    r := &struct{VAL *Person}{person2}
    err := bson.Unmarshal(bb,r)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(person2)
}
