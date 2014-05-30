package main

import (
	"flag"
	"fmt"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
	"syscall"
)

/*
TODO:
- add commandline args
- convert ag's output to sack_shortcuts format
- add -e arg for exec'ing vim w/ appropriate sack_shortcuts format args
*/

var home string = os.Getenv("HOME")
var searchPath string = path.Join(home, ".zsh.d/")

// var flagEdit = flag.Bool("edit", false, "zoom to edit this file")

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
	agBin, err := exec.LookPath(agCmd)

	cmd, err := exec.Command(agBin, flags, searchTerm, searchPath).Output()
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

	ind, err := strconv.Atoi(s)
	check(err)

	selectedLine := lines[ind]
	lineArr := strings.Split(selectedLine, " ")

	env := os.Environ()
	vimBin, err := exec.LookPath("vim")
	check(err)

	plusCmd := fmt.Sprint("+", lineArr[0])
	plussCmd := []string{"vim", lineArr[1], plusCmd}
	fmt.Println("Index entry: ", s)
	fmt.Println("Whole cmd: ", plussCmd)

	if true {
		execErr := syscall.Exec(vimBin, plussCmd, env)
		check(execErr)
	}
}

func display() {}

func setup() []string {
	flag.Parse()
	var args []string = flag.Args()
	return args
}

func checkState() {}

func main() {
	checkState()

	app := cli.NewApp()
	app.Name = "Sack"
	app.Usage = "sack [searchterm] [optional directory]"
	app.Version = "0.1.0"
	app.Flags = []cli.Flag{
		cli.BoolFlag{"edit, e", "edit a given shortcut"},
		cli.BoolFlag{"search, s", "search-ack/ag it"},
	}

	app.Action = func(c *cli.Context) {
		if c.Bool("edit") {
			edit(c.Args()[0])
		} else if c.Bool("search") || true {
			search()
		}
	}
	app.Run(os.Args)
}
