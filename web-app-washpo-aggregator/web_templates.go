package main

import ("fmt"
		"net/http"
		"html/template")

type NewsAggPage struct{
	Title string
	News string
}

func newsAggHandler(w http.ResponseWriter, r *http.Request){
	p := NewsAggPage{Title: "The News Aggregator", News: "Latest News"}
	t, _ := template.ParseFiles("template.html")
	// fmt.Println(t.Execute(w, p)) // check for errors in template.html
	t.Execute(w, p)
}

func indexHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<a href='/agg'><h1>Get The News</h1></a>")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg", newsAggHandler)
	fmt.Println("Running at: http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}

 