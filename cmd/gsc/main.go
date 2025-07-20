package main

import (
	"fmt"
	"github.com/azterixx/GopherScan/internal/adapters/cli"
	"github.com/azterixx/GopherScan/internal/app/subdomain"
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
