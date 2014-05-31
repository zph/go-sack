package sack

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

// TODO: add multiple edits sack -e 1 2 3
func edit(c *cli.Context) {
	lines := content()

	argLen := len(c.Args())

	var ind int

	switch argLen {
	case 0:
		ind = 0
	case 1:
		ind, _ = strconv.Atoi(c.Args()[0])
	default:
		panic(1)
	}

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
