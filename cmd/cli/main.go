package main

import (
	"GopherScan/internal/fetch"
	"GopherScan/internal/provider"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: gosc <domain>")
		os.Exit(1)
	}
	domain := os.Args[1]
	urls := provider.Build(domain, false)
	results := fetch.All(urls)

	for _, result := range results {
		fmt.Println(result)
	}

}
