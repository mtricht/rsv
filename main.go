package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	rsv, _ := os.Open("example.rsv")
	defer rsv.Close()
	scanner := bufio.NewScanner(rsv)
	scanner.Split(bufio.ScanLines)

	var headers []string
	headersRead := false
	var row []string
	for scanner.Scan() {
		value := ""
		if false == headersRead {
			text := scanner.Text()
			parts := strings.Split(text, ",")
			headers = append(headers, parts[0])
			if len(parts) == 1 {
				continue
			}
			headersRead = true
			value = parts[1]
		} else {
			value = scanner.Text()
		}
		parts := strings.Split(value, ",")
		row = append(row, parts[0])
		if len(parts) == 1 {
			continue
		}
		print(headers, row)
		row = []string{parts[1]}
	}
	print(headers, row)
}

func print(headers []string, row []string) {
	log.Println("--------------")
	for index, header := range headers {
		log.Printf("%s: %s", header, row[index])
	}
}
