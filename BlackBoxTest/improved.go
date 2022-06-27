package main

func CheckingTriangleImproved(a, b, c int) int {
	if !ValidateTriangle(a, b, c) {
		return NotATriangle
	}
	if a == b && b == c {
		return Equilateral
	} else if a != b && a != c && b != c {
		return Scalene
	}
	return Isosceles
}
