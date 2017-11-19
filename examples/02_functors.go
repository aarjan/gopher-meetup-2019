/*
	A sample function that works as a closure, and maps to each underlying type as in the given function.
*/

package main

import "fmt"

type functor func(int)int	// functor is a type that accepts a function with defination; func (int) int

func magical(slice ...int) func(functor) []int {
	return func(f functor) []int {
		result := make([]int, len(slice))
		for i, n := range slice {
			result[i] = f(n)
		}
		return result
	}
}

func main() {
	fun := magical(1, 2, 3, 4)	// 'magical()' returns a function; func(functor) []int
	afunctor := func(a int) int {	// A a function defination of type 'functor'
		return a * a * a
	}
	fmt.Println(fun(afunctor))	// Since 'fun' as a function accepts a 'functor' type

}
