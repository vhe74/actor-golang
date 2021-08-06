package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

const CHANSIZE = 5000    //buffer before blocking
const THREADSIZE = 10000 //number of actions per thread
const NBTHREAD = 1       //concurent increasing and decreasing threads
const LOGSTOPS = 1       // log every n actions

var balance int = 100

// action payload struct
type BalanceAction struct {
	Kind   string
	Amount int
}

type Actor struct {
	Inbox chan BalanceAction
	Name  string
	wg    sync.WaitGroup
}

func (A *Actor) Init(chansize int) {
	A.Inbox = make(chan BalanceAction, chansize)
	A.wg.Add(1)
}

func (A *Actor) Run() {
	for {
		select {
		case action, ok := <-A.Inbox:
			if ok {
				if action.Kind == "plus" {
					//balance += action.Amount
					log.Println("[", A.Name, "] plus", action.Amount)
				} else if action.Kind == "minus" {
					//balance -= action.Amount
					log.Println("[", A.Name, "] minus", action.Amount)
				}
			} else {
				log.Println("[", A.Name, "] Inbox closed")
				break
			}
		}
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	log.Println("Starting !")
	actor1 := &Actor{Name: "actor1"}
	actor1.Init(20)
	log.Println(actor1.Name, " Ready !")

	go actor1.Run()

	for i := 1; i <= 10; i++ {
		action := BalanceAction{"plus", i} //rand.Intn(20)}
		actor1.Inbox <- action
		sleepTime := rand.Intn(5)
		time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	}

	time.Sleep(1 * time.Second)

	for i := 1; i <= 40; i++ {
		action := BalanceAction{"minus", i} //rand.Intn(20)}
		actor1.Inbox <- action
		sleepTime := rand.Intn(5)
		time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	}
	time.Sleep(3 * time.Second)
	//actor1.wg.Wait()
}
