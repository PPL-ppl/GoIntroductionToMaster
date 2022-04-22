package main

import (
    "encoding/json"
    "fmt"
    "time"
)

func main() {

    //该定义方式必须逐级定义
    chef1 := Chef{
        "李四",
        2,
        Honor{}, //Honor{"na", time.Now()}建议初始化嵌套结构体 调用可能会出问题
        nil,
    }
    fmt.Println(chef1.GetTime)

    //通过.直接跨级定义
    chef2 := new(Chef)
    chef2.Name = "王五"
    chef2.Age = 22
    honor := *new(Honor)
    honor.Title = "金牌厨师"
    honor.GetTime = time.Now()
    chef2.Honor = honor
    fmt.Println(chef2.Trainee)

    //匿名结构体
    chef3 := struct {
        name string
        age  int
    }{
        "hello",
        2,
    }
    fmt.Println(chef3.age)

    //Struct转Json
    marshal, _ := json.Marshal(chef2)
    fmt.Println(string(marshal))

    chef4 := &Chef{
        "李四",
        2,
        Honor{}, //Honor{"na", time.Now()}建议初始化嵌套结构体 调用可能会出问题
        nil,
    }
    fmt.Printf("%p\n", &chef4)
    cook2 := chef4.FavCook("水煮鱼")
    fmt.Println(chef4.Name) //老六
    fmt.Println(cook2)

    cook3 := chef4.Cook("水煮鱼") //指针调用值方法同样会改变属性值
    fmt.Println(chef4.Name)    //老六
    fmt.Println(cook3)

    cook1 := chef2.Cook("水煮肉片")
    fmt.Println(chef2.Name) //王五
    fmt.Println(cook1)
}

type Chef struct {
    Name string
    Age  int
    Honor
    Trainee *Chef //参数是本身的话需要是指针
}

type Honor struct {
    Title   string
    GetTime time.Time
}

func (c Chef) Cook(cai string) string { //给厨师添加一个Cook方法
    c.Name = "老六"
    return c.Name + "会炒的菜：" + cai
}

func (c *Chef) FavCook(cai string) string { //给厨师添加一个FavCook方法  接收者是指针类型  对Chef的属性修改会修改对象属性
    c.Name = "老六"
    return c.Name + "会炒的菜：" + cai
}
