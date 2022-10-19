package main

import "fmt"

// 手机 贴膜 带壳

// 抽象层

//抽象的手机

type Phone interface {
	Show()
}

//装饰器

type Decorator struct {
	phone Phone //组合
}

func (d *Decorator) Show() {

}

// 实现层
type HuaWei struct {
}

func (d *HuaWei) Show() {
	fmt.Println("show huawei")
}

type XiaoMi struct {
}

func (d *XiaoMi) Show() {
	fmt.Println("show xiaomi")
}

type MoDecorator struct {
	Decorator
}

func (d *MoDecorator) Show() {
	d.phone.Show()
	fmt.Println("贴膜")
}

type KeDecorator struct {
	Decorator
}

func (d *KeDecorator) Show() {
	d.phone.Show()
	fmt.Println("外壳")
}
func NewMoDecorator(phone Phone) Phone {
	return &MoDecorator{Decorator{phone}}
}
func NewKeDecorator(phone Phone) Phone {
	return &KeDecorator{Decorator{phone}}
}
func main() {
	xiaomi := new(XiaoMi)
	xiaomi.Show()
	moDecorator := NewMoDecorator(xiaomi)
	moDecorator.Show()
	KeDecorator := NewKeDecorator(xiaomi)
	KeDecorator.Show()
	decorator := NewMoDecorator(KeDecorator)
	decorator.Show()
}
