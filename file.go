package main

import (
	"io/ioutil"
	"os"
	"strings"
)

func ReadDir(dir string) (fns []string) {
	files, _ := ioutil.ReadDir(dir)
	for _, f := range files {
		fns = append(fns, f.Name())
	}
	return
}

func Rename(name, from, to string) (e error) {
	replaced := strings.Replace(name, from, to, -1)
	e = os.Rename(name, replaced)
	if e != nil {
		return e
	}
	return nil
}
