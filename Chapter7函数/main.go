package main

import (
    "fmt"
    "strconv"
    "sync"
)

func main() {

    fmt.Println(func1("方法一"))

    fmt.Println(func2("admin", "123"))

    var stars = [3]string{"2", "3", "4"}
    fmt.Println(stars)      //[2 3 4]
    strings := func3(stars) //传递数组 值传递
    fmt.Println(strings)    //[2 3 4]
    strs := &[3]string{"2", "3", "4"}
    fmt.Println(strs) //&[2 3 4]
    func4(strs)       //传递数组  指针数组
    fmt.Println(strs) //&[2 3 2]

    strs2 := []string{"2", "3", "3"}
    func5(strs2)
    fmt.Println(strs2) //[2 3 2]

    var hello hi
    hello = func6
    s1 := hello("wzl")
    fmt.Println(s1)

    hello = func7
    s2 := hello("wzl")
    fmt.Println(s2)

    s3 := func8("3", func7)
    fmt.Println(s3)
    /// 匿名函数
    s4 := func8("4", func(name string) string {
        return name
    })
    fmt.Println(s4)
    f := payOrder(func() float64 {
        return 0.8
    }) //返回是个方法
    result := f("肉", 100)
    fmt.Println(result)

    start()
    TestFood()
    end()
}

//func1为函数名 后面第一个括号内是入参  后面第二个括号内是返回值  {函数体}
func func1(userName string) string {
    return userName
}

func func2(userName string, password string) string {
    if userName == "admin" && password == "123" {
        return "true"
    }
    return "false"
}

//传递数组 值传递
func func3(stars [3]string) [3]string {
    for _, str := range stars {
        str = "222"
        fmt.Println(str)
    }
    return stars
}

/// 传递指针数组 地址传递
func func4(stars *[3]string) {
    stars[2] = "2"
}

/// 传递切片  地址传递
func func5(stars []string) {
    stars[2] = "2"
    stars = append(stars, "2") //外面切片大小不会变
}

/// 定义一个func为类型 符合这种类型的方法都可以使用
type hi func(name string) string

func func6(name string) string {
    return name + "hello"
}
func func7(name string) string {
    return name + "hi"
}
func func8(name string, hi2 hi) string {
    s := hi2(name)
    return s
}

type Discount func() float64
type CheckSum func(name string, price float64) float64

func payOrder(discount Discount) CheckSum {
    var total float64
    return func(name string, price float64) float64 {
        fmt.Println(name + ":" + strconv.FormatFloat(price, 'f', -1, 64))
        total = total + price
        if discount == nil {
            return total
        }
        return total * discount()
    }
}

func lookup(key string) int {
    m := make(map[string]int)
    var mu sync.Mutex
    mu.Lock()
    defer mu.Unlock()
    return m[key]
}
func func10() {
    panic("hello")
}

func start() {
    fmt.Println("程序开始执行")
}
func TestFood() {
    foods := []string{"hello", "ni", "hao"}
    defer func() {
        if err := recover(); err != nil {
            fmt.Println(err)

        }
        fmt.Println("defer finish")
    }()
    fmt.Println(foods[5])
}
func end() {
    fmt.Println("程序结束执行")
}
