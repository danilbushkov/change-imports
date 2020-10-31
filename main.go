package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	var files []string
	if len(os.Args) < 2 {
		log.Fatal("Error: No values ​​set")
	}

	fmt.Println(os.Args[1])
	getFiles(".", &files)
	fmt.Println(files)

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
