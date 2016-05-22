package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"
)

func checkState() {}

func splitLine(s string) (string, string, string) {
	a := strings.SplitN(s, ":", 3)
	return a[0], a[1], a[2]
}

func check(e error) {
	if e != nil {
		fmt.Printf("\n----\nError: %#v\n----\n", e)
		panic(e)
	}
}

// TODO: hack, fix properly by learning to work w/ Golang Error Types
func checkCmd(e error) {
	str := fmt.Sprintf("%s", e)
	if e != nil && str != "exit status 1" {
		fmt.Printf("\n----\nError: %#v\n----\n", e)
		panic(e)
	}
}

func content() []string {
	filePath := path.Join(home, shortcutFilename)
	dat, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Unable to open shortcut file. Try doing a search.")
		panic(1)
	}
	lines := strings.Split(string(dat), "\n")
	return lines[0 : len(lines)-1]
}
