package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
x,y:=0,1
count:=0
    return func() int{
	if count==0 {
	count++
	return x
	}
	if count==1 {
	count++
	return y
	}
	}
	y = y(count-1) + y(count-2)
	return y	
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
