package utils

import (
	"home-task11/pkg/crawler"
	"sort"
)

func TargetUrls(docs *[]crawler.Document, ids []int) []string {
	var result []string
	d := *docs

	for _, id := range ids {
		i := sort.Search(len(d), func(i int) bool { return d[i].ID >= id })

		if d[i].ID == id {
			result = append(result, d[i].URL)
		}
	}
	return result
}
