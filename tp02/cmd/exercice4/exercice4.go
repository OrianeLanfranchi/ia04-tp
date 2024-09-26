package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// remplit tab avec la valeur v
func Fill(tab []int, v int) {
	for i := range tab {
		tab[i] = v
	}
}

func FillGoroutine(tab []int, v int, n int) {
	var wg sync.WaitGroup

	step := len(tab) / n

	for i := 0; i < n-1; i++ {
		wg.Add(1)
		sl := tab[i*step : (i+1)*step]
		go func() {
			defer wg.Done()
			Fill(sl, v)
		}()
	}

	sl := tab[(n-1)*step:]
	wg.Add(1)
	go func() {
		defer wg.Done()
		Fill(sl, v)
	}()

	wg.Wait()
}

// applique f sur chaque élément de tab et remplace la valeur initiale par le résultat de f
func ForEach(tab []int, f func(int) int) {
	for i := range tab {
		tab[i] = f(tab[i])
	}
}

func ForEachGoroutine(tab []int, f func(int) int, n int) {
	var wg sync.WaitGroup
	step := len(tab) / n

	for i := 0; i < n-1; i++ {
		wg.Add(1)
		sl := tab[i*step : (i+1)*step]
		go func() {
			defer wg.Done()
			ForEach(sl, f)
		}()
	}

	sl := tab[(n-1)*step:]
	wg.Add(1)
	go func() {
		defer wg.Done()
		ForEach(sl, f)
	}()

	wg.Wait()

	wg.Wait()
}

// copy le tableau src dans dest
func Copy(src []int, dest []int) {
	copy(dest, src)
}

// vérifie que tab1 et tab2 sont identiques
func Equal(tab1 []int, tab2 []int) bool {
	for i := range tab1 {
		if tab1[i] != tab2[i] {
			fmt.Println(i)
			fmt.Println((tab1[i]))
			fmt.Println((tab2[i]))
			return false
		}
	}
	return true
}

func addition(n int) int {
	n++
	return n
}

func main() {
	fmt.Println(runtime.NumCPU())

	//valeur de la constante définie à la compilation -> pas de variable dans les valeurs
	const size int = 100 * (1 << 20)
	const goNb = 500
	var tab1 = make([]int, size)
	var tab2 = make([]int, size)

	start := time.Now()
	Fill(tab1, 1)
	ForEach(tab1, addition)
	elapsed := time.Since(start)
	fmt.Println("Function : ", elapsed)

	start = time.Now()
	FillGoroutine(tab2, 1, goNb)
	ForEachGoroutine(tab2, addition, goNb)
	elapsed = time.Since(start)
	fmt.Println("Goroutine : ", elapsed)

	fmt.Println("tab1 == tab2 :", Equal(tab1, tab2))
}
