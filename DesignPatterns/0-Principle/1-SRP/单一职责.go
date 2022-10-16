package main

/*
类的职责单一，对外只提供一种功能，而引起类变化的原因都应该只有一个
*/

import "fmt"

/*
type Clothes struct {
}

func (c Clothes) OnWork() {

}

func (c Clothes) OnShop() {

}

	func main() {
		c := Clothes{}
		fmt.Print("在工作")
		c.OnWork()
		fmt.Println("在逛街")
		c.OnWork()
	}
*/
type ClothesShop struct {
}

func (c ClothesShop) OnShop() {
}

type ClothesWork struct {
}

func (c ClothesWork) OnWork() {

}
func main() {
	s := ClothesShop{}
	fmt.Println("在工作")
	s.OnShop()
	w := ClothesWork{}
	fmt.Println("在逛街")
	w.OnWork()
}
