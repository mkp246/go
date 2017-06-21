package main

import "fmt"

func main() {
	m:=make(map[string]int)
	m["ans"]=43
	fmt.Println("value:",m["ans"])
	m["ans"]=49
	fmt.Println("value:",m["ans"])
	delete(m,"ans")
	fmt.Println("value:",m["ans"])
	v,ok:=m["ans"]
	fmt.Println("value:",v,"present:",ok)
}