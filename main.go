package main

import (
	"log"
	"os"

	"github.com/allenslian/vanda/cmd"
)

func main() {
	if err := cmd.Execute(os.Args); err != nil {
		log.Fatal(err)
	}
}
