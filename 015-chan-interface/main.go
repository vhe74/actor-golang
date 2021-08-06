package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan interface{}, 10)

	go func() {
		for {
			select {
			case p := <-ch:
				//fmt.Println("Received a ", reflect.TypeOf(p).Name(), " : ", p)
				switch p := p.(type) {
				case string:
					fmt.Printf("Got a string %q\n", p)
				case int:
					fmt.Printf("Got a int %d\n", p)
				default:
					fmt.Printf("Type of p is %T. Value %v\n", p, p)
				}
			}
		}
	}()

	ch <- "Vincent"
	ch <- 31
	ch <- true

	time.Sleep(3 * time.Second)
}