package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	num:=x
	for i:=0;i<10;i++ {
	num=num-(((num*num)-x)/(2*num))
	}
	return num
}
func NewSqrt(x float64) (float64,int){
	num:=x
	count:=0
	var temp float64=0
	for {
	num=num-(((num*num)-x)/(2*num))
		if math.Abs(num-temp)< 0.000001	{
			break;
		}
	temp=num
	count++
	}
	return num,count
}
func main() {
	fmt.Println(Sqrt(11))
	fmt.Println(NewSqrt(11))
}