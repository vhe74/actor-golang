package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

var balance int = 100
var wg sync.WaitGroup

//thread N°1
func T1() {
	for i := 1; i <= 1000; i++ {
		//fmt.Println("T1:", balance)
		balance = balance + 10
		sleepTime := rand.Intn(5)
		time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	}
	wg.Done()
}

//thread N°2
func T2() {
	for i := 1; i <= 1000; i++ {
		//fmt.Println("T2:", balance)
		balance = balance - 10
		sleepTime := rand.Intn(5)
		time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	}
	wg.Done()
}

func main() {
	log.Println("Hello, Bank!")
	log.Println("Initial balance : ", balance)
	wg.Add(2)
	go T1()
	go T2()
	wg.Wait()
	log.Println("Ending balance : ", balance)
}
