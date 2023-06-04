package main

import "golang.org/x/tour/tree"
import "fmt"
// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
};
func Same(t1, t2 *tree.Tree) bool {
var x,y [10]int	
ch:= make(chan int,10)
ch2 := make(chan int,10)
go Walk(t1, ch)
go Walk(t2, ch2)
for i := 0; i < 10; i++ {
		x[i]=<-ch
		y[i]=<-ch2
	}
	return x==y
}

func main() {

fmt.Println(Same(tree.New(1), tree.New(1)))
}
