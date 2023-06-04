package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (x MyReader) Read(m []byte) (i int,e error){
	for a:= range m{
	m[a]='A'
	}
return 1,nil
}

func main() {
	reader.Validate(MyReader{})
}
