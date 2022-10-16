package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//字符串
	//定义变量
	msg := "hello"
	fmt.Println(msg) //hello
	//修改变量
	msg = "hello word"
	fmt.Println(msg) //hello word
	var aaa int = 22
	fmt.Printf("string:%d,内存地址：%p", aaa, &aaa)
	//转换成大写
	upper := strings.ToUpper(msg)
	fmt.Println(upper)
	//字符串拼接的三种方法
	msg1 := "hello"
	msg2 := "word"
	msg3 := msg1 + " " + msg2
	fmt.Println(msg3)
	join := strings.Join([]string{"hello", "word"}, " ")
	fmt.Println(join)
	fmt.Sprintf("%s:%s", msg1, msg2)
	msg4 := bytes.Buffer{}
	msg4.WriteString("hello")
	msg4.WriteString("word")
	fmt.Println(msg4.String())
	msg5 := "hello word"
	//制表符 tab键
	fmt.Println("\t" + msg5)
	//换行符
	fmt.Println("\n" + msg)
	//清除空格 空格可以替换成任何一个数 进行清除
	strings.Trim(msg5, " ")
	//数字
	//整型
	var a int
	a = 2
	fmt.Println(a + 7)
	fmt.Println(a - 7)
	fmt.Println(a * 7)
	fmt.Println(a / 7)
	//浮点型
	var b float32 = 2.00000101 //2.000001 精度为小数点后6位
	fmt.Println(b)
	var c float64 = 2.0000000121312312 //2.0000000121312312  精度为小数点后15位
	fmt.Println(c)

	/*age := 25
	  println("年龄是：" + age) 类型不匹配*/
	age := 2
	fmt.Println("年龄是：" + strconv.Itoa(age)) //对int类型转换成string
	/*变量交换*/
	var aa int = 2
	var cc int = 3
	fmt.Println(aa, cc)
	aa, cc = cc, aa
	fmt.Println(aa, cc)
	/*常量*/

	const (
		c1 = iota
		c2 = iota
		c3 = 222
		c4
		c5 = iota
		c6
	)
	fmt.Println(c1, c2, c3, c4, c5, c6) //0 1 222 222 4 5
}
