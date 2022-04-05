package main

import (
	"flag"
	"fmt"
	"home-task5/pkg/crawler"
	"home-task5/pkg/crawler/spider"
	"home-task5/pkg/index"
	"home-task5/pkg/storage"
	"home-task5/pkg/utils"
	"log"
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
	docs := server.fetchDocs()
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

func (p *parser) fetchDocs() []crawler.Document {
	var docs []crawler.Document

	if storage.FileExists() {
		d, err := p.storage.ReadDocs()
		if err != nil {
			log.Fatalf("Сбой чтения файла: %s", err)
		}

		return d
	}

	docs = p.scanSites()
	file, err := p.storage.NewFile()
	if err != nil {
		log.Fatalf("Сбой создания файла: %s", err)
	}

	err = p.storage.StoreDocs(file, &docs)
	if err != nil {
		log.Fatalf("Сбой записи: %s", err)
	}

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