package index

import (
	"home-task11/pkg/crawler"
	"strings"
)

type Index struct {
	data map[string][]int
}

func New() *Index {
	var index Index
	index.data = make(map[string][]int)
	return &index
}

func (index *Index) Add(docs *[]crawler.Document) {
	for _, doc := range *docs {
		for _, key := range keys(doc.Title) {
			if !exists(index.data[key], doc.ID) {
				index.data[key] = append(index.data[key], doc.ID)
			}
		}
	}
}

func (index *Index) Search(key string) []int {
	return index.data[strings.ToLower(key)]
}

func keys(s string) []string {
	words := strings.Split(s, " ")
	for i := range words {
		words[i] = strings.ToLower(words[i])
	}
	return words
}

func exists(ids []int, item int) bool {
	for _, id := range ids {
		if id == item {
			return true
		}
	}
	return false
}
