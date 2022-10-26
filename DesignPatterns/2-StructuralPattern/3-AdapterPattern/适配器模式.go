package main

import "fmt"

// 被适配的角色

type V220 struct {
}

func (v *V220) Use220v() {
	fmt.Println("220")
}

//适配器

type Adapter struct {
	*V220
}

// 适配器实现V5接口

func (ad *Adapter) Use5v() {
	fmt.Println("使用适配器充电")
	ad.V220.Use220v()
}

//适配器传如V220 返回适配器  适配器调用适配者的方法

func NewAdapter(v220 *V220) *Adapter {
	return &Adapter{v220}
}

// 适配目标

type V5 interface {
	Use5v()
}

// 业务类

type Phone struct {
	V5
}

func NewPhone(v V5) *Phone {
	return &Phone{v}
}

// 普通业务类

func (phone *Phone) Charge() {
	fmt.Println("进行充电")
	phone.V5.Use5v()
}

func main() {
	phone := NewPhone(NewAdapter(new(V220)))
	phone.Charge()

	h := Hero{"概论", new(DaBaoJ)}
	h.Skill()
	D := Hero{"概论", newAdapterAttack(new(PowerOff))}
	D.Skill()

}

//实际案例

// 技能

type Attack interface {
	Fight()
}

// 英雄

type Hero struct {
	Name   string
	attack Attack
}

func (h *Hero) Skill() {
	fmt.Println("使用了技能")
	//最终是调用这个适配者的方法
	h.attack.Fight()
}

//具体的扩展技能

type DaBaoJ struct {
}

func (da *DaBaoJ) Fight() {
	fmt.Println("大宝剑")
}

type PowerOff struct {
}

func (da *PowerOff) ShutDown() {
	fmt.Println("即将关机")
}

// 适配器

type AdapterAttack struct {
	*PowerOff
}

func newAdapterAttack(p *PowerOff) *AdapterAttack {
	return &AdapterAttack{p}
}
func (da *AdapterAttack) Fight() {
	da.PowerOff.ShutDown()
}
