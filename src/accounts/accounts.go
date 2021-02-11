package accounts

import "fmt"

// Account struct
type Account struct { //export하기 위해서 첫 문자는 대문자로 설정한다
	owner   string //private thing
	balance int    //private thing
}

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
func (a Account) Deposit(amount int) {
	fmt.Println("Gonna deposit", amount)
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}
