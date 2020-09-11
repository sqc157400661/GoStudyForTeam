package main

import "fmt"

func main() {
	s := [4]int{2,7,11,19}
	t := 9
	for k, v := range s {
		for i := k+1; i < len(s); i++ {
			if v + s[i] == t {
				fmt.Printf("%d,%d", k, i)
			}
		}
	}
}
