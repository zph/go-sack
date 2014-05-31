package sack

import (
	"github.com/wsxiaoys/terminal/color"
	"bufio"
	"os/exec"
	"path"
	"strconv"
	"strings"
	"syscall"
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)

func display() {
	lines := content()

	// Header
	color.Printf("@r[%2s]@{|} @b%5s@{|}  @g%s@{|}\n", "IDX", "Line", "Path")

	for i, line := range lines {
		li := strings.Split(line, " ")
		s := color.Sprintf("@r[%2d]@{|} @b%5s@{|}  @g%s@{|}", i, li[0], li[1])
		fmt.Println(s)
	}
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

func executeCmd(searchTerm string, searchPath string) []string {
	agBin, err := exec.LookPath(agCmd)

	cmd, err := exec.Command(agBin, flags, searchTerm, searchPath).Output()
	check(err)
	lines := strings.Split(string(cmd), "\n")
	return lines
}

type agLine struct {
	file    string
	line    string
	content string
}

func search(c *cli.Context) {
    argLen := len(c.Args())
    var searchTerm string
    var searchPath string

    switch argLen {
    case 0:
        panic(1)
    case 1:
        searchTerm    = c.Args()[0]
        searchPath, _ = os.Getwd()
    case 2:
        searchTerm = c.Args()[0]
        searchPath = c.Args()[1]
    default:
        searchTerm = c.Args()[0]
        searchPath = c.Args()[1]
        // panic(1)
    }

	lines := executeCmd(searchTerm, searchPath)

	filePath := path.Join(home, shortcutFilename)
	f, err := os.Create(filePath)
	check(err)
	defer f.Close()

	color.Printf("@r[%2s]@{|} @b%5s@{|}  @g%s@{|}\n", "IDX", "Line", "Path")

	for i, line := range lines {

		if line == "" {
			break
		}

		lp := splitLine(line)
		l := agLine{file: lp[0], line: lp[1], content: lp[2]}
		s := color.Sprintf("@r[%2d]@{|} @b%5s@{|}  @g%s@{|} %s", i, l.line, l.file, l.content)
		// TODO: highlight search term using case insensitive regex
		fmt.Println(s)
		o := fmt.Sprint(l.line, " ", l.file, " ", l.content, "\n")

		w := bufio.NewWriter(f)
		_, err := w.WriteString(o)
		check(err)
		w.Flush()
	}
}
