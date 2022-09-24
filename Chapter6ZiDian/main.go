package main

import "fmt"

func main() {
	//字典
	//新建map方式一
	map1 := map[string]int{"hello": 2}
	fmt.Print("----")
	i2 := len(map1)
	fmt.Print(i2)
	fmt.Print("----")
	map1["hj"] = 2

	var map2 map[string]int = map[string]int{"hello": 2}
	map2["hj1"] = 2
	map2["hj2"] = 3
	fmt.Println(map2["hj"])
	//新建map方式二
	errorMap := make(map[string]string)
	errorMap["hello"] = "2"
	i, ok := errorMap["he"]
	if !ok {
		fmt.Println(1)
	} else {
		fmt.Println("2" + i)
	}
	errorMap["hello2"] = "2"
	//delete(errorMap, "hello")
	fmt.Println(errorMap["hello"])
	for i, s2 := range map2 {
		map1[i] = s2
	}
	fmt.Println(map1)

	fmt.Print("----")
	m1 := make(map[string]string, 10)
	i3 := len(m1)
	fmt.Print(i3)
	m2 := make(map[string]string)
	m2["hell0"] = "222"
	i3 = len(m2)
	fmt.Print(i3)
	fmt.Print("----")
}
