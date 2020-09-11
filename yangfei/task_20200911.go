package main

import (
	"fmt"
	"time"
)

func main() {
		fmt.Println(reverse_slice([]int{1, 2, 3, 4}))
		
        slice1 := []string{"a", "b", "c", "d", "e"}
        slice2 := []string{"1", "2", "3", "4", "5"}
		fmt.Println(InsertStringSlice(2, slice1, slice2))
		
		for i := 0; i < 10; i++ {
			go printHelloWord(i)
		}
		time.Sleep(1e9)
		var arr []int = []int{1,2,3,4,5,6,7}
		//return_slice := []int
		searchTargetGroup(&arr,5)
		//fmt.Println(return_slice)
}

// 逆序切片
func reverse_slice(s []int) []int {
        for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
                s[i], s[j] = s[j], s[i]
        }
        return s
}

// 将一个切片slice1插入另一个切片clice2的指定位置
func InsertStringSlice(i int, slice1 []string, slice2 []string) []string {
        var new_clice []string
        new_clice = append(new_clice, slice2[:i]...)
        new_clice = append(new_clice, slice1...)
        new_clice = append(new_clice, slice2[i:]...)
        return new_clice
}

// 使用协程打印10个HelloWorld
func printHelloWord(i int)  {
	fmt.Printf("HelloWorld_%d\n",i)
}

// 寻找目标组合
func searchTargetGroup(arr *[]int, target int){
	var len = len(*arr)
	if len>0 {
		for i := 0; i < len-1; i++ {
			for j := 0; j < i; j++ {
				if i+j==target {
					fmt.Printf("%d,%d\n",i,j)
				}
			}
		}
	}
	
}