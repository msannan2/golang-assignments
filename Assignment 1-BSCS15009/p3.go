package main

import "fmt"

func fibonacci() func() int {
  var a, b, count int = 0, 1, 0
    return func() int {
		for count==0{
        count++
        return a 
		}
		temp:=a
		a=b
		b=temp+b
        return a
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
