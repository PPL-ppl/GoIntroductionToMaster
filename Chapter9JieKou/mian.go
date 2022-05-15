package main

import "fmt"

func main() {
    /*    c := new(Chef)
          c.Age = 2
          c.Name = "王五"
          fmt.Printf("%p\n", &c)
          FavCook1(c, "水煮鱼")
    */
    c2 := Chef{
        Name: "2",
        Age:  22,
    }
    c2.Cook1()
    fmt.Println(c2.Name)
    fmt.Printf("%p\n", &c2)
    FavCook1(&c2, "水煮鱼")

    c1 := C{
        Name: "2",
        Age:  22,
    }
    s := c1.Error()
    fmt.Println(s)
    c1.Cook()

    var x interface{}
    x = "hello"
    //类型分支
    switch x.(type) {
    case string:
        fmt.Println("这是字符串")
    case int:
        fmt.Println("这是数字")

    }
}

type Chef struct {
    Name string
    Age  int
}
type C struct {
    Name string
    Age  int
}
type ChefInterface interface {
    Cook1() bool
    FavCook1(foodName string) bool
    Cook2() bool
    FavCook2(foodName string) bool
}

func (c Chef) Cook1() bool {
    fmt.Printf("%p\n", &c)
    c.Name = "lao"
    return true
}
func (c Chef) FavCook1(foodName string) bool {
    fmt.Println(foodName)
    fmt.Printf("%p\n", &c)
    return true

}
func (c *Chef) Cook2() bool {
    fmt.Printf("%p\n", &c)
    return true
}
func (c *Chef) FavCook2(foodName string) bool {
    fmt.Println(foodName)
    fmt.Printf("%p\n", &c)
    c.Age = 23
    return true
}

func (c *C) Cook() bool {
    fmt.Printf("%p\n", &c)
    return true
}

func FavCook1(c *Chef, foodName string) bool {
    var ci ChefInterface = c //结构体实现了接口接口所有方法，那么就实现了这个接口  可以直接赋值给接口 只实现一个不算
    fmt.Printf("%p\n", &c)
    fmt.Printf("%p\n", &ci)
    c.Name = "老六"
    ci.Cook2()
    ci.FavCook2(foodName)

    c2 := ci.(*Chef) //断言 如果（内没有实现接口就报错）  ci.(*C)会保错
    fmt.Println(c2.Name)
    c3 := C{
        Name: "2",
        Age:  2,
    }
    var ci2 error = &c3
    c4 := ci2.(*C) //断言 如果（内没有实现接口就报错）  此时不报错
    fmt.Println(c4.Name)

    c5 := ci.(ChefInterface) //断言为true得c5为type ChefInterface 还是 *Chef
    fmt.Println(c5)
    return true

}

//实现error接口 自定义错误信息
func (c *C) Error() string {
    return "错了"
}

type inter1 interface {
    getName() string
}

func (c Chef) getName() string {
    return c.Name
}
func (c *Chef) setName(name string) {
    c.Name = name
}

//测试代码
/*
c1 := Chef{
Name: "2",
Age:  22,
}
fmt.Println(c1.Name)//2
var c2 inter1=c1
fmt.Println(c1.getName())//2
c1.SetName("3")
fmt.Println(c1.getName())//3  修改内存地址
fmt.Println(c2.getName())//2  接口取到的是副本
*/
