package main

import (
	"fmt"
	"sync"
	"time"
)

func subRout(c chan int) {
	c <- 11
}

func main() {
	c1 := make(chan int)
	//c2 := make(chan int)

	go subRout(c1)
	//go subRout(c2)

	// blocks the gouroutine until data is received on channels unless "default is supplied"
	select {

	case l, ok := <-c1:
		if ok {
			fmt.Println(l)
		}

	case l, ok := <-time.After(time.Second * 3):
		if ok {
			fmt.Println("Timeout broo", l)
		}
	}

	// sync routines - wait for them
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)

	i := 0
	for i < 2 {
		go func(wg *sync.WaitGroup) {
			fmt.Println("Routine processing..")
			waitGroup.Done()
		}(&waitGroup)
		i++
	}

	waitGroup.Wait()

	fmt.Println("Done!")

	//fmt.Println(l)
}
