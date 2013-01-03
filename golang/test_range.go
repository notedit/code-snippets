package main

import(
    "fmt"
    "time"
)


func testRange(errc chan error){
    errc <- fmt.Errorf("fafafa")
    time.Sleep(time.Duration(2 * time.Second))
    errc = nil
}

func main(){
    errchan := make(chan error)

    go testRange(errchan)
    for e := range errchan {
        fmt.Println(e)
    }
}
