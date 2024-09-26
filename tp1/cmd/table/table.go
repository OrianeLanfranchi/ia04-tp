package main

import (
	"fmt"
	"tp1/pkg/table"
)

func main() {
	var sl = make([]int, 11)
	table.Fill(sl) //modifies sl
	table.Compte(sl)
	fmt.Println(table.ValeursCentrales(sl))
	sl1 := append([]int{}, sl...)  //copies sl content in sl1
	fmt.Println("SL copy : ", sl1) //test

	//copy :
	var sl2 []int //slice pas initialisée
	copy(sl2, sl) //pas de réallocation de mémoires
}
