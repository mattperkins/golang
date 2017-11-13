// [5 5]int = array
// []int = slice

package main

import ("fmt"
"net/http"
"io/ioutil"
"encoding/xml")

// var washPostXML = []byte(`<sitemapindex>
// 	<sitemap></sitemap>
// 	<loc>http://exampleofusinglocalxml.com</loc>
// 	</sitemapindex>`)

type SitemapIndex struct {
	Locations []Location `xml:"sitemap"`
}

type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	// bytes := washPostXML

	var s SitemapIndex
	xml.Unmarshal(bytes, &s)

	fmt.Println(s.Locations)
}
