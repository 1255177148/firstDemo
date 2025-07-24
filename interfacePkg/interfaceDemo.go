package main

import "fmt"

type PayMethod interface {
	Account
	Pay(amount float64) bool
}

type Account interface {
	GetBalance() float64
}

/*
信用卡结构体
*/
type CreditCard struct {
	balance float64
	limit   float64
}

func (c *CreditCard) GetBalance() float64 {
	return c.balance
}

func (c *CreditCard) Pay(amount float64) bool {
	if c.balance+amount > c.limit {
		fmt.Println("信用卡支付失败，超出额度")
		return false
	}
	c.balance += amount
	fmt.Printf("信用卡支付成功：%f", amount)
	return true
}

type DebitCard struct {
	balance float64
}

func (d *DebitCard) GetBalance() float64 {
	return d.balance
}

func (d *DebitCard) Pay(amount float64) bool {
	if amount > d.balance {
		fmt.Println("支付失败，借记卡余额不足")
	}
	d.balance -= amount
	fmt.Printf("借记卡支付成功，%f", amount)
	return true
}

func handlePay(p PayMethod, amount float64) {
	if p.Pay(amount) {
		switch t := p.(type) {
		case *CreditCard:
			fmt.Println("购买成功，已用额度：", t.GetBalance())
		case *DebitCard:
			fmt.Println("购买成功，剩余余额：", t.GetBalance())
		default:
			fmt.Println("未知卡")
		}
	} else {
		fmt.Println("支付失败")
	}
}

func main() {
	creditCard := &CreditCard{balance: 30, limit: 100}
	var debitCard PayMethod = &DebitCard{balance: 100}
	handlePay(creditCard, 20)

	handlePay(debitCard, 30)
}
