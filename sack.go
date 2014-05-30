package main

import (
	// "bufio"
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

const agCmd string = "ag"
const flags string = "-i"

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

func executeCmd(searchTerm string, searchPath string) []string {
	agBin, err := exec.LookPath(agCmd)

	fmt.Println("PPath: ", searchPath)
	cmd, err := exec.Command(agBin, flags, searchTerm, searchPath).Output()
	check(err)
	lines := strings.Split(string(cmd), "\n")
	fmt.Println(lines)
	return lines
}

func search(c *cli.Context) {
	searchTerm := c.Args()[0]
	searchPath := c.Args()[1]
	// TODO: allow PWD as default value for searchPath
	//     searchPath, _ = os.Getwd()

	fmt.Println("Term ", searchTerm, " Path: ", searchPath)
	lines := executeCmd(searchTerm, searchPath)
	fmt.Println(strings.Join(lines, "\n"))
	// TODO: transform to sack_shortcuts format
	// TODO: write to ~/.sack_shortcuts
}

func edit(c *cli.Context) {
	lines := content()

	ind, err := strconv.Atoi(c.Args()[0])
	check(err)

	selectedLine := lines[ind]
	lineArr := strings.Split(selectedLine, " ")

	env := os.Environ()
	vimBin, err := exec.LookPath("vim")
	check(err)

	plusCmd := fmt.Sprint("+", lineArr[0])
	plussCmd := []string{"vim", lineArr[1], plusCmd}

	if c.Bool("debug") {
		fmt.Println("Whole cmd: ", plussCmd, " Index: ", c.Args()[0])
	}

	if true {
		execErr := syscall.Exec(vimBin, plussCmd, env)
		check(execErr)
	}
}

func display() {
	lines := content()
	fmt.Println(len(lines))
	length := len(lines) - 1

	for i := 0; i < length; i++ {
		fmt.Println("[", i, "]", lines[i])
	}
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
		cli.BoolFlag{"print, p", "display existing shortcuts"},
		cli.BoolFlag{"debug, d", "show all the texts"},
	}

	app.Action = func(c *cli.Context) {

		if c.Bool("debug") {
			fmt.Printf("Context %#v\n", c)
		}

		if c.Bool("edit") {
			edit(c)
		} else if c.Bool("search") {
			search(c)
		} else if c.Bool("print") || true {
			display()
		}
	}
	app.Run(os.Args)
}
