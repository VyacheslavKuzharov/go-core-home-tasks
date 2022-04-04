package storage

import (
	"bufio"
	"encoding/json"
	"errors"
	"home-task5/pkg/crawler"
	"io"
	"os"
)

const FileName = "sites-data.json"

type Storage struct {}

func New() *Storage {
	return &Storage{}
}

func (st *Storage) NewFile() (*os.File, error) {
	newFile, err := os.Create(FileName)
	if err != nil {
		return nil, err
	}

	return newFile, nil
}

func (st *Storage) StoreDocs(file *os.File, docs *[]crawler.Document ) error {
	data, err := json.MarshalIndent(*docs, "", " ")
	if err != nil {
		return err
	}

	err = saveToStorage(file, data)
	if err != nil {
		return err
	}

	err = closeStorage(file)
	if err != nil {
		return err
	}

	return nil
}

func (st *Storage) ReadDocs() ([]crawler.Document, error)  {
	var docs []crawler.Document

	reader, err := openStorage()
	if err != nil {
		return nil, err
	}

	b, err := readStorage(reader)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &docs)
	if err != nil {
		return nil, err
	}

	err = closeStorage(reader)
	if err != nil {
		return nil, err
	}

	return docs, nil
}

func openStorage() (*os.File, error) {
	file, err := os.Open(FileName)
	if err != nil {
		return nil, err
	}

	return file, nil
}

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

func saveToStorage(w io.Writer, data []byte) error {
	_, err := w.Write(data)
	return err
}

func closeStorage(c io.Closer) error {
	err := c.Close()
	return err
}

func FileExists() bool {
	_, err := os.Stat(FileName)
	return !errors.Is(err, os.ErrNotExist)
}