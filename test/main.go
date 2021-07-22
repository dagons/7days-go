package main

import (
	"fmt"
)

func d()(t int){
	defer func(i int){
		fmt.Println(i)
		fmt.Println(t)
	}(t)
	return 2
}

func d2()(t int){
	defer func(i int){
		fmt.Println(i)
		fmt.Println(t)
	}(t)
	t = 1
	return 1
}
func main() {
	d()
	d2()
}