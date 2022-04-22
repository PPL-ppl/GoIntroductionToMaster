package main

import (
    "fmt"
)

/*
最先执行的方法
执行完后自上向下执行import中的init方法
同一个包只执行一次
*/
//全局变量 放在堆中
var name int

func init() {
    //在init方法中进行初始化
    name = 2
}
func main() {
    fmt.Println("hello word Go")
    //变量定义 局部变量 放在栈上 超过栈的大小分配到堆上
    var a int
    fmt.Println(a) //默认值0  int8 int16 int等区别在于字节数不一样
    var b string
    fmt.Println(b) //默认值""
    var c bool
    fmt.Println(c) //默认值false
    var d map[string]int
    fmt.Println(d) //默认为nil
    fmt.Println(len(d))
    //短变量
    e := 2
    fmt.Println(e) //e的类型取决与赋值
    //指针
    var point string
    point1 := &point //取point地址给指针point1
    *point1 = "2"
    point2 := *point1 //通过*point1获取point的值
    *point1 = "3"
    fmt.Println(point1 != nil) //true
    fmt.Println(point + "0")   //30
    fmt.Println(point1)        //内存地址
    fmt.Println(point2 + "1")  //20
    //内置函数 new 可使用在表达式中
    fmt.Println(new(int))
}

//作用域
var out string = "在外面吃"

func In() {
    out := "在家里吃"
    {
        out := "在桌子上吃"
        fmt.Println(out) //在桌子上吃
    }
    fmt.Println(out) //在家里吃
}
