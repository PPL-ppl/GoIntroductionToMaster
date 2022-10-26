package main

import "fmt"

// 人 装备  用不同装备战斗方式不一样

// 武器策略抽象

type WeaponStrategy interface {
	UseWeapon()
}

// 具体策略

type Ak47 struct {
}

func (a Ak47) UseWeapon() {
	fmt.Println("射击")
}

type Knife struct {
}

func (a Knife) UseWeapon() {
	fmt.Println("刺刀")
}

type Person struct {
	strategy WeaponStrategy
}

// 设置策略

func (p *Person) SetWeaponStrategy(w WeaponStrategy) {
	p.strategy = w
}

func main() {
	p := new(Person)
	p.SetWeaponStrategy(new(Knife))
	p.strategy.UseWeapon()
}
