package main

import (
	"errors"
	"flag"
	"os"
)

// Config holds the cli flag values
type Config struct {
	schema string
	folder string
}

// Setup processes the passed cli flag
func (c *Config) Setup() {
	flag.StringVar(&c.schema, "schema", "", "schema files location")
	flag.StringVar(&c.folder, "folder", "", "xml files location")
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
