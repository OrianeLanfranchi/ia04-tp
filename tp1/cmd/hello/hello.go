package main

import (
	"fmt"
)

func isPair(x int) bool {
	return (x % 2) == 0
}

func main() {
	for i := 0; i < 1001; i++ {
		if isPair(i) {
			fmt.Println(i)
		}
	}

}
