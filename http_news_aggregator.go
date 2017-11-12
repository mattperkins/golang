// [5 5]int = array
// []int = slice

package main

import ("fmt"
"net/http"
"io/ioutil"
"encoding/xml")



func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
}


 
