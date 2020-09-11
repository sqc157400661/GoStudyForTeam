package main

import "fmt"

func main()  {
	fmt.Println(reverse_slice([]string{"1","2","3","4"}))
}

func reverse_slice(s []string) []string  {
	for i,j := 0,len(s)-1; i < j; i,j=i+1,j-1 {
			s[i],s[j] = s[j],s[i]
	}
	return s
}

