package main

import (
    "fmt"
    "time"
)

//协程与通道
func main() {
    func1()
    func2()
    ch1 := make(chan string)
    ch2 := make(chan string)
    //go关键字开启协程 可直接是方法 也可以是写func
    go SendChan(ch1)      //先写入一个数
    go SaveChan(ch1, ch2) //读取出来，写入ch2
    Cook(ch2)             //读取ch2
    fmt.Println("洗碗")

    //无缓冲chan  提供了同步机制，每次发送操作都要对应的接收的操作保持同步
    ch3 := make(chan bool)
    go Eat("生蚝", ch3)
    b := <-ch3
    fmt.Println(b)

    //缓冲chan 发送和接收操作解耦  如果chan满了继续发送会死锁
    ch4 := make(chan bool, 7)
    go Eat("生蚝1", ch4)
    go Eat("生蚝2", ch4)
    b2 := <-ch4
    fmt.Println(b2)
    b3 := <-ch4
    fmt.Println(b3)

    ch5 := make(chan string)
    ch6 := make(chan string)
    ch7 := make(chan string)
    go GetFood1(ch5)
    go GetFood2(ch6)
    go GetFood3(ch7)
    select {
    case r := <-ch5:
        fmt.Println(r)
        break
    case r := <-ch6:
        fmt.Println(r)
        break
    case r := <-ch7:
        fmt.Println(r)
        break
    default:
        fmt.Println("菜还没好")
    }
    //close(ch6)
    time.Sleep(3000)
    V, ok := <-ch6
    if ok {
        fmt.Println(V)
    }
}
func func1() {
    for i := 0; i < 10; i++ {
        go func() {
            fmt.Println(i)
        }()
    }
}
func func2() {
    fmt.Println("打印开始")
    go work()
    fmt.Println("打印结束")
    time.Sleep(1000)
}
func work() {
    fmt.Println("正在工作")
}

//channal通道
func Chan() {
    var chan1 = make(chan bool)
    chan1 <- true //发送
    <-chan1       //接收
}

//单向chan -写
func SendChan(c chan<- string) {
    c <- "买菜"
    close(c)
}

//单向chan -ch1读，ch2写   参数写法c1 <-chan string, c2 chan<- string
func SaveChan(c1 <-chan string, c2 chan<- string) {
    r := <-c1
    c2 <- r + "买肉"
    close(c2)
}

//单向chan -读
func Cook(c <-chan string) {
    for s := range c {
        fmt.Println(s + "已经准备好了")
    }
}

func Eat(foodName string, c chan bool) {
    fmt.Println("我正在吃" + foodName)
    c <- false //写入
}
func GetFood1(c chan string) {
    time.Sleep(3 * time.Second)
    close(c)
}

func GetFood2(c chan string) {
    c <- "蒸好鱼了"
    time.Sleep(3 * time.Second)
}
func GetFood3(c chan string) {
    c <- "生蚝好了"
    time.Sleep(2 * time.Second)
}
