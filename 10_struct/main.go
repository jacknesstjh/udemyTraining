package main

import(
	"fmt"
)

type person struct{
	name string
	age int
}

var p person
func main(){
	p.name="jackness"
	p.age=32
	fmt.Println("The person's name is ",p.name)
}


