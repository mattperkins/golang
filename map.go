// in terminal run: gofmt go run file.go (auto format file)

package main

import "fmt"

func main(){
	// key = string value = float32
	// map is a reference type (doesnt have any values - so we use make to initialise)
	Grades := make(map[string]float32)

	Grades["Freddy"] = 42
	Grades["Sandy"] = 93
	Grades["Bob"] = 58

	fmt.Println(Grades)

	FreddyGrade := Grades["Freddy"]
	fmt.Println("Freddy scored:", FreddyGrade)

	delete(Grades,"Bob")
	fmt.Println("Deleted : Bob")
	fmt.Println(Grades)

	for key, value := range Grades {
		fmt.Println(key, ":", value)
	}
}