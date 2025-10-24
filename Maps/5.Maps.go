package main

import "fmt"

func main(){
	m:= make(map[string]string)	// key value pair

	m["akshat"] = "Kumar"

	fmt.Println(m["akshat"])
}