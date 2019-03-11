// Package bank provides a concurrency-safe bank with one account.
package bank

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdraw = make(chan int)
var boolWithdraw = make(chan bool)

func Desposit(amount int) { deposits <- amount }
func Balance() int        { return <-balances }
func Withdraw(amount int) bool {
	withdraw <- amount
	return <-boolWithdraw
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
			// Do nothing
		case amount := <-withdraw:
			if amount < balance {
				boolWithdraw <- false
			} else {
				boolWithdraw <- true
			}
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
