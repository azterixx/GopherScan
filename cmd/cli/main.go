package main

import (
	"GopherScan/internal/fetch"
	"GopherScan/internal/pinger"
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
	alive := pinger.Active(results)

	for _, result := range alive {
		fmt.Println(result)
	}

}
