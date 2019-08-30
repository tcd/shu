package main

import (
	"log"

	"github.com/tcd/shu/cmd"
)

func init() {
	log.SetFlags(0)
}

func main() {
	cmd.Execute()
}
