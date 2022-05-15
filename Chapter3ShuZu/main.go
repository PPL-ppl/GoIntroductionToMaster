package main

import (
    "fmt"
    "strconv"
)

func main() {
    //数组
    var list = [2]string{"2", "2"}
    fmt.Println(list)
    list[0] = "1"
    list[1] = "2"
    fmt.Println(list)
    var list2 = [2]string{"1", "2"}
    fmt.Println(list2 == list) //true
    //i为下标 s为值
    for i, s := range list2 {
        fmt.Println(strconv.Itoa(i) + ":" + s)
    }
}
