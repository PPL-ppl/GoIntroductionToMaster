package main

import "fmt"

// 整个流程是固定的

type Beverage interface {

	//煮开水

	BoilWater()

	//冲泡

	Brew()

	//倒入杯中

	PourInCap()

	//添加佐料

	AddThing()

	//是否添加类图

	AddWangThings() bool
}

type template struct {
	b Beverage
}

func (t *template) MakeBeverage() {
	if t == nil {
		return
	}
	//固定流程
	t.b.BoilWater()
	t.b.Brew()
	t.b.PourInCap()

	//是否添加佐料 Hook
	if t.b.AddWangThings() {
		t.b.AddThing()
	}
}

type MakeCoffee struct {
	template
}

func (m *MakeCoffee) BoilWater() {
	fmt.Println("泡水")
}

func (m *MakeCoffee) Brew() {
	fmt.Println("水冲咖啡豆")
}

func (m *MakeCoffee) PourInCap() {
	fmt.Println("将咖啡倒入被子")
}

func (m *MakeCoffee) AddThing() {
	fmt.Println("加入佐料")
}
func (m *MakeCoffee) AddWangThings() bool {
	return true
}
func NewMakeCoffee() *MakeCoffee {
	m := new(MakeCoffee)
	m.b = m
	return m
}

type MakeTea struct {
	template
}

func (m *MakeTea) BoilWater() {
	fmt.Println("泡水")
}

func (m *MakeTea) Brew() {
	fmt.Println("水冲茶叶")
}

func (m *MakeTea) PourInCap() {
	fmt.Println("将茶倒入被子")
}

func (m *MakeTea) AddThing() {
	fmt.Println("加入佐料")
}
func (m *MakeTea) AddWangThings() bool {
	return false
}
func NewMakeTea() *MakeTea {
	m := new(MakeTea)
	m.b = m
	return m
}

func main() {
	coffee := NewMakeCoffee()
	coffee.MakeBeverage()

	tea := NewMakeTea()
	tea.MakeBeverage()
}
