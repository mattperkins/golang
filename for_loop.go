package main

import "fmt"

func main() {

	for j:=0; j<3; j++ {
		fmt.Println("A standard For loop")
	}

	i:=0 
	for i<10 {
		fmt.Println(i)
		// increment in steps of 3
		i+=3
	}

	x:=5
	for {
		 fmt.Println("Do stuff", x)
		 x+=3
		 if x > 25{
			 break
		 }
	 }
	 
	for x:=5; x<25; x+=3{
		 fmt.Println("Alternate syntax to 'Do stuff loop above'",x)}

	a:=3
	for x:=5; a<25; x+=3{
		fmt.Println("Mixed variable loop", x)
		a+=4
	}
}