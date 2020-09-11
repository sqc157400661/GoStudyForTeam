package main

import "fmt"

func main() {
	a := []int{1,2,3}
	b := []int{4,5,6}
	fmt.Println(insertSlice(a, b, 1))
		
}

func insertSlice(a, b []int, i int) []int {
	return append(append(a[:i], b...), a[i:len(a)]...)
}
