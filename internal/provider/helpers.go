package provider

import (
	"fmt"
	"net/url"
	"strings"
)

func unique(in []string) []string {
	seen := make(map[string]struct{}, len(in))
	for _, v := range in {
		if v != "" {
			seen[v] = struct{}{}
		}

	}
	out := make([]string, 0, len(seen))
	for k := range seen {
		out = append(out, k)
	}
	return out

}

func hostFromURL(raw string) (string, error) {

	raw = strings.TrimSpace(raw)
	raw = strings.TrimSuffix(raw, "/")

	if !strings.Contains(raw, "://") {
		raw = "https://" + raw
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
