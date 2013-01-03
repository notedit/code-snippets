package main

import (
    "fmt"
    "database/sql"
    _ "github.com/lxn/go-pgsql"
)

func main() {

    db,err := sql.Open("postgres","dbname=database user=user password=password port=5432")
    
    if err != nil {
        fmt.Println(err)
    }
    defer db.Close()

    //rs,err := db.Query("SELECT id,title,author_userid,last_modify_userid,last_reply_userid,body,ctype,recommend_count,date_create,date_last_reply,show,disable_reply,tag_id FROM content WHERE  show=true AND id=151")
    rs,err := db.Query("SELECT aaa FROM test8")
    if err != nil {
        fmt.Println(err)
    }
    rs.Next()
    fmt.Println(rs)
    var bb string
    err = rs.Scan(&bb)
    fmt.Println(err)
    if err != nil {
        fmt.Println(bb)
    }
   // err = db.QueryRow("SELECT lastval()").Scan(&rid)
   // fmt.Println(rid)
   // fmt.Printf("%#n\n",r)
   // if err != nil {
   //     fmt.Println(err)
   //}
}

