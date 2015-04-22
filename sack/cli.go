package sack

import (
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
		cli.StringFlag{"flags, f", "-i", "flags to pass to ag"},
	}

	app.Commands = []cli.Command{
		{
			Name:  "init",
			Usage: "shell init script",
			Action: func(c *cli.Context) {
				shellInit(c)
			},
		},
		{
			Name:  "eval",
			Usage: "shell eval command to insert into .{zsh,bash}rc",
			Action: func(c *cli.Context) {
				shellEval(c)
			},
		},
	}
	app.Action = func(c *cli.Context) {

		debug("Execute:cli.Context %v", c)

		switch true {
		case c.Bool("edit"):
			edit(c)
		case c.Bool("search"):
			search(c)
		case c.Bool("print"):
			display()
		default:
			display()
		}
	}

	app.Run(os.Args)
}
