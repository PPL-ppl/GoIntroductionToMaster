package main

import "fmt"

/*
接口隔离
接口不应该拥有他不使用的方法
*/
/*
合成复用

如果使用继承，会导致父类的任何变换都可能影响到子类的行为。
如果使用对象组合，就降低了这种依赖关系。
对于继承和组合，优先使用组合。

*/

type Cat struct {
}

func (c *Cat) Eat() {

}

// 增加睡觉--继承实现
type CatB struct {
	Cat
}

func (c *CatB) Sleep() {
	fmt.Print("睡觉")
}

// 增加睡觉--组合实现
type CatC struct {
	c Cat
}

func (cc *CatC) Eat(c *Cat) {
	//cc.c.Eat
	c.Eat()
}
func (cc *CatC) Sleep() {
	fmt.Print("睡觉")
}
func main() {

}
