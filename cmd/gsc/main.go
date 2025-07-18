package main

import (
	"GopherScan/internal/adapters/cli"
	"GopherScan/internal/app/subdomain"
	"fmt"
	"os"
)

func main() {
	cfg, err := cli.ParseFlags(os.Args[1:])
	if err != nil {
		os.Exit(1)
	}
	results := subdomain.Scan(cfg.Domain, cfg.History)
	for _, r := range results {
		fmt.Println(r)
	}
}
