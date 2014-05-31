package sack

import (
    "fmt"
	"github.com/codegangsta/cli"
)

func shellInit(c *cli.Context) {
    sh := `
    sack=$(which sack)

    alias S="${sack} -s"
    alias F="${sack} -e"
    `

    fmt.Println(sh)
}

func shellEval(c * cli.Context){
    sh := "eval \"$(sack init)\""
    fmt.Println(sh)

}
