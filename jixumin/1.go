package main

import "fmt"

func main() {
	a := []int{1,2,3,4,5}
	BackPrint(a)
}

func BackPrint(a []int) {
	for i:=len(a); i>0; i-- {
		fmt.Println(a[i-1])
	}
}
