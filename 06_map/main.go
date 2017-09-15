package main

import(
	"fmt"
)

func main(){
	dict :=make(map[string]string)
	dict["go"]="Golang"
	dict["cs"]="Csharp"
	dict["rb"]="Ruby"
	dict["py"]="Python"
	dict["js"]="Javascript"
	for k,v :=range dict{
		fmt.Printf("Key: %s  Value: %s\n",k,v)
	}
	if lan,ok := dict["go"]; ok{
		fmt.Println(lan,ok)
	}
}