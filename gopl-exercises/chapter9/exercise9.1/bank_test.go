package bank

import (
	"fmt"
	"testing"
	"time"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		Deposit(200)
		fmt.Println("=", Balance())
		done <- struct{}{}
	}()

	go func() {
		ok := Withdraw(300)
		fmt.Println(ok, "=", Balance())
		done <- struct{}{}
	}()
	// Bob
	go func() {
		time.Sleep(time.Second)
		Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done
	<-done

	if got, want := Balance(), 100; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
