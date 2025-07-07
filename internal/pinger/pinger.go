package pinger

import (
	"io"
	"net"
	"net/http"
	"sort"
	"sync"
	"time"
)

const ua = "Go-Pinger/1.2"

var transport = &http.Transport{
	DialContext: (&net.Dialer{
		Timeout:   3 * time.Second,
		KeepAlive: 30 * time.Second,
	}).DialContext,
	TLSHandshakeTimeout:   3 * time.Second,
	ResponseHeaderTimeout: 4 * time.Second,
	MaxIdleConns:          200,
	MaxIdleConnsPerHost:   20,
	DisableCompression:    true,
}

var httpClient = &http.Client{
	Transport: transport,
	Timeout:   5 * time.Second,
	CheckRedirect: func(_ *http.Request, _ []*http.Request) error {
		return http.ErrUseLastResponse
	},
}

func Alive(host string) bool {
	return probe("https://"+host) || probe("http://"+host)
}

func Active(domains []string) []string {
	var wg sync.WaitGroup
	out := make(chan string, len(domains))
	for _, d := range domains {
		wg.Add(1)
		go func(h string) {
			defer wg.Done()
			if Alive(h) {
				out <- h
			}
		}(d)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	var alive []string
	for h := range out {
		alive = append(alive, h)
	}
	sort.Strings(alive)
	return alive
}

func probe(u string) bool {
	req, err := http.NewRequest(http.MethodHead, u, nil)
	if err != nil {
		return false
	}
	req.Header.Set("User-Agent", ua)
	resp, err := httpClient.Do(req)
	if err != nil {
		return false
	}
	_, _ = io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if resp.StatusCode == http.StatusMethodNotAllowed ||
		resp.StatusCode == http.StatusRequestedRangeNotSatisfiable {
		req, err = http.NewRequest(http.MethodGet, u, nil)
		if err != nil {
			return false
		}
		req.Header.Set("User-Agent", ua)
		req.Header.Set("Range", "bytes=0-0")
		resp, err = httpClient.Do(req)
		if err != nil {
			return false
		}
		_, _ = io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
	}
	return resp.StatusCode < 400

}
