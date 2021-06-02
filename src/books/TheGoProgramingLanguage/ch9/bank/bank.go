package main

var balance int

// Deposit 银行储蓄
func Deposit(money int) {
	balance = balance + money
}

// Balance 查看储蓄
func Balance() int {
	return balance
}
