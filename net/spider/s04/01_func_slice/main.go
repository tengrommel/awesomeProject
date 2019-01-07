package main

import "fmt"

func main() {
	scores1 := []float32{91, 82, 99}
	fmt.Printf("average: %.2f\n", avg(scores1))
	fmt.Printf("average: %.2f\n", avgList(scores1...))
	fmt.Printf("average: %.2f\n", avgList(float32(9), float32(3), float32(5)))
	fmt.Printf("average: %.2f\n", avgList())
}

func avgList(list ...float32) float32 {
	if len(list) == 0 {
		return float32(0)
	}
	var total = float32(0)
	for i := range list {
		total += list[i]
	}
	return total / float32(len(list))
}

func avg(scores []float32) float32 {
	var total float32
	for _, score := range scores {
		total += score
	}
	return total / float32(len(scores))
}
