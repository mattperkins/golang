package main

import ("fmt"
		"net/http")

func index_handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, `<h1>Heading level one</h1>
		<h2>Multi line formatting</h2>
		`)
	fmt.Fprintf(w, "<h3> %s </h3>", "Variable"	)
}

func main() {
	http.HandleFunc("/", index_handler)
	fmt.Println("Running at: http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}

 