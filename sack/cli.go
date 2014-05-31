package sack

import (
	"github.com/codegangsta/cli"
    "fmt"
    "os"
)

func Execute() {
	checkState()

	app := cli.NewApp()
	app.Name = "Sack"
	app.Usage = "sack [searchterm] [optional directory]"
	app.Version = Version()
	app.Flags = []cli.Flag{
		cli.BoolFlag{"edit, e", "edit a given shortcut"},
		cli.BoolFlag{"search, s", "search-ack/ag it"},
		cli.BoolFlag{"print, p", "display existing shortcuts"},
		cli.BoolFlag{"debug, d", "show all the texts"},
	}

	app.Action = func(c *cli.Context) {

		if c.Bool("debug") {
			fmt.Printf("Context %#v\n", c)
		}

		if c.Bool("edit") {
			edit(c)
		} else if c.Bool("search") {
			search(c)
		} else if c.Bool("print") || true {
			display()
		}
	}
	app.Run(os.Args)
}
