// Package main provides ...
package main

import (
	"fmt"
	bank "gopl.io/ch9/bank1"
	"strconv"
)

func main() {
	// main2()
	testRace()
}

func testRace() {
	go func() {
		bank.Deposit(200)
		fmt.Println("当前储蓄:" + strconv.Itoa(bank.Balance()))
	}()

	go bank.Deposit(100)
	fmt.Println(bank.Balance())
}

func main2() {
	done := make(chan bool)
	m := make(map[string]string)
	m["name"] = "world"
	go func() {
		m["name"] = "data race"
		done <- true
	}()
	fmt.Println("Hello,", m["name"])
	<-done
}
