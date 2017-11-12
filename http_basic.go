package main

import ("fmt"
		"net/http" )

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!!")
}
func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the 'About' page")
}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about", about_handler)
	fmt.Println("Server running at http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}