package storage

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"home-task5/pkg/crawler"
	"io"
	"log"
	"os"
)

const FileName = "sites-data.json"

type Storage struct {
	file *os.File
}

func New() *Storage {
	return &Storage{}
}

func (st *Storage) NewFile() *os.File {
	newFile, err := os.Create(FileName)
	if err != nil {
		fmt.Printf("Сбой создания файла: %s", err)
	}

	st.file = newFile
	return st.file
}

func (st *Storage) StoreDocs(docs *[]crawler.Document )  {
	data, err := json.MarshalIndent(*docs, "", " ")
	if err != nil {
		fmt.Printf("Сбой маршалинга JSON: %s", err)
	}

	saveToStorage(st.file, data)
	closeStorage(st.file)
}

func (st *Storage) ReadDocs() []crawler.Document  {
	var docs []crawler.Document

	reader := openStorage()
	b, err := readStorage(reader)
	if err != nil {
		fmt.Printf("Сбой чтения из хранилища: %s", err)
	}

	err = json.Unmarshal(b, &docs)
	if err != nil {
		fmt.Printf("Сбой демаршалинга JSON: %s", err)
	}

	closeStorage(reader)

	return docs
}

// сейчас хранилище это json файл
func openStorage() *os.File {
	file, err := os.Open(FileName)
	if err != nil {
		log.Fatal(err)
	}

	return file
}

// сейчас хранилище это json файл
func readStorage(r io.Reader) ([]byte, error) {
	var b []byte

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		b = append(b, []byte(scanner.Text()+"\n")...)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return b, nil
}

func saveToStorage(w io.Writer, data []byte)  {
	_, err := w.Write(data)
	if err != nil {
		fmt.Printf("Сбой записи в хранилище: %s", err)
	}
}

func closeStorage(c io.Closer) {
	err := c.Close()
	if err != nil {
		fmt.Printf("Сбой при закрытии хранилища: %s", err)
	}
}

func FileExists() bool {
	_, err := os.Stat(FileName)
	return !errors.Is(err, os.ErrNotExist)
}