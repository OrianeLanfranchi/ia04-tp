package table

import (
	"fmt"
	"math/rand"
	"sort"
)

func Fill(sl []int) {
	for i := range sl {
		sl[i] = rand.Intn(100-1) + 1
	}
}

func Moyenne(sl []int) float64 {
	var count float64

	for _, e := range sl {
		count += float64(e)
	}
	return count / float64(len(sl))
}

func ValeursCentrales(sl []int) []int {
	var sl2 = make([]int, len(sl))
	copy(sl2, sl)
	sort.Ints(sl2)
	if (len(sl2) % 2) == 0 {
		return []int{sl2[(len(sl2)/2)-1], sl[(len(sl2) / 2)]}
	}
	return []int{sl2[len(sl2)/2]}
}

func Plus1(sl []int) {
	for i := range sl {
		sl[i] += 1
	}
}

func Compte(sl []int) {
	fmt.Println(sl)
}
