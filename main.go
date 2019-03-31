package main

import (
	"flag"
	"path/filepath"
)

var from = flag.String("from", "", "from rule")
var to = flag.String("to", "", "to rule")
var path = flag.String("path", "", "file path")

func main() {
	flag.Parse()

	dir, file := filepath.Split(*path)
	files := []string{file}

	if dir == *path {
		files = ReadDir(dir)
	}

	for _, file := range files {
		e := Rename(filepath.Join(dir, file), *from, *to)
		if e != nil {
			panic(e)
			return
		}
	}
}
