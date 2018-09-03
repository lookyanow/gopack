package main

import "fmt"

func filter(numbers []int, callback func(int) bool) ([]int, []int) {
	xsEven := []int{}
	xsOdd := []int{}

	for _, n := range numbers {
		if callback(n) {
			xsEven = append(xsEven, n)
		} else {
			xsOdd = append(xsOdd, n)
		}

	}
	return xsEven, xsOdd
}

func checkNum(n int) bool {
	return n%2 == 0
}

func main() {
	xs1, xs2 := filter([]int{1, 2, 3, 4, 5, 6, 7, 8}, checkNum)
	fmt.Println(xs1)
	fmt.Println(xs2)
	fmt.Println(len(xs1))
	fmt.Println(len(xs2))
}
