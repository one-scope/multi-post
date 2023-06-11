package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	CONFIG_FILE = "config.yaml"
)

func main() {
	tChannel := flag.String("channel", "default", "Set channel")
	flag.Parse()
	tContent := flag.Arg(0)

	tConfig, tError := ReadConfig(CONFIG_FILE)
	if tError != nil {
		fmt.Fprintln(os.Stderr, tError)
		os.Exit(1)
	}
}
