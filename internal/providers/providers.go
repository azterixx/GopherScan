package providers

import (
	"fmt"
	"net/url"
	"strings"
)

type Provider struct {
	Name string
	URL  string
}

var builtinProviders = []Provider{
	{"crt.sh", "crt.sh/?q=%s&output=json"},
	{"UrlScan", "urlscan.io/api/v1/search/?q=domain:%s"},
	{"HackerTarget", "api.hackertarget.com/hostsearch/?q=%s"},
}

var outdatedProviders = []Provider{
	{"AnubisDB", "anubisdb.com/anubis/subdomains/%s"},
}

func Build(domain string, outdated bool) []Provider {

	domain, _ = hostFromURL(domain)
	out := make([]Provider, len(builtinProviders))
	for i, p := range builtinProviders {
		out[i] = Provider{p.Name, "https://" + fmt.Sprintf(p.URL, domain)}

	}
	if outdated {
		out = append(out, outdatedProviders...)

	}
	return out

}

func hostFromURL(raw string) (string, error) {
	raw = strings.TrimSpace(raw)
	raw = strings.TrimSuffix(raw, "/")

	if !strings.Contains(raw, "://") {
		raw = "http://" + raw
	}

	u, err := url.Parse(raw)
	if err != nil {
		return "", fmt.Errorf("parse %q: %w", raw, err)
	}

	host := u.Hostname()
	if host == "" {
		return "", fmt.Errorf("no hostname in %q", raw)
	}
	return host, nil
}
