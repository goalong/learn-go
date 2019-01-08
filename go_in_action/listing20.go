package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)


var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}


func main() {
	court := make(chan int)
	wg.Add(2)
	go player("Nadal", court)
}


func player(name string, court chan int) {
	defer wg.Done()
	for {
		ball, ok := <-court
		if !ok {
			fmt.Println("")
		}
	}
}
