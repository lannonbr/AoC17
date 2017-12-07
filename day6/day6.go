package main

import (
	"fmt"
)

func getMax(arr []int) int {
	maxIdx := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] > arr[maxIdx] {
			maxIdx = i
		}
	}

	return maxIdx
}

func refill(input []int, idx int) []int {
	num := input[idx]
	input[idx] = 0
	idx++
	for num > 0 {
		input[idx%len(input)]++
		idx++
		num--
	}
	return input
}

func main() {
	input := []int{4, 10, 4, 1, 8, 4, 9, 14, 5, 1, 14, 15, 0, 15, 3, 5}

	// input := []int{0, 2, 7, 0}

	instances := make(map[string]int)

	cycles := 0

	for {
		idx := getMax(input)

		input = refill(input, idx)

		cycles++
		inputStr := fmt.Sprintf("%v", input)

		if v, ok := instances[inputStr]; ok {
			fmt.Println("[Pt2] This many cycles:", v)
			break
		} else {
			instances[inputStr] = 0
			for k := range instances {
				instances[k]++
			}
		}
	}

	fmt.Println("[Pt1] Cycles:", cycles)
}
