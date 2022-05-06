package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func userInput() (string, error) {
	var reader = bufio.NewReader(os.Stdin)

	fmt.Print("Введите запрос: ")

	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	input = strings.Replace(input, "\n", "", -1)

	return input, nil

}

func main() {
	for {
		query, _ := userInput()
		if query == "exit" {
			break
		}
		log.Println("query -->", query)
		conn, err := net.Dial("tcp4", "localhost:13")
		if err != nil {
			log.Fatal(err)
		}
		_, err = conn.Write([]byte(query))

		msg, err := io.ReadAll(conn)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Ответ от сервера:", string(msg))
	}
}
