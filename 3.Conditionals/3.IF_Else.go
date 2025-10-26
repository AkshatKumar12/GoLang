package main

import "fmt"

func main(){
	age:=17

	if age>=18{
		fmt.Println("You can vote")
		} else if age==17{
			fmt.Println("Congrats 1 year to go")
		}else{
			fmt.Print("You cant vote")
		}
}