package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// TODO: fix linter errors
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

func xmlLint(schema string, xml string) {
	cmd := exec.Command("xmllint", "--noout", "--schema", schema, xml)

	stderr, err := cmd.StderrPipe()
	if err != nil {
		panic(err)
	}

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	slurp, _ := ioutil.ReadAll(stderr)

	if strings.TrimSpace(string(slurp)) != xml+" validates" {
		message := string(slurp)
		message = strings.Replace(message, xml+" fails to validate\n", "", 1)

		fmt.Printf("%s\n", message)
	}

	cmd.Wait()
}

func main() {
	c := Config{}
	c.Setup()
	flag.Parse()

	schema := c.schema
	if schema == "" {
		fmt.Println("--schema is missing")
		return
	}
	folder := c.folder
	if folder == "" {
		fmt.Println("--folder is missing")
		return
	}

	_, err := os.Stat(schema)
	if err != nil {
		fmt.Println("Schema file was not found")
		return
	}

	_, err = os.Stat(folder)
	if err != nil {
		fmt.Println("XML folder was not found")
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
