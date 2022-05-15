package main

import "fmt"

func main() {
    //数组
    var a = [6]int{0, 1, 2, 3, 4, 5}
    //创建切片方式一  0为起始下标 6为结束的下标
    var b = a[0:6] //含0不含6
    //创建切片方式二
    var c []int
    c = append(c, 2)
    //创建切片方式三
    ints := make([]int, 2, 3) //2为长度 3为容量 3可省略
    ints = append(ints, 2)

    b = append(b, 3) //会自动扩容 并且会改变之前数组的值 超出之前数组范围的不会影响
    fmt.Println(len(b))
    fmt.Println(cap(b))
    fmt.Println(a[5])
    fmt.Println(b[6])

    var d = [6]int{0, 1, 2, 3, 4, 5} //数组长度6
    e := d[2:3]                      //切片长度为1 容量为数组长度-2=4
    fmt.Println(len(e))
    fmt.Println(cap(e))

    //切片复制
    var ss = []int{0, 8, 2, 3, 4, 5}
    var s = ss[:]
    fmt.Println(s)
    var x = []int{6, 7}
    copy(x, ss)     //长度小的会被修改掉
    fmt.Println(x)  //[0 8]
    fmt.Println(ss) //[0 8 2 3 4 5]

    //在索引2的位置添加值
    ss = append(ss[2:], append([]int{2}, ss[2:]...)...) //[0 8 2 2 3 4 5] 在下标2前面把2加进去  且把2之后的数放到2前面

    ss = append(ss[:2], append([]int{2}, ss[:2]...)...) //[0 8 2 0 8] 在下标2后面面把2加进去  且把2之后的数放到2前面

    ss = append(ss[:2], append([]int{2}, ss[:]...)...) //[0 8 2 0 8 2 2 3 4 5]  在下标2后面面把2加进去  且把加了2后的所有数放到2后面
    ss = append(ss[:2], ss[3:]...)                     //在下标2后面加上下标3后面的数 删除了3

    ss2 := make([]*int, len(ss))
    for i, i2 := range ss {
        ss2[i] = &i2 //获取的i2的物理地址 最终都会等于ss的最后一个值
    }
    //正确写法
    for i, _ := range ss {
        ss2[i] = &ss[i] //获取的ss的每个坐标下的值
    }

}
