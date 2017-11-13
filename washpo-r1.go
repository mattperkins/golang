package main

import ("fmt"
"net/http"
"io/ioutil"
"encoding/xml")

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	// bytes := washPostXML

	var s SitemapIndex
	xml.Unmarshal(bytes, &s)

	fmt.Println(s.Locations)

	for _, Location := range s.Locations{
		fmt.Printf("\n%s",Location)
	}
}
