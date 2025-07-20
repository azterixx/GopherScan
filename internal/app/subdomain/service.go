package subdomain

import (
	"github.com/azterixx/GopherScan/internal/adapters/fetch"
	"github.com/azterixx/GopherScan/internal/adapters/pinger"
	"github.com/azterixx/GopherScan/internal/adapters/provider"
)

func Scan(domain string, history bool) []string {
	urls := provider.Build(domain, history)
	results := fetch.All(urls)
	if !history {
		results = pinger.Active(results)
	}
	return results

}
