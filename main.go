package main

import (
	"log"
	"os"

	_ "goinaction/rssreader/matchers"
	"goinaction/rssreader/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
