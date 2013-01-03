

package main


import (
    "fmt"
    "hash/crc32"
)


func main(){

    cs := crc32.ChecksumIEEE([]byte("1234667"))
    fmt.Println(cs)

    cs = crc32.ChecksumIEEE([]byte("dfdfdfdfdf"))
    fmt.Println(cs)
}
