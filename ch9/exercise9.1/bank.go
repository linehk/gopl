package bank

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

type draw struct {
	amount  int
	succeed chan bool
}

var withdraws = make(chan draw)

func Withdraw(amount int) bool {
	succeed := make(chan bool)
	withdraws <- draw{amount, succeed}
	return <-succeed
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case draw := <-withdraws:
			if draw.amount <= balance {
				balance -= draw.amount
				draw.succeed <- true
			} else {
				draw.succeed <- false
			}
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
