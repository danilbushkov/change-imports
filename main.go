package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/danilbushkov/change-imports/file"
)

func main() {
	var filePaths []string
	var path string
	var old string
	var new string
	var buffer *bytes.Buffer
	var f *os.File
	var err error
	if len(os.Args) < 3 {
		fmt.Println("Necessary parameters:")
		fmt.Println("    app <path> <old name import> <new name import>")
		os.Exit(1)
	}
	path = os.Args[1]
	old = os.Args[2]
	new = os.Args[3]

	file.GetFiles(path, &filePaths)
	for _, p := range filePaths {
		buffer = file.ChangeTextFile(p, old, new)
		if buffer != nil {
			f, err = os.Create(p)
			if err != nil {
				log.Fatal(err)
			}
			f.WriteString(buffer.String())
			f.Close()
		}
	}
}
