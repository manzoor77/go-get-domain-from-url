package main

import (
	"fmt"
	"net/url"
	"strings"
)

func getdomain(urlSlice ...string) {
	//for loop to iterate the whole slice
	for _, urls := range urlSlice {
		fmt.Println(urls)
		u, err := url.Parse(urls)
		if err != nil {
			panic(err)
		}
		//Split string into parts and store into map
		parts := strings.Split(u.Hostname(), ".")
		fmt.Println(parts)
		fmt.Println(len(parts))
		fmt.Println(parts[0])
		fmt.Println(parts[1])

		domain := parts[len(parts)-2] + "." + parts[len(parts)-1]
		fmt.Println(domain)

	}
}

func main() {
	// URLS Slice
	urlSlice := []string{"https://mail.google.com/mail/u/0/#inbox", "https://github.com/manzoor77", "https://linkedin.com/in/manzoor-/"}
	//PASS URL Slice to getdomain function by using Variadic function
	getdomain(urlSlice...)
}
