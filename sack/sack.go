package sack

import (
	// "bufio"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/wsxiaoys/terminal/color"
	"os"
	"os/exec"
	// "path"
	"strconv"
	"strings"
	"syscall"
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
