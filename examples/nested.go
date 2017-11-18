/*
	A simple program to flatten  a nested array to a single dimension array
*/
package main

import "fmt"

func main() {
	nestedArr := [][]int{{2, 4, 5, 5}, {3, 3, 52, 1},{1,2,3}} 	// A nested array with variable lenght
	ar := make([]int, 0)
	fmt.Println(array(ar, nestedArr))
}


// Todo: Make it less hacky; use either first param or the return type
func array(ar []int, arr [][]int) []int {
	l := len(arr)
	if l == 0 {
		return ar
	}
	return array(append(ar, arr[l-1]...), arr[:l-1])
}
