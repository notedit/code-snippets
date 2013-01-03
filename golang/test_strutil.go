package strutil

import (
    "fmt"
)

func convert16to36(string input) string {
    var nummap map[byte]int = {
        "0":0,"1":1,"2":2,"3":3,"4":4,"5":5,
        "6":6,"7":7,"8":8,"9":9,"a":10,"b":11,"c":12,"d":13,"e":14,"f":15,
        }
    var table string = "0123456789abcdefghijklmnopqrstuvwxyz"
    var length int = len(input)
    var result string = ""
    var newlen  int = 0


    for i:= 0;i<length;++i {
        value = value*16 + nummap[input[i]]
        if value >= 36 {

        }
    }

    for  newlen != 0 {
        value := 0
        newlen = 0

        for i:= 0;i < length; ++i {
            value = value * 16 + nummap[input[i]]
            if value >= 36 {
                
            }
        }
    }


}
