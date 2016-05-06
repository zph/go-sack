package main

import (
	"github.com/wsxiaoys/terminal/color"
	"os"
	"path"
)

var home string = os.Getenv("HOME")

const agCmd string = "ag"
const ptCmd string = "pt"
const grepCmd string = "grep"

var filePath string = path.Join(home, shortcutFilename)
var termPath string = path.Join(home, searchTermFile)

const searchTermFile string = ".sack_searchterm"
const shortcutFilename string = ".sack_shortcuts"

var header string = color.Sprintf("@r[%2s]@{|} @b%5s@{|}  @g%s@{|}\n", "IDX", "Line", "Path")
