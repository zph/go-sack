package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/wsxiaoys/terminal/color"
)

var pwdRegex = regexp.MustCompile("^" + pwd + "/")

func display() {
	lines := content()

	dat, err := ioutil.ReadFile(termPath)
	check(err)

	// Header
	fmt.Print(header)

	for i, line := range lines {
		li := strings.SplitN(line, " ", 3)
		relativeFile := strings.Replace(li[1], pwd, "", 1)
		s := displayLines(string(dat), i, li[0], relativeFile, li[2])
		fmt.Printf(s)
	}
}

func absoluteToRelativePath(absPath string) string {
	return pwdRegex.ReplaceAllString(absPath, "")
}

func displayLines(term string, ind int, line string, file string, content string) string {
	str := fmt.Sprint("(?i)", "(", term, ")")
	reg, _ := regexp.Compile(str)
	hiContent := reg.ReplaceAllString(content, color.Sprintf("@{r!}$1"))
	path := absoluteToRelativePath(file)
	s := color.Sprintf("@r[%2d]@{|} @g%s@{|}:@b%s@{|}  %s", ind, path, line, hiContent)
	return s
}
