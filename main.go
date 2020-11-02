package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	//var files []string
	if len(os.Args) < 2 {
		log.Fatal("Error: No values ​​set")
	}

	//fmt.Println(os.Args[1])
	//getFiles(".", &files)
	//fmt.Println(files)
	buffer := changeTextFile("test/t", "test", "new")
	// err := os.Remove("test.go")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	file, err := os.Create("test/t")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.WriteString(buffer.String())
}

func changeTextFile(path string, old string, new string) *bytes.Buffer {
	var isImport bool
	var text string
	buffer := bytes.Buffer{}
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text = scanner.Text() + "\n"

		if !isImport && strings.Contains(text, "import") && strings.Contains(text, "(") {
			isImport = true
		} else if !isImport && strings.Contains(text, "import") {
			if strings.Contains(text, old) {
				text = strings.ReplaceAll(text, old, new)
			}
		}

		if isImport && !strings.Contains(text, ")") {
			if strings.Contains(text, old) {
				text = strings.ReplaceAll(text, old, new)
			}
		} else {
			isImport = false
		}

		buffer.WriteString(text)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	//fmt.Println(buffer.String())
	return &buffer
}

func getFiles(p string, sl *[]string) {
	err := filepath.Walk(p,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			*sl = append(*sl, path)
			return nil
		})
	if err != nil {
		log.Fatal(err)
	}
}
