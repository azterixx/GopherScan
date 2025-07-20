package provider

import (
	"fmt"
	"github.com/azterixx/GopherScan/internal/platform/utils"
)

type Provider struct {
	Name    string
	URL     string
	Paths   []string
	SplitNL bool
	Parser  Parse
}

var builtinProviders = []Provider{
	{
		Name:    "crt.sh",
		URL:     "crt.sh/?q=%s&output=json",
		Paths:   []string{"#.name_value"},
		SplitNL: true,
	},
	{
		Name:   "HackerTarget",
		URL:    "api.hackertarget.com/hostsearch/?q=%s",
		Parser: csvParse,
	},
}

var outdatedProviders = []Provider{
	{Name: "AnubisDB", URL: "anubisdb.com/anubis/subdomains/%s", Paths: []string{""}},
}

func Build(domain string, outdated bool) []Provider {
	domain, _ = utils.HostFromURL(domain)
	out := make([]Provider, len(builtinProviders))
	for i, p := range builtinProviders {
		out[i] = Provider{
			p.Name,
			"https://" + fmt.Sprintf(p.URL, domain),
			p.Paths,
			p.SplitNL,
			p.Parser,
		}

	}
	if outdated {
		out = append(out, outdatedProviders...)

	}
	return out

}
