/*
	An exercise on Binary Tree to check if the underlying values in the given two trees are true or not.
	Based on https://tour.golang.org/concurrency/8
*/

package main

import "fmt"
import "golang.org/x/tour/tree"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	ch <- t.Value
	if t.Left != nil {
		Walk(t.Left,ch)
	}
	if t.Right != nil {
		Walk(t.Right,ch)
	}
	if t.Left ==nil && t.Right == nil {
		return
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1:= make(chan int)
	m := make(map[int]bool)
	go Walk(t1,c1)
	go Walk(t2,c1)
	for i:=0;i<10;i++ {
			m[<-c1] = true
	}
	for i:=0;i<10;i++ {
		if _,ok := m[<-c1]; !ok {
			return false
		}
	}

	return true
}

func main() {
	t1 := tree.New(1)
	t2 := tree.New(2)
	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(Same(t1,t2))
}








