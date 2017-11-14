package main

import ("fmt"
"time"
"sync")

var wg sync.WaitGroup

func cleanup(){
	defer wg.Done()
	if r := recover(); r != nil {
		fmt.Println("Recovered in cleanup    ", r)
	}
}

func say(s string){
	defer cleanup()
	for i := 0; i < 3; i++ {
		time.Sleep(time.Millisecond*100)
		fmt.Println(s)
		if i == 2 {
			panic("Panic!!!")
		}	
	}
		
}

func main(){
	wg.Add(1)
	go say("Go!")
	wg.Add(1)
	go say("Go Again!")
	wg.Wait()
}	 