package main

import (
    "log"
    "time"
    "fmt"
    "database/sql"
    _ "github.com/bmizerany/pq"
)

type Res struct {
    a time.Time
}

func main(){
    db,err := sql.Open("postgres","sslmode=disable user=user password=password dbname=database")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    var bbb bool
    err = db.QueryRow("SELECT a FROM test3 where id=14;").Scan(&bbb)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Printf("%v\n",bbb)
    //if !r.Next() {
    //    log.Print("not next")
    //}

    //res := Res{}
    //err = r.Scan(&res.a)
    //if err != nil {
    //    log.Print("scan error")
    //    log.Print(err)
    //}

    //fmt.Printf("%#v\n",res)
}
