package main

import (
	"fmt"
	"sync"
)

var l sync.Mutex // 0 value = Unlock status
var n = 0

func f() {

	n++
}

func f2(c chan int) {
	n += 1
	c <- n
}

func main() {
	c := make(chan int)

	for i := 0; i < 10000; i++ {
		go f2(c)
		_ = <-c
	}

	fmt.Println("Appuyez sur entrée")
	fmt.Scanln()
	fmt.Println("n:", n)
	close(c)
}
