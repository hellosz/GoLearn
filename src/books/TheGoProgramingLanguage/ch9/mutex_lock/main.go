// 使用sync.Mutex实现锁
package main

import "sync"

func main() {

}

var (
	balances int
	sema     sync.Mutex // 互斥锁
)

// Balance 查看当前储蓄
func Balance() int {
	// 加锁（互斥量保护的变量应该紧跟在后面，否则需要注释进行说明）
	sema.Lock()
	// 保存当前银行金额的快照
	snapshot := balances
	// 释放锁
	sema.Unlock()
	return snapshot
}

// Deposit 储蓄
func Deposit(balance int) {
	// 加锁
	sema.Lock()
	balances = balances + balance
	// 释放锁
	sema.Unlock()
}
