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

// content should be "plus" or "minus" and an amount
var mailbox = make(chan BalanceAction, CHANSIZE)

//to avoid anticipated exit
var wg sync.WaitGroup

// the balance Manager
func balanceManager(c chan BalanceAction) {
	opsNo := 0
	for {
		select {
		case action, ok := <-c:
			if ok {
				opsNo++
				//log.Println(action)
				if opsNo%LOGSTOPS == 0 {
					log.Println("[", opsNo, "]", "Action asked is :  ", action, " starting balance : ", balance)
				}

				if action.Kind == "plus" {
					balance += action.Amount
					//log.Printf("The new balance is : %d\n", balance)
				} else if action.Kind == "minus" {
					balance -= action.Amount
					//log.Printf("The new balance is : %d\n", balance)
				}
			} else {
				log.Println("Mailbox closed!")
			}
		default:
			//log.Println("No value ready, moving on. opsNo :", opsNo, " balance :", balance)
		}
	}
}

func Tplus() {
	for i := 1; i <= THREADSIZE; i++ {
		action := BalanceAction{"plus", rand.Intn(20)}
		mailbox <- action
		sleepTime := rand.Intn(5)
		time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	}
	wg.Done()
}

func Tminus() {
	for i := 1; i <= THREADSIZE; i++ {
		mailbox <- BalanceAction{"minus", rand.Intn(20)}
		sleepTime := rand.Intn(5)
		time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	}
	wg.Done()
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("Starting Buffered Chan sample")

	for i := 0; i < NBTHREAD; i++ {
		wg.Add(2)
		go Tplus()
		go Tminus()
	}
	go balanceManager(mailbox)
	wg.Wait()
	log.Printf("current balance is : %d\n", balance)
	time.Sleep(2 * time.Second)
	log.Println("Ending balance after : ", balance)
}
