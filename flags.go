package main

import (
	"flag"
)

type Config struct {
	schema string
	folder string
}

func (c *Config) Setup() {
	flag.StringVar(&c.schema, "schema", "", "schema files location")
	flag.StringVar(&c.folder, "folder", "", "xml files location")
}
