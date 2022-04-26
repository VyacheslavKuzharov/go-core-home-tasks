package main

import (
	"fmt"
	"home-task11/pkg/crawler"
	"home-task11/pkg/crawler/spider"
	"home-task11/pkg/index"
	"home-task11/pkg/utils"
	"sort"
	"strings"
)

type parser struct {
	sites   []string
	scanner *spider.Service
	index   *index.Index
}

func main() {
	var query = "language"

	server := new()
	docs := server.fetchDocs()
	server.index.Add(&docs)

	ids := server.index.Search(query)

	if len(ids) > 0 {
		urls := utils.TargetUrls(&docs, ids)

		fmt.Println("Результаты поиска: ")
		fmt.Println(strings.Join(urls, "\n"))
	} else {
		fmt.Printf("Строка: %v не найдена.", query)
	}
}

func new() *parser {
	return &parser{
		sites:   []string{"https://go.dev"},
		scanner: spider.New(),
		index:   index.New(),
	}
}

func (p *parser) fetchDocs() []crawler.Document {
	var docs []crawler.Document

	docs = p.scanSites()

	return docs
}

func (p *parser) scanSites() []crawler.Document {
	var allDocs []crawler.Document

	for _, site := range p.sites {
		fmt.Println("Сканирование сайта: ", site)
		docs, err := p.scanner.Scan(site, 2)
		if err != nil {
			fmt.Println("ошибка при добавлении сканировании документов:", err)
		}
		allDocs = append(allDocs, docs...)
	}

	for id := range allDocs {
		allDocs[id].ID = id
		id++
	}
	sort.Slice(allDocs, func(i, j int) bool { return allDocs[i].ID < allDocs[j].ID })

	fmt.Println("Сканирование завершено")
	return allDocs
}
