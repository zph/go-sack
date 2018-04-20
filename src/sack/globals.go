package main

import (
	"os"
	"path"

	"github.com/wsxiaoys/terminal/color"
)

var home string = os.Getenv("HOME")

var pwd = os.Getenv("PWD")

const agCmd string = "ag"
const ptCmd string = "pt"
const rgCmd string = "rg"
const grepCmd string = "grep"

var filePath string = path.Join(home, shortcutFilename)
var termPath string = path.Join(home, searchTermFile)

const searchTermFile string = ".sack_searchterm"
const shortcutFilename string = ".sack_shortcuts"

var header string = color.Sprintf("@r[%2s]@{|} @b%5s@{|}  @g%s@{|}\n", "IDX", "Line", "Path")
