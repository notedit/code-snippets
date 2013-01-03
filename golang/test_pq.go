package main

import (
    "log"
    "time"
    "fmt"
    "database/sql"
    "utils"
    _ "github.com/bmizerany/pq"
)

type Res struct {
    a time.Time
}

func main(){
    db,err := sql.Open("postgres","sslmode=disable user=user port=5432 password=password dbname=database")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    salt := utils.GenSalt(10)
    h := sha1.New()
    
}
