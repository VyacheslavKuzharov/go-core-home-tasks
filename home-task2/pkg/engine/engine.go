package engine

import (
	"flag"
	"home-task2/pkg/crawler"
	"strings"
)

func Search(docs []crawler.Document) []string {
	var query string
	var results []string
	flag.StringVar(&query, "s", "golang", "You know... for search.")
	flag.Parse()

	if flagPresent("s") {
		for _, doc := range docs {
			res := strings.Contains(strings.ToLower(doc.Title), strings.ToLower(query))

			if res {
				results = append(results, doc.URL)
			}
		}
	}

	return results
}

func flagPresent(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
