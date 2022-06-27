package main

import (
	"testing"
)

func TestCheckingTriangleImproved(t *testing.T) {
	cases := []struct {
		Name     string
		A, B, C  int
		Expected int
	}{
		{"test 1", 100, 100, 1, Isosceles},
		{"test 2", 100, 100, 2, Isosceles},
		{"test 3", 100, 100, 100, Equilateral},
		{"test 4", 100, 100, 199, Isosceles},
		{"test 5", 100, 100, 200, NotATriangle},
		{"test 6", 100, 1, 100, Isosceles},
		{"test 7", 100, 2, 100, Isosceles},
		{"test 8", 100, 100, 100, Equilateral},
		{"test 9", 100, 199, 100, Isosceles},
		{"test10", 100, 200, 100, NotATriangle},
		{"test11", 1, 100, 100, Isosceles},
		{"test12", 2, 100, 100, Isosceles},
		{"test13", 100, 100, 100, Equilateral},
		{"test14", 199, 100, 100, Isosceles},
		{"test15", 200, 100, 100, NotATriangle},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if ans := CheckingTriangleImproved(c.A, c.B, c.C); ans != c.Expected {
				t.Fatalf("Triangle %d %d %d expected %d, but %d got",
					c.A, c.B, c.C, c.Expected, ans)
			}
		})
	}
}

func TestImprovedBoundary(t *testing.T) {
	cases := []struct {
		Name     string
		A, B, C  int
		Expected int
	}{
		{"test 1", 100, 100, 1, Isosceles},
		{"test 2", 100, 100, 2, Isosceles},
		{"test 3", 100, 100, 100, Equilateral},
		{"test 4", 100, 100, 199, Isosceles},
		{"test 5", 100, 100, 200, NotATriangle},
		{"test 6", 100, 1, 100, Isosceles},
		{"test 7", 100, 2, 100, Isosceles},
		{"test 8", 100, 100, 100, Equilateral},
		{"test 9", 100, 199, 100, Isosceles},
		{"test10", 100, 200, 100, NotATriangle},
		{"test11", 1, 100, 100, Isosceles},
		{"test12", 2, 100, 100, Isosceles},
		{"test13", 100, 100, 100, Equilateral},
		{"test14", 199, 100, 100, Isosceles},
		{"test15", 200, 100, 100, NotATriangle},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if ans := CheckingTriangleImproved(c.A, c.B, c.C); ans != c.Expected {
				t.Fatalf("Triangle %d %d %d expected %d, but %d got",
					c.A, c.B, c.C, c.Expected, ans)
			}
		})
	}
}
