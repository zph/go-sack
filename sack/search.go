package sack

import (
	"bufio"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/wsxiaoys/terminal/color"
	"os"
	"os/exec"
	"path"
	"strings"
)

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
		searchTerm = c.Args()[0]
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
