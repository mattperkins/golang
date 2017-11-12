package main

import "fmt"

type car struct {
	accelerator uint16 // min 0 max 65535
	brake uint16
	steering_wheel uint16
	top_speed float64 }

func main () {
	// fmt.Println("Test...")
	a_car := car {
		accelerator: 22341,
		brake: 0,
		steering_wheel: 12561,
		top_speed: 120 }
	// shorthand but less descriptive
	// a_car := car{22341,0,12561,120} 

	fmt.Println(a_car.steering_wheel)
}


