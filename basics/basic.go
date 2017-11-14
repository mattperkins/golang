package main

import ("fmt" 
		"math/rand"
		"math")

func foo(){
	fmt.Println("Hello, World!")
}
func bar() {
	fmt.Println("Square root", math.Sqrt(9))
}
func baz() {
	fmt.Println("Random number:", rand.Intn(101))
}

func add(x,y,z float64) float64 {
	return x + y + z
}

func zoo(a,b string) (string, string) {
	return a, b
}

	// type conversion
	// var c int = 62
	// var d float64 = float64(c)
	// e := c // e will be type int


func main(){
	foo()
	bar()
	baz()
	// var num1 float64 = 1.5
	num1, num2, num3 := 1.5,6.3,14.5
	fmt.Println(add(num1,num2,num3))
	w1,w2 := "Hello", "World!"
	fmt.Println(zoo(w1,w2))
}


