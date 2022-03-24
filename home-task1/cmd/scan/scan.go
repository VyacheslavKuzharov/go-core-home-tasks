package main

import (
	"fmt"
	"home-task1/pkg/crawler"
	"home-task1/pkg/crawler/spider"
	"home-task1/pkg/engine"
	"strings"
)

func main() {
	docs := scanSites("https://go.dev", "https://golang.org/")
	results := engine.Search(docs)

	if len(results) > 0 {
		fmt.Println("Результаты поиска: ")
		fmt.Println(strings.Join(results, "\n"))
	}
}

func scanSites(sites ...string) []crawler.Document {
	var allDocs []crawler.Document
	scanner := spider.New()

	for _, site := range sites {
		fmt.Println("Сканирование сайта: ", site)
		docs, err := scanner.Scan(site, 2)
		if err != nil {
			fmt.Println("ошибка при добавлении сканировании документов:", err)
		}

		allDocs = append(allDocs, docs...)
	}

	fmt.Println("Сканирование завершено")
	return allDocs
}
