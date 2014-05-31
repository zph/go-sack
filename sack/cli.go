package sack

import (
	"fmt"
	"github.com/codegangsta/cli"
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

    app.Commands = []cli.Command{
        {
            Name:      "init",
            Usage:     "shell init script",
            Action: func(c *cli.Context) {
                shellInit(c)
            },
        },
    }
	app.Action = func(c *cli.Context) {

		if c.Bool("debug") {
			fmt.Printf("Context %#v\n", c)
		}

        switch true { 
        case c.Bool("edit"):
			edit(c)
        case c.Bool("search"):
            search(c)
        case c.Bool("print"):
            display()
        default:
            search(c)
        }
	}

	app.Run(os.Args)
}
