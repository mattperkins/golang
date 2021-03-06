package main

import ("fmt"
"time"
"sync")

var wg sync.WaitGroup

func say(s string){
	defer wg.Done()
	for i:=0; i<3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond*100)
	}
	// wg.Done() replaced by 'defer statement' above
}

func main(){
	wg.Add(1)
	go say("Go!")
	wg.Wait()
	wg.Add(1)
	go say("Wait!")
	wg.Wait()
	wg.Add(1)
	go say("Go Again!")
	wg.Wait()
}	 