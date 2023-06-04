package main

import "fmt"
import "strconv"


type IPAddr [4]byte
func (x IPAddr) String() string{
	var s string=""
	s=s+strconv.Itoa(int(x[0]))+string(".")+strconv.Itoa(int(x[1]))+string(".")+strconv.Itoa(int(x[2]))+string(".")+strconv.Itoa(int(x[3]))
	return s
}
	
func main() {
	hosts:=map[string]IPAddr{
		"loopback":{127,0,0,1},
		"googleDNS":{8,8,8,8},
	}
	for name,ip:=range hosts{
		fmt.Printf("%v: %v\n",name,ip)
	}
}
