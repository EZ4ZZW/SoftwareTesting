package main

import "fmt"

const (
	Equilateral int = iota
	Isosceles
	Scalene
	NotATriangle
	Error
)

var KindOfTri map[int]string = map[int]string{
	0: "Equilateral",
	1: "Isosceles",
	2: "Scalene",
	3: "Not A Triangle",
}

var (
	a int
	b int
	c int
)

func ValidateTriangle(a, b, c int) bool {
	return a+b > c && a+c > b && b+c > a
}

func main() {
	// loop get edge
	//	for {
	//		num, err := fmt.Scanf("%d%d%d", &a, &b, &c)
	//		if err != nil {
	//			fmt.Println("Input Error")
	//			continue
	//		}
	//		if num != 3 {
	//			fmt.Println("Not Permit Input")
	//			continue
	//		}
	//		if !(ValidateVar("a", a) && ValidateVar("b", b) && ValidateVar("c", c)) {
	//			continue
	//		}
	//		break
	//	}
	fmt.Scanf("%d%d%d", &a, &b, &c)
	fmt.Println(a, b, c)
	fmt.Println(KindOfTri[CheckingTriangleTraditional(a, b, c)])
}
