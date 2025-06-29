package main

import (
	"errors"
	"fmt"
	"sync"
)

type account struct {
	balance int
	sync.RWMutex
}

func (a *account) Balance() int {
	a.RLock()
	defer a.RUnlock()
	return a.balance
}

func (a *account) deposite(amount int) {
	a.Lock()
	a.balance = a.balance + amount
	a.Unlock()
}

func (a *account) withdraw(amount int) error {
	a.Lock()
	defer a.Unlock()
	if a.balance >= amount {
		a.balance = a.balance - amount
		return nil
	}
	return errors.New("balance invalid")
}

func main() {

	done := make(chan struct{})
	acc := account{balance: 1000}

	go func() {
		defer func() { done <- struct{}{} }()
		err := acc.withdraw(500)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("successfully withdraw")
	}()

	go func() {
		defer func() { done <- struct{}{} }()
		acc.deposite(200)

		fmt.Println("deposite successfully")
	}()

	<-done
	<-done
	fmt.Println(acc.balance)

}
