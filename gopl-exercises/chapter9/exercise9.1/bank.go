// Package bank provides a concurrency-safe bank with one account.
package bank

var deposits = make(chan int)              // send amount to deposit
var balances = make(chan int)              // receive balance
var withdraws = make(chan withdrawRequest) // receive balance
type withdrawRequest struct {
	amount int
	ok     chan<- bool
}

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func Withdraw(amount int) bool {
	var withdrawsResults = make(chan bool) //每个withdraw应该创建自己的信道， 不应该用公用的信道，否则无法区分是谁的返回结果
	withdraws <- withdrawRequest{amount, withdrawsResults}
	return <-withdrawsResults
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case withdrawRequest := <-withdraws:
			amount, ok := withdrawRequest.amount, withdrawRequest.ok
			if amount > balance {
				ok <- false
			} else {
				balance -= amount
				ok <- true
			}
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
