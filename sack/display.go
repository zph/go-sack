package sack

import (
	"fmt"
	"github.com/wsxiaoys/terminal/color"
	"io/ioutil"
	"regexp"
	"strings"
)

func display() {
	lines := content()

	dat, err := ioutil.ReadFile(termPath)
	check(err)

	// Header
	fmt.Print(header)

	for i, line := range lines {
		li := strings.SplitN(line, " ", 3)
		s := displayLines(string(dat), i, li[0], li[1], li[2])
		fmt.Println(s)
	}
}

func displayLines(term string, ind int, line string, file string, content string) string {
	str := fmt.Sprint("(?i)", "(", term, ")")
	reg, _ := regexp.Compile(str)
	hiContent := reg.ReplaceAllString(content, color.Sprintf("@{r!}$1"))
	s := color.Sprintf("@r[%2d]@{|} @b%5s@{|}  @g%s@{|} %s", ind, line, file, hiContent)
	return s
}
