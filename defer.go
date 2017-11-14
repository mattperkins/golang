package main

import "fmt"

func foo(){
	// defer statements are first in last out/last in first out
	defer fmt.Println("I am first but will run last")
	defer fmt.Println("I am second in the list")
	fmt.Println("I am last but will run first")
}

func bar(){
	for i:=0; i<3; i++{
		defer fmt.Println(i)
	}
}

func main(){
	foo()
	bar()
}