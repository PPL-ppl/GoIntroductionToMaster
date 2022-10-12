package main

type Banker struct {
}

// 存款
func (c Banker) Save() {

}

// 转账
func (c Banker) Transfer() {

}

// 支付
func (c Banker) Pay() {

}

// 股票
func (c Banker) Shares() {

}

func main() {
	b := Banker{}
	b.Save()
	b.Transfer()
	b.Pay()
	b.Shares()
}
