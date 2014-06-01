package sack

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"
)

func checkState() {}

func splitLine(s string) []string {
	arr := strings.SplitN(s, ":", 3)
	return arr
}

func check(e error) {
	if e != nil {
		fmt.Printf("\n----\nError: %#v\n----\n", e)
		panic(e)
	}
}

func content() []string {
	filePath := path.Join(home, shortcutFilename)
	dat, err := ioutil.ReadFile(filePath)
	check(err)
	lines := strings.Split(string(dat), "\n")
	return lines[0 : len(lines)-1]
}
