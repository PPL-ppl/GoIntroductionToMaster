package main

import "fmt"

// 水果 没有简单工厂模式

type Fruit struct {
}

func (f *Fruit) Show(name string) {
	if name == "apple" {
		fmt.Print("Apple")
	} else if name == "banana" {
		fmt.Print("Banana")
	} else if name == "pear" {
		fmt.Print("Pear")
	}
}

//创建一个Fruit对象

func NewFruit(name string) *Fruit {
	fruit := new(Fruit)
	if name == "apple" {
		//创建Apple逻辑
	} else if name == "banana" {
		//创建Banana逻辑
	} else if name == "pear" {
		//创建Pear逻辑
	}
	return fruit
}

// 简单工厂模式
type fruit interface {
	Show()
}
type Apple struct {
	fruit
}

func (apple Apple) Show() {
	fmt.Println("apple")
}

type Banana struct {
	fruit
}

func (banana Banana) Show() {
	fmt.Println("banana")
}

type Pear struct {
	fruit
}

func (pear Pear) Show() {
	fmt.Println("pear")
}

type Factory struct {
}

func (fuc Factory) CreateFruit(kind string) fruit {
	var fruit fruit

	if kind == "apple" {
		fruit = new(Apple)
	} else if kind == "banana" {
		fruit = new(Banana)
	} else if kind == "pear" {
		fruit = new(Pear)
	}
	return fruit
}
func main() {
	factory := new(Factory)
	apple := factory.CreateFruit("apple")
	apple.Show()
	banana := factory.CreateFruit("banana")
	banana.Show()
	pear := factory.CreateFruit("pear")
	pear.Show()
}
