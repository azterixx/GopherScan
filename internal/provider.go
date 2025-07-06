package internal

import (
	"fmt"
	"net/url"
)

type provider struct {
	Name string
	URL  string
}

var providers = []provider{
	{"crt.sh", "https://crt.sh/?q=%s&output=json"},
	{"ThreatMiner", "https://api.threatminer.org/v2/domain.php?q=%s&rt=5"},
	{"AnubisDB", "https://anubisdb.com/anubis/subdomains/%s"},
	{"WaybackCDX", "https://web.archive.org/cdx/search/cdx?url=%s&output=json"},
	{"ThreatCrowd", "https://www.threatcrowd.org/searchApi/v2/domain/report/?domain=%s"},
	{"CertSpotter", "https://api.certspotter.com/v1/issuances?domain=%s"},
}

func BuildProviders(domain string) []provider {
	out := make([]provider, len(providers))
	domain = url.QueryEscape(domain)
	for i, p := range providers {
		out[i] = provider{p.Name, fmt.Sprintf(p.URL, domain)}

	}
	return out

}
