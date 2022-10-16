package main

/*
里氏代换

任何抽象类(interface接囗）出现的地方都可以用他的实现类进行替换，
实际就是虚拟机制，语言级别实现面向对象功能。
*/
/*
依赖倒转

依赖于抽象（接囗），不要依赖具体的实现（类),也就是针对接囗编程。
*/
import "fmt"

//司机张三、李四 汽车 奔驰、宝马
//1-张三 开 奔驰
//2-李四开宝马

type Benz struct {
}

func (b *Benz) Run() {
	fmt.Print("Benz is running")
}

type BMW struct {
}

func (b *BMW) Run() {
	fmt.Print("BMW is running")
}

// 司机张三

type Zhang3 struct {
}

func (b *Zhang3) RunBenz(benz *Benz) {
	fmt.Print("Benz is running")
	benz.Run()
}
func (b *Zhang3) RunBMW(BMW *BMW) {
	fmt.Print("BMW is running")
	BMW.Run()
}

//司机李四

type Li4 struct {
}

func (b *Li4) RunBMW(bmw *BMW) {
	fmt.Print("BMW is running")
	bmw.Run()
}
func (b *Li4) RunBenz(benz *Benz) {
	fmt.Print("Benz is running")
	benz.Run()
}

// 旧版本
/*func main() {
	a := Li4{}
	b := &BMW{}
	a.RunBMW(b)
}*/

//----->抽象层<-------

type Car interface {
	Run()
}

type Driver interface {
	Driver(car Car)
}

// ----->实现层<-------
type BenZ struct {
}

func (b *BenZ) Run() {
}

type Bmw struct {
}

func (b *Bmw) Run() {
}

type ZhangSan struct {
}
type liSi struct {
}

func (b *ZhangSan) Driver(car Car) {
	car.Run()
}
func (b *liSi) Driver(car Car) {
	car.Run()
}

//----->逻辑层<-------

func main() {
	var benz Car
	benz = new(BenZ)
	var zhangSan Driver
	zhangSan = new(ZhangSan)
	zhangSan.Driver(benz)

	var bmw Car
	bmw = new(Bmw)
	var lisi Driver
	lisi = new(liSi)
	lisi.Driver(bmw)

	zhangSan.Driver(bmw)
}
