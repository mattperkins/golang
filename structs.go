package main

import "fmt"

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

type car struct {
	accelerator uint16 // min 0 max 65535
	brake uint16
	steering_wheel uint16
	top_speed_kmh float64 }

// pointer receiver
func (c *car) kmh() float64 {
	c.top_speed_kmh = 500
	return float64(c.accelerator) * (c.top_speed_kmh/usixteenbitmax)
}

// value receiver
func (c car) mph() float64 {
	return float64(c.accelerator) * (c.top_speed_kmh/usixteenbitmax/kmh_multiple)
}

func (c *car) new_top_speed(newspeed float64){
	c.top_speed_kmh = newspeed
}

func main () {
	// fmt.Println("Test...")
	a_car := car {
		accelerator: 65000,
		brake: 0,
		steering_wheel: 12561,
		top_speed_kmh: 225.0 }
	// shorthand but less descriptive
	// a_car := car{22341,0,12561,120} 

	fmt.Println(a_car.steering_wheel)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())
	a_car.new_top_speed(500)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())
}


