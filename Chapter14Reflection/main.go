package main

import (
    "fmt"
    "reflect"
)

func main() {
    ValueOf1()
    ValueOf2()
    StructTag()
    P2P()
    IfFly()
}

type T struct {
}

//动态调用无参方法
func ValueOf1() {
    name := "GoodDinner1"
    t := &T{}
    reflect.ValueOf(t).MethodByName(name).Call(nil)
}
func (t *T) GoodDinner1() {
    fmt.Println("吃顿好的")
}

//动态调用有参方法
func (t *T) GoodDinner2(a int, b string) {
    fmt.Println("吃顿好的")
}
func ValueOf2() {
    name := "GoodDinner2"
    t := &T{}
    param1 := reflect.ValueOf(66)
    param2 := reflect.ValueOf("红烧肉")
    params := []reflect.Value{param1, param2}
    reflect.ValueOf(t).MethodByName(name).Call(params)
}

//动态struct tag解析
type person1 struct {
    Age  int    `json:"age" test:"testAge"`
    Name string `json:"name" test:"testName"`
}

func StructTag() {
    p := person1{Name: "王五", Age: 22}
    reflectType := reflect.TypeOf(p)
    for i := 0; i < reflectType.NumField(); i++ {
        field := reflectType.Field(i)
        if jsonItem, ok := field.Tag.Lookup("json"); ok {
            fmt.Println(jsonItem)
        }
        testItem := field.Tag.Get("test")
        fmt.Println(testItem)
    }
}

type person struct {
    Age  int    `json:"NewAge"`
    Name string `json:"NewName"`
}
type newPerson struct {
    NewAge  int
    NewName string
}

func P2P() {
    t := person{
        Age:  22,
        Name: "王五",
    }
    refType := reflect.TypeOf(t)
    refValue := reflect.ValueOf(t)
    newPerson2 := &newPerson{}
    newVaule := reflect.ValueOf(newPerson2)

    for i := 0; i < refType.NumField(); i++ {
        field := refType.Field(i)
        newTag := field.Tag.Get("json")
        tValue := refValue.Field(i)
        newVaule.Elem().FieldByName(newTag).Set(tValue)
    }
    fmt.Println(newPerson2)
}
func kind() {
    v := "红烧肉"
    t := reflect.TypeOf(v)

    switch t.Kind() {
    case reflect.Int:
        fmt.Println("int")
    case reflect.String:
        fmt.Println("string")
    case reflect.Bool:
        fmt.Println("bool")
    }
}

type IFly interface {
    //接口方法
    Fly()
}
type Bird struct {
    Name string
}

//Bird实现IFly接口
func (b *Bird) Fly() {
}
func IfFly() {
    bird := &Bird{}
    t := reflect.TypeOf((*IFly)(nil)).Elem()
    refType := reflect.TypeOf(bird)
    fmt.Println(refType.Implements(t))
}
