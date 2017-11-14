package main

import ("fmt"
		"net/http"
		"html/template"
		"io/ioutil"
		"encoding/xml"
		"sync"	
	)

var wg sync.WaitGroup

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword string
	Location string
}	

type NewsAggPage struct{
	Title string
	News map[string]NewsMap
}

func indexHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<a href='/agg'><h1>Get The News</h1></a>")
}

func newsRoutine(c chan News, Location string) {
	defer wg.Done()
	var n News
	resp, _ := http.Get(Location)
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &n)
	resp.Body.Close()
	c <- n
}

func newsAggHandler(w http.ResponseWriter, r *http.Request){
	var s SitemapIndex
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)
	news_map := make(map[string]NewsMap) 
	resp.Body.Close()
	queue := make(chan News, 30)
	for _, Location := range s.Locations{
		wg.Add(1)
		go newsRoutine(queue, Location) 
	}

	wg.Wait()
	close(queue)

	for elem := range queue {
			for idx, _ := range elem.Titles {
			news_map[elem.Titles[idx]] = NewsMap{elem.Keywords[idx], elem.Locations[idx]}
		}
	}

	for idx, data := range news_map{
		 fmt.Println("\n\n\n",idx)
		 fmt.Println("\n",data.Keyword)
		 fmt.Println("\n",data.Location)
	 }

	p := NewsAggPage{Title: "WashPo Aggregator", News: news_map}
	t, _ := template.ParseFiles("template-r1.html")
	// fmt.Println(t.Execute(w, p)) // check for errors in template.html
	t.Execute(w, p)
}


func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg", newsAggHandler)
	fmt.Println("Running at: http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}

 