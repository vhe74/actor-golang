package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

const CHANSIZE = 5000   //buffer before blocking
const THREADSIZE = 1000 //number of actions per thread
const NBTHREAD = 20     //concurent increasing and decreasing threads
const LOGSTOPS = 500    // log every n actions

var balance int = 100

// content should be "plus" or "minus"
var mailbox = make(chan string, CHANSIZE)

//to avoid anticipated exit
var wg sync.WaitGroup

// the balance Manager
func balanceManager(c chan string) {
	opsNo := 0
	for {
		select {
		case action, ok := <-c:
			if ok {
				opsNo++
				if opsNo%LOGSTOPS == 0 {
					log.Println("[", opsNo, "]", "Action asked is :  ", action, " balance : ", balance)
				}
				if action == "plus" {
					balance += 10
					//log.Printf("The new balance is : %d\n", balance)
				} else if action == "minus" {
					balance -= 10
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
		mailbox <- "plus" //send "plus" to the channel
		sleepTime := rand.Intn(5)
		time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	}
	wg.Done()
}

func Tminus() {
	for i := 1; i <= THREADSIZE; i++ {
		mailbox <- "minus" //send "minus" to the channel
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
