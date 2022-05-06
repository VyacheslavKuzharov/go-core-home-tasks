package main

import (
	"bufio"
	"fmt"
	"home-task11/pkg/crawler"
	"home-task11/pkg/crawler/spider"
	"home-task11/pkg/index"
	"home-task11/pkg/utils"
	"log"
	"net"
	"os"
	"sort"
	"strings"
)

type parser struct {
	sites   []string
	scanner *spider.Service
	index   *index.Index
}

func handler(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	buf := make([]byte, 1024)
	log.Println("buf1 -->", buf)

	n, err := reader.Read(buf)

	query := string(buf[:n])

	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	log.Println("msg -->", query)

	server := new()
	docs := server.fetchDocs()
	server.index.Add(&docs)

	ids := server.index.Search(query)

	if len(ids) > 0 {
		urls := utils.TargetUrls(&docs, ids)

		//fmt.Println("Результаты поиска: ")
		//fmt.Println(strings.Join(urls, "\n"))
		str := strings.Join(urls, "\n")

		_, err = conn.Write([]byte(str))
		if err != nil {
			return
		}
	} else {
		fmt.Printf("Строка: %v не найдена.", query)
	}
}

func main() {
	listener, err := net.Listen("tcp4", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Запуск сервера на 0.0.0.0:8000 ")
	// цикл обработки клиентских подключений
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handler(conn)
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
