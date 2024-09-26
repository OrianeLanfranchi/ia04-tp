package main

import (
	"fmt"
)

func compte(n int) {
	for i := range n {
		fmt.Println(i)
	}
}

func compteMsg(n int, msg string) {
	for i := range n {
		fmt.Println(msg)
		fmt.Println(i)
	}
}

func compteMsgFromTo(start int, end int, msg string) {
	for i := start; i <= end; i++ {
		fmt.Println(msg)
		fmt.Println(i)
	}
}

func main() {
	//go compte(10)

	//go compteMsg(100, "Goroutine 1")
	//go compteMsg(100, "Goroutine 2")

	/*var count int

	for i := 0; i <= 91; i += 10 {
		go compteMsgFromTo(i, i+10, "goroutine "+strconv.Itoa(count))
		count++
	}*/

	fmt.Scanln()
}
