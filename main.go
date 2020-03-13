package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// TODO: add unit tests

/* func readFileToBuffer(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buffer, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return buffer, nil
} */

func main() {
	c := Config{}
	c.Setup()
	flag.Parse()

	result := c.Verify()
	if result != "" {
		fmt.Println(result)
		return
	}

	schema := c.schema
	folder := c.folder

	var files []string

	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}

	var found int

	for _, file := range files {
		if strings.HasSuffix(file, ".xml") {
			found++
			xmlLint(schema, file)
		}
	}

	if found == 0 {
		fmt.Println("No XML files found")
	} else {
		fmt.Printf("Files validated: %v\n", found)
	}
}
