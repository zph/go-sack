package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func execute() {
	checkState()

	app := cli.NewApp()
	app.Name = "Sack"
	app.Usage = "sack [searchterm] [optional directory]"
	app.Version = Version()
	app.Flags = []cli.Flag{
		cli.BoolFlag{Name: "edit, e", Usage: "edit a given shortcut"},
		cli.BoolFlag{Name: "search, s", Usage: "search-ack/ag it"},
		cli.BoolFlag{Name: "print, p", Usage: "display existing shortcuts"},
		cli.BoolFlag{Name: "debug, d", Usage: "show all the texts"},
		cli.StringFlag{Name: "flags, f", Usage: "flags to pass to ag"},
	}

	app.Commands = []cli.Command{
		{
			Name:  "init",
			Usage: "shell init script",
			Action: func(c *cli.Context) error {
				shellInit(c)
				return nil
			},
		},
		{
			Name:  "eval",
			Usage: "shell eval command to insert into .{zsh,bash}rc",
			Action: func(c *cli.Context) error {
				shellEval(c)
				return nil
			},
		},
	}
	app.Action = func(c *cli.Context) error {

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
		return nil
	}

	app.Run(os.Args)
}
