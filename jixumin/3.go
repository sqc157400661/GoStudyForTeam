package main

import "fmt"

func main() {
	c := make(chan bool, 10)
	for i:=0; i<10; i++ {
		go func(){
			fmt.Printf("hello world\n")
			c <- true
		}()	
	}
	
	for j:=0; j<10; j++ {
		<- c
	}
}
