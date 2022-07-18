//Разработать программу, которая в рантайме способна определить тип переменной: 
// int, string, bool, channel из переменной типа interface{}.

package main

import "fmt"

func WhoIsType(a interface{}){
	switch a.(type) {
	case bool:
		fmt.Println("This is bool")
	case string:
		fmt.Println("This is string")
	case int:
		fmt.Println("This is int")
	default:
		fmt.Println("I don't know what that type is")
	}
}

func main(){
	WhoIsType(false)
	WhoIsType(10)
	WhoIsType("Hello")
	WhoIsType(23.45)
}
