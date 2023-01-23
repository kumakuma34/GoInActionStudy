package main

import (
	_ "github.com/kumakuma34/GoInActionStudy/chapter2/sample/matchers"
	"github.com/kumakuma34/GoInActionStudy/chapter2/sample/search"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("Sherlock Holmes")
}
