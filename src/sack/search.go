package main

import (
	"bufio"
	"fmt"
	"github.com/codegangsta/cli"
	. "github.com/tj/go-debug"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
)

type Line struct {
	file    string
	line    string
	content string
}

func (l *Line) truncatedContent() string {
	// Lines longer than maxLength are often junk results
	line := l.content
	actualContentLen := len(line)
	maxLength := 200
	var contentLen int
	if actualContentLen > maxLength {
		contentLen = maxLength
	} else {
		contentLen = len(line)
	}

	return line[:contentLen]
}

func (l *Line) display(s SearchArgs, i int) string {
	return displayLines(s.term, i, l.line, l.file, l.truncatedContent())
}

func (l *Line) toString() string {
	return fmt.Sprint(l.line, " ", l.file, " ", l.truncatedContent(), "\n")
}

type SearchArgs struct {
	cmd   string
	term  string
	path  string
	flags []string
}

func (s *SearchArgs) bin() string {
	b, err := exec.LookPath(s.cmd)
	check(err)
	return b
}

func executeCmd(s SearchArgs) []string {
	var debug = Debug("sack:search")

	if _, err := exec.LookPath(ptCmd); err == nil {
		s.cmd = ptCmd
	} else if _, err := exec.LookPath(agCmd); err == nil {
		s.cmd = agCmd
	} else {
		s.cmd = grepCmd
		s.flags = []string{"--ir", "--line-number"}
	}

	debug("Executing search: %+v", s)
	return genericSearch(s)
}

func scanner(dst *[]string, stdout *io.ReadCloser) {
	reader := bufio.NewReader(*stdout)
	for {
		t, err := reader.ReadString('\n')
		if err == nil {
			*dst = append(*dst, t)
		} else if err == io.EOF {
			break
		} else {
			log.Printf("Error is: %v", err)
			break
		}
	}
}

func cmdAndParse(bin string, ax []string) []string {
	var dst []string

	cmd := exec.Command(bin, ax...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	go scanner(&dst, &stdout)

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}

	return dst
}

func genericSearch(a SearchArgs) []string {
	ax := []string{a.term, a.path}
	if len(a.flags) != 0 {
		ax = append(ax, a.flags...)
	}
	return cmdAndParse(a.bin(), ax)
}

func setTermPath(c *cli.Context) SearchArgs {
	args := c.Args()
	argLen := len(args)
	debug("setTermPath:argLen: %v", argLen)
	switch argLen {
	case 0:
		panic(1)
	case 1:
		pwd, _ := os.Getwd()
		return SearchArgs{term: args[0], path: pwd}
	default:
		return SearchArgs{term: args[0], path: args[1]}
	}
}

func getFlags(fx string) []string {
	if fx == "" {
		return []string{}
	} else {
		return strings.Split(fx, " ")
	}
}

func printer(c chan string, wg *sync.WaitGroup) {
	for {
		if v, ok := <-c; ok == false {
			fmt.Println(v)
			wg.Done()
			break
		} else {
			fmt.Println(v)
		}
	}
}

func displayAndWriteLines(s SearchArgs, lines []string) {
	if len(lines) == 0 {
		var debug = Debug("sack:search:error")
		debug("nolines:lines: %v", lines)
		os.Exit(1)
		return
	}

	t := []byte(s.term)
	err := ioutil.WriteFile(termPath, t, 0644)
	check(err)

	f, err := os.Create(filePath)
	check(err)
	defer f.Close()

	c1 := make(chan string)
	var wg sync.WaitGroup

	wg.Add(1)
	go printer(c1, &wg)

	c1 <- header

	for i, line := range lines {

		if line == "" {
			break
		}

		a, b, c := splitLine(line)

		l := Line{a, b, c}

		c1 <- l.display(s, i)

		w := bufio.NewWriter(f)
		_, err := w.WriteString(l.toString())
		check(err)
		w.Flush()
	}

	close(c1)
	// Necessary or some output is lost due to premature exit
	wg.Wait()
}

func search(c *cli.Context) {
	s := setTermPath(c)
	s.flags = getFlags(c.String("flags"))
	lines := executeCmd(s)
	displayAndWriteLines(s, lines)
}
