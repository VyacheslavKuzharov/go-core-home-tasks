package main

import (
	"flag"
	"fmt"
	"home-task5/pkg/crawler"
	"home-task5/pkg/crawler/spider"
	"home-task5/pkg/index"
	"home-task5/pkg/storage"
	"home-task5/pkg/utils"
	"sort"
	"strings"
)

type parser struct {
	sites []string
	scanner *spider.Service
	index *index.Index
	storage *storage.Storage
}

func main() {
	var query string

	flag.StringVar(&query, "s", "golang", "You know... for search.")
	flag.Parse()

	server := new()
	docs := server.FetchDocs()
	server.index.Add(&docs)

	if utils.FlagPresent("s") {
		ids := server.index.Search(query)

		if len(ids) > 0 {
			urls := utils.TargetUrls(&docs, ids)

			fmt.Println("Результаты поиска: ")
			fmt.Println(strings.Join(urls, "\n"))
		} else {
			fmt.Printf("Строка: %v не найдена.", query)
		}
	}
}

func new() *parser {
	return &parser{
		sites: []string{"https://go.dev", "https://golang.org/"},
		scanner: spider.New(),
		index: index.New(),
		storage: storage.New(),
	}
}

func (p *parser) FetchDocs() []crawler.Document {
	var docs []crawler.Document

	if storage.FileExists() {
		return p.storage.ReadDocs()
	}

	docs = p.scanSites()
	p.storage.NewFile()
	p.storage.StoreDocs(&docs)

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