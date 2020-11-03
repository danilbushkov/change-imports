package file

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"strings"
)

// ChangeTextFile returns the buffer with the changed import
func ChangeTextFile(path string, old string, new string) *bytes.Buffer {
	var isImport bool
	var text string
	var changed bool
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
				changed = true
			}
		}

		if isImport && !strings.Contains(text, ")") {
			if strings.Contains(text, old) {
				text = strings.ReplaceAll(text, old, new)
				changed = true
			}
		} else {
			isImport = false
		}

		buffer.WriteString(text)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	if changed {
		return &buffer
	}
	return nil
}
