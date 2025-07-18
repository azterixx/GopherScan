package cli

import (
	"flag"
	"fmt"
)

type Config struct {
	Domain  string
	History bool
}

func ParseFlags(args []string) (Config, error) {
	fs := flag.NewFlagSet("gsc", flag.ContinueOnError)

	cfg := Config{}
	fs.BoolVar(&cfg.History, "history", false, "include historical sub-domains")

	fs.Usage = func() {
		fmt.Fprintf(fs.Output(), "usage: %s [flags] <domain>\n\nflags:\n", fs.Name())
		fs.PrintDefaults()
	}

	if err := fs.Parse(args); err != nil {
		return cfg, err
	}

	if fs.NArg() < 1 {
		fs.Usage()
		return cfg, fmt.Errorf("domain is required")
	}
	cfg.Domain = fs.Arg(0)

	return cfg, nil
}
