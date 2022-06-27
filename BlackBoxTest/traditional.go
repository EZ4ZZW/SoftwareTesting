package main

import (
	"fmt"
)

func ValidateVar(name string, value int) bool {
	// Rules
	if value > 200 || value < 1 {
		fmt.Printf("Value of %s is not in the range of permitted values\n", name)
		return false
	}
	return true
}

func CheckingTriangleTraditional(a, b, c int) int {
	if !(a >= 1 && b >= 1 && c >= 1) {
		return Error
	}
	Match := 0
	if a == b {
		Match += 1
	}
	if a == c {
		Match += 2
	}
	if b == c {
		Match += 3
	}
	if Match == 0 {
		if ValidateTriangle(a, b, c) {
			return Scalene
		}
		return NotATriangle
	}
	if Match == 1 {
		if a+b <= c {
			return NotATriangle
		}
		return Isosceles
	} else if Match == 2 {
		if a+c <= b {
			return NotATriangle
		}
		return Isosceles
	} else if Match == 3 {
		if b+c <= a {
			return NotATriangle
		}
		return Isosceles
	}
	return Equilateral
}
