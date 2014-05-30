package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
)

/*
TODO:
- add commandline args
- convert ag's output to sack_shortcuts format
- add -e arg for exec'ing vim w/ appropriate sack_shortcuts format args
*/

var home string = os.Getenv("HOME")
var searchPath string = path.Join(home, ".zsh.d/")
var flagEdit = flag.Bool("edit", false, "zoom to edit this file")

const agCmd string = "ag"
const flags string = "-i"
const searchTerm string = "ruby"
const shortcutFilename string = ".sack_shortcuts"

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
	return lines
}

func splitLine(s string) []string {
	arr := strings.Split(s, ":")
	return arr
}

func executeCmd() []string {
	cmd, err := exec.Command(agCmd, flags, searchTerm, searchPath).Output()
	check(err)
	lines := strings.Split(string(cmd), "\n")
	return lines
}

func search() {
	lines := executeCmd()
	// firstLine := lines[0]
	// lineArr   := splitLine(firstLine)
	// fmt.Println(strings.Join(lineArr, "---"))
	fmt.Println(strings.Join(lines, "\n"))
}

func edit(s string) {
	lines := content()
	fmt.Println("Index entry: ", s)

	ind, err := strconv.Atoi(s)
	check(err)

	selectedLine := lines[ind]
	lineArr := strings.Split(selectedLine, " ")
	fmt.Println(strings.Join(lineArr, "---"))
}

func setup() []string {
	flag.Parse()
	var args []string = flag.Args()
	return args
}

func checkState() {}

func main() {
	checkState()
	args := setup()
	if *flagEdit {
		edit(args[0])
	} else {
		search()
	}
}
