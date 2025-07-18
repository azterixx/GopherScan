package subdomain

import (
	"GopherScan/internal/adapters/fetch"
	"GopherScan/internal/adapters/pinger"
	"GopherScan/internal/adapters/provider"
)

func Scan(domain string, history bool) []string {
	urls := provider.Build(domain, history)
	results := fetch.All(urls)
	if !history {
		results = pinger.Active(results)
	}
	return results

}
