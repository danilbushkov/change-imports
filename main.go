package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"path/filepath"
)

func main() {
	//var files []string
	if len(os.Args) < 2 {
		log.Fatal("Error: No values ​​set")
	}

	//fmt.Println(os.Args[1])
	//getFiles(".", &files)
	//fmt.Println(files)
	buffer := getTextFile("test.go")
	err := os.Remove("test.go")
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("test.go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.WriteString(buffer.String())
}

func getTextFile(path string) *bytes.Buffer {
	buffer := bytes.Buffer{}
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		buffer.WriteString(scanner.Text() + "\n")
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
