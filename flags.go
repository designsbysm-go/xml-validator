package main

import (
	"errors"
	"flag"
	"os"
)

// Config holds the cli flag values
type Config struct {
	folder   string
	progress string
	schema   string
}

// Setup processes the passed cli flag
func (c *Config) Setup() {
	flag.StringVar(&c.folder, "folder", "", "xml files location")
	flag.StringVar(&c.progress, "progress", "show", "show/hide progress bar")
	flag.StringVar(&c.schema, "schema", "", "schema files location")
}

// Verify the cli flags exsists and the values are valid
func (c *Config) Verify() (string, string, error) {
	if c.schema == "" {
		return "", "", errors.New("--schema is missing")
	}

	if c.folder == "" {
		return "", "", errors.New("--folder is missing")
	}

	_, err := os.Stat(c.schema)
	if err != nil {
		return "", "", errors.New("Schema file was not found")
	}

	_, err = os.Stat(c.folder)
	if err != nil {
		return "", "", errors.New("XML folder was not found")
	}

	return c.schema, c.folder, nil
}
