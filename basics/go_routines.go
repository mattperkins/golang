package main

import ("fmt"
"time")

func say(s string){
	for i:=0; i<3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond*500)
	}
}

func main(){
	
	say("Hello")
	time.Sleep(time.Second)
	go say("Hello") // go routine running concurrently
	say("World!")
}
