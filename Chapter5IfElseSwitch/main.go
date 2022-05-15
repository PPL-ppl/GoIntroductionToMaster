package main

import "fmt"

func main() {

    name := "你好"
    if name == "你好" {
        fmt.Println("你好")
    } else if name == "" {
        fmt.Println("不好")
    } else {
        fmt.Println("不好")
    }

    price := []int{32, 12, 32, 11, 10}
    switch 32 {
    case price[0]:
        fmt.Println("0")
    case price[3]:
        fmt.Println("2")
        break
    }

    var s = "中国人"
    for i, v := range s {
        fmt.Printf("%d %s 0x%x\n", i, string(v), v)
    }
    var sl = []int{1, 2, 3, 4, 5, 6}

    for i := 0; i < len(sl); i++ {
        if sl[i]%2 == 0 { // 忽略切片中值为偶数的元素
            continue //偶数就不向下执行
        } else {
            fmt.Println(i)
        }
    }
loop: //与上面等同
    for i := 0; i < len(sl); i++ {
        if sl[i]%2 == 0 { // 忽略切片中值为偶数的元素
            continue loop //偶数就不向下执行
        } else {
            fmt.Println(i)
        }
    }
    var sl2 = [][]int{
        {1, 34, 26, 35, 78},
        {3, 45, 13, 24, 99},
        {101, 13, 38, 7, 127},
        {54, 27, 40, 83, 81},
    }
loop1:
    for i := 0; i < len(sl2); i++ {
        for j := 0; j < len(sl2[i]); j++ {
            if sl2[i][j] == 13 {
                fmt.Printf("found 13 at [%d, %d]\n", i, j)
                continue loop1 //跳出内层循环
            } else {
                fmt.Printf("found 13 at [%d, %d]\n", i, j)
            }
        }

    }
loop2:
    for i := 0; i < len(sl2); i++ {
        for j := 0; j < len(sl2[i]); j++ {
            if sl2[i][j] == 22 {
                fmt.Printf("found gold at [%d, %d]\n", i, j)
                break loop2 //结束外层循环
            }
        }
    }
}
