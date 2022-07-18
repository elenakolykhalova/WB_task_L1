// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

package main

import "fmt"

type Human struct {
	Name 	string
	Age		int
	Weight	int
}

type Action struct {
	Human
	run		bool	
}

func (h Human) Run() bool {
	if h.Weight > 65 {
		return true
	} else {
		return false
	}
}

func main() {
	a := Action{Human{"Nikolay", 26, 45}, false}
	fmt.Println(a.Run())
}
