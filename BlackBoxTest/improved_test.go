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
		{"test 16", 1, 1, 1, Equilateral},
		{"test 17", 1, 1, 2, NotATriangle},
		{"test 18", 1, 1, 100, NotATriangle},
		{"test 19", 1, 1, 199, NotATriangle},
		{"test 20", 1, 1, 200, NotATriangle},
		{"test 21", 1, 2, 1, NotATriangle},
		{"test 22", 1, 2, 2, Isosceles},
		{"test 23", 1, 2, 100, NotATriangle},
		{"test 24", 1, 2, 199, NotATriangle},
		{"test 25", 1, 2, 200, NotATriangle},
		{"test 26", 1, 100, 1, NotATriangle},
		{"test 27", 1, 100, 2, NotATriangle},
		{"test 28", 1, 100, 100, Isosceles},
		{"test 29", 1, 100, 199, NotATriangle},
		{"test 30", 1, 100, 200, NotATriangle},
		{"test 31", 1, 199, 1, NotATriangle},
		{"test 32", 1, 199, 2, NotATriangle},
		{"test 33", 1, 199, 100, NotATriangle},
		{"test 34", 1, 199, 199, Isosceles},
		{"test 35", 1, 199, 200, NotATriangle},
		{"test 36", 1, 200, 1, NotATriangle},
		{"test 37", 1, 200, 2, NotATriangle},
		{"test 38", 1, 200, 100, NotATriangle},
		{"test 39", 1, 200, 199, NotATriangle},
		{"test 40", 1, 200, 200, Isosceles},
		{"test 41", 2, 1, 1, NotATriangle},
		{"test 42", 2, 1, 2, Isosceles},
		{"test 43", 2, 1, 100, NotATriangle},
		{"test 44", 2, 1, 199, NotATriangle},
		{"test 45", 2, 1, 200, NotATriangle},
		{"test 46", 2, 2, 1, Isosceles},
		{"test 47", 2, 2, 2, Equilateral},
		{"test 48", 2, 2, 100, NotATriangle},
		{"test 49", 2, 2, 199, NotATriangle},
		{"test 50", 2, 2, 200, NotATriangle},
		{"test 51", 2, 100, 1, NotATriangle},
		{"test 52", 2, 100, 2, NotATriangle},
		{"test 53", 2, 100, 100, Isosceles},
		{"test 54", 2, 100, 199, NotATriangle},
		{"test 55", 2, 100, 200, NotATriangle},
		{"test 56", 2, 199, 1, NotATriangle},
		{"test 57", 2, 199, 2, NotATriangle},
		{"test 58", 2, 199, 100, NotATriangle},
		{"test 59", 2, 199, 199, Isosceles},
		{"test 60", 2, 199, 200, Scalene},
		{"test 61", 2, 200, 1, NotATriangle},
		{"test 62", 2, 200, 2, NotATriangle},
		{"test 63", 2, 200, 100, NotATriangle},
		{"test 64", 2, 200, 199, Scalene},
		{"test 65", 2, 200, 200, Isosceles},
		{"test 66", 100, 1, 1, NotATriangle},
		{"test 67", 100, 1, 2, NotATriangle},
		{"test 68", 100, 1, 100, Isosceles},
		{"test 69", 100, 1, 199, NotATriangle},
		{"test 70", 100, 1, 200, NotATriangle},
		{"test 71", 100, 2, 1, NotATriangle},
		{"test 72", 100, 2, 2, NotATriangle},
		{"test 73", 100, 2, 100, Isosceles},
		{"test 74", 100, 2, 199, NotATriangle},
		{"test 75", 100, 2, 200, NotATriangle},
		{"test 76", 100, 100, 1, Isosceles},
		{"test 77", 100, 100, 2, Isosceles},
		{"test 78", 100, 100, 100, Equilateral},
		{"test 79", 100, 100, 199, Isosceles},
		{"test 80", 100, 100, 200, NotATriangle},
		{"test 81", 100, 199, 1, NotATriangle},
		{"test 82", 100, 199, 2, NotATriangle},
		{"test 83", 100, 199, 100, Isosceles},
		{"test 84", 100, 199, 199, Isosceles},
		{"test 85", 100, 199, 200, Scalene},
		{"test 86", 100, 200, 1, NotATriangle},
		{"test 87", 100, 200, 2, NotATriangle},
		{"test 88", 100, 200, 100, NotATriangle},
		{"test 89", 100, 200, 199, Scalene},
		{"test 90", 100, 200, 200, Isosceles},
		{"test 91", 199, 1, 1, NotATriangle},
		{"test 92", 199, 1, 2, NotATriangle},
		{"test 93", 199, 1, 100, NotATriangle},
		{"test 94", 199, 1, 199, Isosceles},
		{"test 95", 199, 1, 200, NotATriangle},
		{"test 96", 199, 2, 1, NotATriangle},
		{"test 97", 199, 2, 2, NotATriangle},
		{"test 98", 199, 2, 100, NotATriangle},
		{"test 99", 199, 2, 199, Isosceles},
		{"test 100", 199, 2, 200, Scalene},
		{"test 101", 199, 100, 1, NotATriangle},
		{"test 102", 199, 100, 2, NotATriangle},
		{"test 103", 199, 100, 100, Isosceles},
		{"test 104", 199, 100, 199, Isosceles},
		{"test 105", 199, 100, 200, Scalene},
		{"test 106", 199, 199, 1, Isosceles},
		{"test 107", 199, 199, 2, Isosceles},
		{"test 108", 199, 199, 100, Isosceles},
		{"test 109", 199, 199, 199, Equilateral},
		{"test 110", 199, 199, 200, Isosceles},
		{"test 111", 199, 200, 1, NotATriangle},
		{"test 112", 199, 200, 2, Scalene},
		{"test 113", 199, 200, 100, Scalene},
		{"test 114", 199, 200, 199, Isosceles},
		{"test 115", 199, 200, 200, Isosceles},
		{"test 116", 200, 1, 1, NotATriangle},
		{"test 117", 200, 1, 2, NotATriangle},
		{"test 118", 200, 1, 100, NotATriangle},
		{"test 119", 200, 1, 199, NotATriangle},
		{"test 120", 200, 1, 200, Isosceles},
		{"test 121", 200, 2, 1, NotATriangle},
		{"test 122", 200, 2, 2, NotATriangle},
		{"test 123", 200, 2, 100, NotATriangle},
		{"test 124", 200, 2, 199, Scalene},
		{"test 125", 200, 2, 200, Isosceles},
		{"test 126", 200, 100, 1, NotATriangle},
		{"test 127", 200, 100, 2, NotATriangle},
		{"test 128", 200, 100, 100, NotATriangle},
		{"test 129", 200, 100, 199, Scalene},
		{"test 130", 200, 100, 200, Isosceles},
		{"test 131", 200, 199, 1, NotATriangle},
		{"test 132", 200, 199, 2, Scalene},
		{"test 133", 200, 199, 100, Scalene},
		{"test 134", 200, 199, 199, Isosceles},
		{"test 135", 200, 199, 200, Isosceles},
		{"test 136", 200, 200, 1, Isosceles},
		{"test 137", 200, 200, 2, Isosceles},
		{"test 138", 200, 200, 100, Isosceles},
		{"test 139", 200, 200, 199, Isosceles},
		{"test 140", 200, 200, 200, Equilateral},
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
