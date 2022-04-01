package main

import (
	"flag"
	"fmt"
	"home-task5/pkg/crawler"
	"home-task5/pkg/crawler/spider"
	"home-task5/pkg/index"
	"sort"
	"strings"
)

type parsing struct {
	sites []string
	scanner *spider.Service
	ind *index.Index
}

func new() *parsing {
	return &parsing{
		sites: []string{"https://go.dev"},
		scanner: spider.New(),
		ind: index.New(),
	}
}

func (p *parsing) scanSites() []crawler.Document {
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

func (p *parsing) indexing(docs []crawler.Document)  {
	p.ind.Add(docs)
}

func main() {
	var query string
	flag.StringVar(&query, "s", "golang", "You know... for search.")
	flag.Parse()

	server := new()
	docs := server.scanSites()
	server.indexing(docs)

	if flagPresent("s") {
		ids := server.ind.Search(query)

		if len(ids) > 0 {
			urls := targetUrls(docs, ids)

			fmt.Println("Результаты поиска: ")
			fmt.Println(strings.Join(urls, "\n"))
		} else {
			fmt.Printf("Строка: %v не найдена.", query)
		}
	}
}
func targetUrls(docs []crawler.Document, ids []int) []string {
	var result []string

	for _, id := range ids {
		i := sort.Search(len(docs), func(i int) bool { return docs[i].ID >= id })

		if docs[i].ID == id {
			result = append(result, docs[i].URL)
		}
	}
	return result
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