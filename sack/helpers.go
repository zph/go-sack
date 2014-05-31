package sack

import (
	"io/ioutil"
	"path"
	"strings"
)

func checkState() {}

func splitLine(s string) []string {
	arr := strings.Split(s, ":")
	return arr
}

func check(e error) {
	if e != nil {
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
