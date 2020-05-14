package main

import(
	"fmt"
)

/*
	Problem: Combinations from n arrays picking one element from each array
	Refer: https://www.geeksforgeeks.org/combinations-from-n-arrays-picking-one-element-from-each-array/
*/

func findUniqueCombinations(arr [][]int){
	// number of arrays     
	n := len(arr)
	//  to keep track of next element in each of the n arrays 
	indices := make([]int, n)
	// indices is initialized with all zeros

	for true {
		// prcurrent combination 
        for i:=0; i<n; i++ {
			fmt.Print(arr[i][indices[i]], " ")
		}
		fmt.Println()

		// find the rightmost array that has more 
        // elements left after the current element 
        // in that array 
        next := n - 1
        for (next >= 0 && (indices[next] + 1 >= len(arr[next]))) {
			next-=1
		}

		// no such array is found so no more combinations left 
        if next < 0 {
			return
		}

		// if found move to next element in that array 
		indices[next] += 1
		
		// for all arrays to the right of this 
        // array current index again points to 
        // first element
		for i := next + 1; i < n; i++ {
			indices[i] = 0
		}	            
	}
}

func main(){
	arr := [][]int{ {1,2, 3}, {4}, {5, 6} }
	fmt.Println(arr)	
	findUniqueCombinations(arr)
}
