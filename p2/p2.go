package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
   //split the individual words 
	split:=strings.Fields(s)
	m:=make(map[string]int)
	//iterate for loop unitl the end of string
	for _,word:=range split{
	m[word]++ // increment the value of every matching key
	}
	w:=m["Word"]
	fmt.Printf("\n******************The instance \"Word\" was found %d times*******************\n",w)
	return m
}

func main() {
	wc.Test(WordCount)
}
