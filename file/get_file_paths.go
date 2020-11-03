package file

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

//GetFilePaths returns a list of paths .go files
func GetFilePaths(p string, sl *[]string) {
	err := filepath.Walk(p,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.Contains(path, ".go") {
				*sl = append(*sl, path)
			}
			return nil
		})
	if err != nil {
		log.Fatal(err)
	}
}
