// 使用通道实现锁机制
package main

func main() {

}

var (
	balances int
	sema     = make(chan struct{}, 1) // 缓冲池长度为1
)

// Balance 查看当前储蓄
func Balance() int {
	// 加锁
	sema <- struct{}{}
	// 保存当前银行金额的快照
	snapshot := balances
	// 释放锁
	<-sema
	return snapshot
}

// Deposit 储蓄
func Deposit(balance int) {
	// 加锁
	sema <- struct{}{}
	balances = balances + balance
	// 释放锁
	<-sema
}
