package main

import (
	"fmt"
	"net/url"
)

func main() {
	urlExamples := []string{
		"http://www.example1.com",
		"http://www.example2.com/",
		"www.example3.com",
		"www.example4.com/",
		"example5.com/",
		"example6.com",
		"https://sc2-10-185-16-54.eng.vmware.com",
	}

	for i, v := range urlExamples {
		fmt.Println(i)
		u, err := url.Parse(v)
		if err != nil {
			fmt.Println(err)
		}
		if u.Scheme == "" {
			u.Scheme = "https"
		}
		u.Path = "/test"
		fmt.Println("first", u)

		u.Path = "test/"
		fmt.Println("second", u)
	}

}
