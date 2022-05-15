package main

import (
    "fmt"
    "strconv"
    "sync"
    "sync/atomic"
    "time"
)

func main() {
    //互斥锁
    for i := 0; i < 10; i++ {
        addMutex(strconv.Itoa(i))
    }
    //读写锁
    s := []string{"老张", "老王", "老李"}
    for _, s2 := range s {
        go CheckSeat(s2)
        go GetSeat(s2)
    }
    time.Sleep(8 * time.Second)
    buySth()
}

//互斥锁
var m1 sync.Mutex

func addMutex(name string) {
    m1.Lock()
    defer m1.Unlock() //函数结束时调用  同一个函数多个defer 逆序执行 从后到前执行  后使用先关闭
    fmt.Print(name + "加锁")
    fmt.Println(name + "解锁")
}

//读写锁
var m2 sync.RWMutex

func GetSeat(name string) { //抢位置
    m2.Lock()
    defer m2.Unlock()
    fmt.Println(name + "抢到了座位")
    time.Sleep(1 * time.Second)
    fmt.Println(name + "离开了座位")
}

func CheckSeat(name string) { //查看是否有空位置
    m2.RLock()
    defer m2.RUnlock()
    fmt.Println(name + "查看了座位")
    time.Sleep(1 * time.Second)
}

//案例
func buySth() {
    var buySth string
    var box sync.RWMutex
    sendCond := sync.NewCond(&box)
    revCond := sync.NewCond(box.RLocker())
    flag := make(chan bool, 5)
    go func() { //协程
        defer func() {
            fmt.Println("已投递，完成~")
            flag <- true
        }()
        box.Lock() //持有box锁
        for buySth == "已投递" {
            sendCond.Wait() //等待通知
        }
        buySth = "已投递"
        box.Unlock()
        revCond.Signal() //单一通知
    }()

    go func() {
        defer func() {
            fmt.Println("货物已经取走")
            flag <- true
        }()
        box.RLock()
        for buySth == "" {
            revCond.Wait()
        }
        buySth = ""
        box.RUnlock()
        sendCond.Signal()
    }()

    a1 := <-flag
    fmt.Println(a1)
    a2 := <-flag
    fmt.Println(a2)
}

func Atomic() { //原子操作
    //加法
    var num1 uint64
    atomic.AddUint64(&num1, 70)
    //比较和交换
    var num2 uint64
    swapUint64 := atomic.CompareAndSwapUint64(&num2, 70, 60)
    fmt.Println(swapUint64)
    //加载
    var num3 uint64
    atomic.AddUint64(&num3, 70)
    //存储
    var num4 uint64
    atomic.AddUint64(&num4, 70)
    //交换
    var num5 uint64
    atomic.AddUint64(&num5, 70)
}
