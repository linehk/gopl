package bank

import (
	"fmt"
	"sync"
	"testing"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		Deposit(200)
		fmt.Println("=", Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}

func TestWithdraw(t *testing.T) {
	Deposit(100)

	tests := []struct {
		amount      int
		wantSucceed bool
		wantBalance int
	}{
		{50, true, 50},
		{100, false, 50},
	}
	for i, tt := range tests {
		if gotSucceed := Withdraw(tt.amount); gotSucceed != tt.wantSucceed {
			t.Errorf("%d. succeed: got %v, want %v", i, gotSucceed, tt.wantSucceed)
		}
		if gotBalance := Balance(); gotBalance != tt.wantBalance {
			t.Errorf("%d. balance: got %v, want %v", i, gotBalance, tt.wantBalance)
		}
	}
}

func TestWithdrawConcurrent(t *testing.T) {
	Deposit(10000)

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(amount int) {
			Withdraw(amount)
			wg.Done()
		}(i)
	}

	wg.Wait()

	if got, want := Balance(), 5050; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
