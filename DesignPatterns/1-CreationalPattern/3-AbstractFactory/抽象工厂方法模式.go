package main

// 使用场景  固定选项

import "fmt"

// ------抽象层

type AbstractApple interface {
	ShowApple()
}
type AbstractBanana interface {
	ShowBanana()
}
type AbstractPear interface {
	ShowPear()
}

// 抽象工厂

type AbstractFactory interface {
	CreateApple() AbstractApple
	CreateBanana() AbstractBanana
	CreatePear() AbstractPear
}

// 实现层
/*中国产品组*/

type ChinaApple struct {
}

func (ca *ChinaApple) ShowApple() {
	fmt.Println("中国苹果")
}

type ChinaBanana struct {
}

func (ca *ChinaBanana) ShowBanana() {
	fmt.Println("中国苹果")
}

type ChinaPear struct {
}

func (ca *ChinaPear) ShowPear() {
	fmt.Println("中国苹果")
}

type ChinaFactory struct {
}

func (cf *ChinaFactory) CreateApple() AbstractApple {
	var apple AbstractApple
	apple = new(ChinaApple)
	return apple
}

func (cf *ChinaFactory) CreateBanana() AbstractBanana {
	var apple AbstractBanana
	apple = new(ChinaBanana)
	return apple
}
func (cf *ChinaFactory) CreatePear() AbstractPear {
	var apple AbstractPear
	apple = new(ChinaPear)
	return apple
}
func main() {
	var chinaFactory AbstractFactory
	chinaFactory = new(ChinaFactory)
	var chinaApple AbstractApple
	chinaApple = chinaFactory.CreateApple()
	chinaApple.ShowApple()

	var computer Computer
	computer = new(InterFactory)
	cpu := computer.GetCpu()
	cpu.Calculate()

}

//这种模式  Computer增加产品 所有供应商都要增加这个产品

// 提供显卡、CPU、内存

type Computer interface {
	GetXianKa() XianKa
	GetNeiCun() NeiCun
	GetCpu() Cpu
}

// 显卡、CPU、内存

type XianKa interface {
	DisPlay()
}
type NeiCun interface {
	Storage()
}
type Cpu interface {
	Calculate()
}

// 英特尔产品

type InterXianKa struct {
}

func (ca *InterXianKa) DisPlay() {
	fmt.Println("因特尔显卡")
}

type InterNeiCun struct {
}

func (inter *InterNeiCun) Storage() {
	fmt.Println("因特尔内存")
}

type InterCpu struct {
}

func (inter *InterCpu) Calculate() {
	fmt.Println("因特尔Cpu")
}

//因特尔工厂供货

type InterFactory struct {
}

func (inter *InterFactory) GetXianKa() XianKa {
	var xianKa XianKa
	xianKa = new(InterXianKa)
	return xianKa
}
func (inter *InterFactory) GetNeiCun() NeiCun {
	var neiCun NeiCun
	neiCun = new(InterNeiCun)
	return neiCun
}
func (inter *InterFactory) GetCpu() Cpu {
	var cpu Cpu
	cpu = new(InterCpu)
	return cpu
}
