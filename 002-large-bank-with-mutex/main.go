package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

var balance int = 100
var wg sync.WaitGroup
var lock = sync.Mutex{}

const nbThread = 4

//thread N°1
func Tplus() {
	for i := 1; i <= 1000; i++ {
		//fmt.Println("T1:", balance)
		lock.Lock()
		balance = balance + 10
		//lock.Unlock() //releasing lock immediatly
		sleepTime := rand.Intn(5)
		time.Sleep(time.Duration(sleepTime) * time.Millisecond)
		lock.Unlock() //releasing Lock after some computation spending up to 5ms
	}
	wg.Done()
}

//thread N°2
func Tminus() {
	for i := 1; i <= 1000; i++ {
		//fmt.Println("T1:", balance)
		lock.Lock()
		balance = balance - 10
		//lock.Unlock() //releasing lock immediatly
		sleepTime := rand.Intn(5)
		time.Sleep(time.Duration(sleepTime) * time.Millisecond)
		lock.Unlock() //releasing Lock after some computation spending up to 5ms
	}
	wg.Done()
}

func main() {
	log.Println("Hello, Bank!")
	log.Println("Initial balance : ", balance)
	for i := 0; i < nbThread; i++ {
		wg.Add(1)
		go Tplus()
		go Tminus()
	}
	wg.Wait()
	log.Println("Ending balance : ", balance)
}
