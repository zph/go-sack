package sack

import (
	"bufio"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/wsxiaoys/terminal/color"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func executeCmd(term string, path string, flags string) []string {

	var lines []string
	_, err := exec.LookPath(agCmd)
	if err == nil {
		lines = agSearch(term, path, flags)
	} else {
		lines = grepSearch(term, path, flags)
	}

	return lines
}

func agSearch(term string, path string, flags string) []string {
	bin := getPath(agCmd)

	// Blows up if flags == "" without this conditional
	var cmdOut []byte
	var err error
	if flags == "" {
		cmdOut, err = exec.Command(bin, term, path).Output()
	} else {
		cmdOut, err = exec.Command(bin, flags, term, path).Output()
	}
	check(err)

	return outputLines(cmdOut)
}

func getPath(bin string) string {
	agBin, err := exec.LookPath(bin)
	check(err)
	return agBin
}

func grepSearch(term string, path string, flags string) []string {
	bin := getPath(grepCmd)

	flag1 := "-ir"
	flag2 := "--line-number"

	cmdOut, err := exec.Command(bin, flag1, flag2, term, path).Output()
	check(err)

	return outputLines(cmdOut)
}

func outputLines(b []byte) []string {
	return strings.Split(string(b), "\n")
}

type agLine struct {
	file    string
	line    string
	content string
}

func setTermPath(c *cli.Context) (string, string) {
	argLen := len(c.Args())
	var term string
	var path string
	switch argLen {
	case 0:
		panic(1)
	case 1:
		term = c.Args()[0]
		path, _ = os.Getwd()
	case 2:
		term = c.Args()[0]
		path = c.Args()[1]
	default:
		term = c.Args()[0]
		path = c.Args()[1]
	}
	return term, path
}

func search(c *cli.Context) {

	term, searchPath := setTermPath(c)

	lines := executeCmd(term, searchPath, c.String("flags"))

	t := []byte(term)
	err := ioutil.WriteFile(termPath, t, 0644)
	check(err)

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

		s := displayLines(term, i, l.line, l.file, l.content)
		fmt.Println(s)
		o := fmt.Sprint(l.line, " ", l.file, " ", l.content, "\n")

		w := bufio.NewWriter(f)
		_, err := w.WriteString(o)
		check(err)
		w.Flush()
	}
}
