package main

import "fmt"

type Attendee struct {
	Topic string
	Person Person
}
                            // define data structure
type Person struct {
	Name string
	City string
	Company string
}

type Listener struct {
	is_interested bool
	Person Person
}
							// bind method to structure via passing type of arg
							
// Attendee/Person as receiver(argument)
func (a Attendee) SayHi() {
	fmt.Printf("Hi, my name is %s and today I will talk about %s\n", a.Person.Name, a.Topic)
}

func (a Attendee) About() {
	fmt.Printf("I'm from %s, currently working at %s\n", a.Person.City, a.Person.Company)
}

func (p Person) SayHi() {
	fmt.Printf("Hi, my name is %s!\n", p.Name)
}

func (p Person) About() {
	fmt.Printf("I am from &s and currently working in %s\n", p.City, p.Company)
}

func (l Listener) Yawn() {
	if l.is_interested != true {
		fmt.Println("Yawwwwwnn!!!!")
	}
}

                 // define var of this data 'type' with this method attached
func main() {
	guy := Attendee {
		Topic: "'Why we should kill the morning stand up and sit down more'",
		Person: Person {
			Name: "John Doe",
			City: "Doeville",
			Company: "Doe and Co",
		},
	}
	harry := Listener {
			is_interested: false,
		Person: Person {
			Name: "Harry",
			City: "LA",
			Company: "Siliconia",
		},
		}
	guy.SayHi()
	guy.About()
	harry.Yawn()
}