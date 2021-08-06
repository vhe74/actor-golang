package main

import (
	"log"
	"math/rand"
	"time"
)

// action payload struct
type BalanceAction struct {
	Kind   string
	Amount int
}

type Actor struct {
	Inbox   chan BalanceAction
	Name    string
	balance int
}

func (A *Actor) Init(chansize int) {
	A.Inbox = make(chan BalanceAction, chansize)
}

func (A *Actor) Run() {
	for {
		select {
		case action, ok := <-A.Inbox:
			if ok {
				if action.Kind == "plus" {
					A.balance += action.Amount
					log.Println("[", A.Name, "] plus", action.Amount, " Balance is now ", A.balance)
				} else if action.Kind == "minus" {
					A.balance -= action.Amount
					log.Println("[", A.Name, "] minus", action.Amount, " Balance is now ", A.balance)
				}
			} else {
				log.Println("[", A.Name, "] Inbox closed")
				break
			}
		}
	}
}

func (A *Actor) SetBalance(b int) { A.balance = b }

func (A *Actor) GetBalance() int { return A.balance }

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	log.Println("Starting !")
	actor1 := &Actor{Name: "actor1"}
	actor1.Init(20)
	actor1.SetBalance(100)
	log.Println("Initial balance : ", actor1.GetBalance())
	log.Println(actor1.Name, " Ready !")

	go actor1.Run()

	for i := 1; i <= 100; i++ {
		action := BalanceAction{"plus", i} //rand.Intn(20)}
		actor1.Inbox <- action
		sleepTime := rand.Intn(5)
		time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	}

	for i := 1; i <= 100; i++ {
		action := BalanceAction{"minus", i} //rand.Intn(20)}
		actor1.Inbox <- action
		sleepTime := rand.Intn(5)
		time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	}
	time.Sleep(3 * time.Second)
	log.Println("End, balance is now ", actor1.GetBalance())
}
