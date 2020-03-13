package main

import (
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
func (c *Config) Verify() string {
	if c.schema == "" {
		return "--schema is missing"
	}

	if c.folder == "" {
		return "--folder is missing"
	}

	_, err := os.Stat(c.schema)
	if err != nil {
		return "Schema file was not found"
	}

	_, err = os.Stat(c.folder)
	if err != nil {
		return "XML folder was not found"
	}

	return ""
}
