package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

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
