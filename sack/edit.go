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

func edit(c *cli.Context) {
	lines := content()

    argLen := len(c.Args())

    var ind int

    if argLen == 0 {
        ind = 0
    } else {
        ind, _ = strconv.Atoi(c.Args()[0])
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
