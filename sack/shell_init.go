package sack

import (
    "fmt"
	"github.com/codegangsta/cli"
)

func shellInit(c *cli.Context) {
    sh := `
    set -e

    sack=$(which sack)

    alias S="${sack} -s"
    alias F="${sack} -e"
    `

    fmt.Println(sh)
}

