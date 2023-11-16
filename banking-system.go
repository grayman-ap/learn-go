package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	AccountHolder string
	Balance       float64
}

func createAccount(holder string) Account {
	return Account{
		AccountHolder: holder,
		Balance:       0,
	}
}

func deposit(acct *Account, amount float64, wg *sync.WaitGroup) {
	defer wg.Done()

	// Simulate time for transaction to complete
	time.Sleep(time.Millisecond * 300)

	acct.Balance += amount

	fmt.Printf("[%s] Deposit of $%.2f Completed, Total Balance = $%.2f\n", acct.AccountHolder, amount, acct.Balance)
}

func withdraw(acct *Account, amount float64, wg *sync.WaitGroup) {
	defer wg.Done()

	// Simulate time for withdraw to complete
	time.Sleep(time.Millisecond * 500)

	if acct.Balance >= amount {
		acct.Balance -= amount
		fmt.Printf("[%s] Withdraw of $%.2f Completed, Current Balance = $%.2f\n", acct.AccountHolder, amount, acct.Balance)
	} else {
		fmt.Printf("[%s] Insufficient balance of $%.2f, Current Balance = $%.2f\n", acct.AccountHolder, amount, acct.Balance)
	}
}

func getBalance(acct Account) float64 {
	return acct.Balance
}

func main() {
	var wg sync.WaitGroup

	// Create 2 Accounts
	account1 := createAccount("Peter Adeshina")
	account2 := createAccount("Jane Light")

	wg.Add(2)
	go deposit(&account1, 10.00, &wg)
	go withdraw(&account2, 50.00, &wg)

	wg.Wait()

	fmt.Printf("[%s] Final Balance: $%.3f\n", account1.AccountHolder, getBalance(account1))
	fmt.Printf("[%s] Final Balance: $%.3f\n", account2.AccountHolder, getBalance(account2))
}
