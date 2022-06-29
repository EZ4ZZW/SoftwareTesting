package main

import "fmt"

// {1, 2, 199, 200, 100}^3

func script() {
	numbers := []int{1, 2, 100, 199, 200}
	id := 16
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			for k := 0; k < 5; k++ {
				res := CheckingTriangleImproved(numbers[i], numbers[j], numbers[k])
				fmt.Printf("{\"test %d\", %d, %d, %d, %s},\n", id, numbers[i], numbers[j], numbers[k], KindOfTri[res])
				id++
			}
		}
	}
}
