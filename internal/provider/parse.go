package provider

import (
	"github.com/tidwall/gjson"
	"strings"
)

type Parse func([]byte) ([]string, error)

func (p Provider) Parse(body []byte) ([]string, error) {
	if p.Parser != nil {
		return p.Parser(body)
	}
	var out []string
	for _, path := range p.Paths {
		gjson.GetBytes(body, path).ForEach(
			func(_, v gjson.Result) bool {
				if p.SplitNL {
					out = append(out, strings.Split(v.String(), "\n")...)
				} else {
					out = append(out, v.String())
				}
				return true
			},
		)
	}
	for i, r := range out {
		out[i] = strings.TrimPrefix(r, "*.")
	}
	return unique(out), nil
}

func csvParse(b []byte) ([]string, error) {
	lines := strings.Split(strings.TrimSpace(string(b)), "\n")
	out := make([]string, 0, len(lines))
	for _, l := range lines {
		if idx := strings.IndexByte(l, ','); idx > 0 {
			out = append(out, l[:idx])
		}
	}
	return unique(out), nil
}
