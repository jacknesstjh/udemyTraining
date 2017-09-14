package main

import(
	"strings"
)

func main(){
	s := "I am a string\nContaining new lines"
	s = strings.Replace(s,"\n","<br>",-1)
	println(s)
}