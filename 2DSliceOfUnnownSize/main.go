package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	
	arr := [][]int{}
	
	r1 := []int{}
	r1 = append(r1, 1)
	r1 = append(r1, 2)
	r1 = append(r1, 3)
	arr = append(arr, r1)
	
	r2 := []int{}
	r2 = append(r2, 10)
	r2 = append(r2, 20)
	r2 = append(r2, 30)
	r2 = append(r2, 40)
	
	arr = append(arr, r2)
	
	fmt.Println(arr)
	
}
