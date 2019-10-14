package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"regexp"
)

func main() {
	url := "https://www.ltd.org/system-map/route_91/"
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	var space = regexp.MustCompile(`(\s)(\s+)`)
	var clean string
	doc.Find(".alert").Each(func(i int, s *goquery.Selection) {
		clean += space.ReplaceAllString(s.Text(), "$1")
	})

	fmt.Println(clean)
}
