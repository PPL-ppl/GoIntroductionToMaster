package main

import "fmt"

//抽象层

// 水果类(抽象的接口）

type Fruit interface {
	Show()
}

// 工厂类(抽象的接口）

type AbstractFactory interface {
	CreateFruit() Fruit //生产水果类的(抽象）的生产器方法
}

//====基础模块层

type Apple struct {
	Fruit
}

func (apple *Apple) Show() {
	fmt.Println("apple")
}

type Banana struct {
	Fruit
}

func (banana *Banana) Show() {
	fmt.Println("banana")
}

type Pear struct {
	Fruit
}

func (pear *Pear) Show() {
	fmt.Println("pear")
}

// ====基础工厂层
// 具体的苹果工厂

type AppleFactory struct {
	AbstractFactory
}

func (fac *AppleFactory) CreateFruit() Fruit {
	var fruit Fruit
	fruit = new(Apple)
	return fruit
}

// 具体的香蕉工厂

type BananaFactory struct {
	AbstractFactory
}

func (fac *BananaFactory) CreateFruit() Fruit {
	var fruit Fruit
	fruit = new(Banana)

	return fruit
}

// 具体的梨工厂

type PearFactory struct {
	pear AbstractFactory
}

func (fac *PearFactory) CreateFruit() Fruit {
	var fruit Fruit
	fruit = new(Pear)

	return fruit
}
func main() {

	/*	s := new(student)
		s.name = "22"
		name := s.GetName()
		fmt.Println(name)*/
	//需求1：需要一个具体的苹果对象
	//1-先要一个具体的苹果工厂
	var appleFac AbstractFactory
	appleFac = new(AppleFactory)
	//2-生产一个具体的水果
	var apple Fruit
	apple = appleFac.CreateFruit()
	apple.Show()
}
