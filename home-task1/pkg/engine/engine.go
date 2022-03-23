package engine

import (
	"flag"
	"home-task1/pkg/crawler"
	"strings"
)

var query string

func Search(docs []crawler.Document) []string {
	var results []string
	flag.StringVar(&query, "s", "golang", "You know... for search.")
	flag.Parse()

	if isFlagPassed("s") {
		for _, doc := range docs {
			res := strings.Contains(strings.ToLower(doc.Title), strings.ToLower(query))

			if res {
				results = append(results, doc.URL)
			}
		}
	}

	return results
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
