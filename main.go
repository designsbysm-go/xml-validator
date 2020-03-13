package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// TODO: add unit tests

func main() {
	c := Config{}
	c.Setup()
	flag.Parse()

	schema, folder, err := c.Verify()
	if err != nil {
		fmt.Println(err)
		return
	}

	var files []string

	err = filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
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
