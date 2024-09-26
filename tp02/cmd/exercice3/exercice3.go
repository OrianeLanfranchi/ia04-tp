package main

import (
	"fmt"
	"time"
)

func happyNewYearSleep() {
	for i := 5; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Bonne année !!")
}

func happyNewYearAfter() {
	for i := 5; i > 0; i-- {
		fmt.Println(i)
		<-time.After(1 * time.Second)
	}
	fmt.Println("Bonne année !!")
}

func happyNewYearTick() {
	tick := time.Tick(1 * time.Second) //on crée un ticker à chaque fois qu'on appelle time.Tick donc on le définit une seule fois ici (sinon ça causait une fuite mémoire)
	for i := 5; i > 0; i-- {
		fmt.Println(i)
		<-tick
	}
	fmt.Println("Bonne année !!")
}

func main() {
	happyNewYearSleep()
	happyNewYearAfter()
	happyNewYearTick()
}
