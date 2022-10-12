package main

// 抽象的银行业务
type AbstractBanker interface {
	DoBuSi()
}

// 存款业务
type SaveBank struct {
}

func (SaveBank) DoBuSi() {

}

// 转账业务
type TransferBank struct {
}

func (TransferBank) DoBuSi() {

}

// 股票业务
type SharesBanker struct {
}

func (SharesBanker) DoBuSi() {

}

// 多态
func BankBusiness(banker AbstractBanker) {
	banker.DoBuSi()
}
func main() {
	cb := SaveBank{}
	cb.DoBuSi()
	tb := TransferBank{}
	tb.DoBuSi()
	c := SharesBanker{}
	c.DoBuSi()
	//一个方法 多种结果
	BankBusiness(c)
}
