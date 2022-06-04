package main

import (
	"fmt"
	"math/rand"
)

var dadu = []int{1, 2, 3, 4, 5, 6}

func Input(N, M int) []int {
	var result []int
	var player []int
	var dadu []int

	for i := 1; i <= M; i++ {
		player = append(player, i)
	}
	for j := 1; j <= N; j++ {
		dadu = append(dadu, j)
	}

	l := dadu[rand.Intn(len(dadu))]
	for k := 1; k <= l; k++ {
		result = append(result, k)
	}
	return result
}

func main() {
	fmt.Println(Input(3, 4))
}
