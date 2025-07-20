package fetch

import (
	"context"
	p "github.com/azterixx/GopherScan/internal/adapters/provider"
	"io"
	"net/http"
	"sync"
	"time"
)

func All(list []p.Provider) []string {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	client := &http.Client{Timeout: 15 * time.Second}

	var (
		wg  sync.WaitGroup
		mu  sync.Mutex
		set = make(map[string]struct{})
	)

	for _, prov := range list {
		wg.Add(1)
		go func(pr p.Provider) {
			defer wg.Done()

			req, err := http.NewRequestWithContext(ctx, http.MethodGet, pr.URL, nil)
			if err != nil {
				return
			}

			resp, err := client.Do(req)

			if err != nil {
				return
			}
			body, _ := io.ReadAll(resp.Body)
			resp.Body.Close()

			urls, err := pr.Parse(body)
			if err != nil {
				return
			}

			mu.Lock()
			for _, u := range urls {
				if u != "" {
					set[u] = struct{}{}
				}
			}
			mu.Unlock()
		}(prov)
	}
	wg.Wait()

	out := make([]string, 0, len(set))
	for u := range set {
		out = append(out, u)
	}
	return out
}
