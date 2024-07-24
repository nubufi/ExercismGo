package account

import (
	"sync"
)

// Define the Account type here.
type Account struct {
	balance int64
	isOpen bool
	mu *sync.Mutex
} 
func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	
	return &Account{amount,true,&sync.Mutex{}}
} 
 
func (a *Account) Balance() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	return a.balance,a.isOpen
} 
 
func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if !a.isOpen {
		return 0,false
	}
	if amount < 0 && a.balance + amount <  0{
		return 0,false
	}

	a.balance += amount

	return a.balance,true
}

func (a *Account) Close() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	
	if !a.isOpen {
		return 0,false
	}

	var balance int64
	balance += a.balance

	a.balance = 0
	a.isOpen = false

	return balance, true
}
