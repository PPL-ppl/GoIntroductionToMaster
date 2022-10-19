package main

import "fmt"

// 代练 帮助升级等其他操作
//
// 海外代购(代理)  除了下面的购物  还有辨别真伪 海关安检
//
// 韩国购物  美国购物  非洲购物
type Goods struct {
	Kind string
	Fact bool
}

// 抽象层

type Shopping interface {
	Buy(goods *Goods)
}

//实现层

type KoreaShopping struct {
}

func (ks *KoreaShopping) Buy(goods *Goods) {
	fmt.Println("在韩国购物,买了:" + goods.Kind)
}

type AmericaShopping struct {
}

func (am *AmericaShopping) Buy(goods *Goods) {
	fmt.Println("在美国购物,买了:" + goods.Kind)
}

type AfricaShopping struct {
}

func (as *AfricaShopping) Buy(goods *Goods) {
	fmt.Println("在非洲购物,买了:" + goods.Kind)
}

type OverSeasProxy struct {
	shopping Shopping
}

func (overSeasProxy *OverSeasProxy) Buy(goods *Goods) {
	//1 辨别真伪
	//2 调用要被代理的购物方式
	//3 海关安检
	if overSeasProxy.distinguish(goods) == true {
		overSeasProxy.shopping.Buy(goods) //overSeasProxy.shoppings取实际购买方法
		overSeasProxy.check(goods)
	}

}

func NewProxy(shopping Shopping) Shopping {
	return &OverSeasProxy{shopping}
}

// 质检
func (overSeasProxy *OverSeasProxy) distinguish(goods *Goods) bool {
	fmt.Println("对" + goods.Kind + "辨别真伪")
	if goods.Fact == false {
		fmt.Println("发现假货" + goods.Kind + "，不要购买")
	}
	return goods.Fact
}

// 海关安检
func (overSeasProxy *OverSeasProxy) check(goods *Goods) {
	fmt.Println("对" + goods.Kind + "安检")
}

func main() {
	gi := new(Goods)
	gi.Kind = "无聊"
	gi.Fact = true
	var shopping Shopping
	shopping = new(KoreaShopping)
	shopping.Buy(gi)

	var krShopping Shopping
	krShopping = new(KoreaShopping)
	newProxy := NewProxy(krShopping)
	newProxy.Buy(gi)

	//美女
	wonman := new(MeiNv)
	//通过代理见美女
	business := Proxy(wonman)
	business.TieTie()
}

//实际案例

type WoMan interface {
	TieTie()
}
type MeiNv struct {
}

func (lp MeiNv) TieTie() {
	fmt.Println("贴贴")
}

type Business struct {
	WoMan
}

func Proxy(woman WoMan) WoMan {
	return &Business{woman}
}
func (lp *Business) TieTie() {
	lp.WoMan.TieTie()
}
