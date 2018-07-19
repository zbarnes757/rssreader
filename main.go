package main

import (
	"log"
	"os"

	_ "rssreader/matchers"
	"rssreader/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
