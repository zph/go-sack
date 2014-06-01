package sack

import (
	"os"
	"path"
)

var home string = os.Getenv("HOME")

const agCmd string = "ag"
const grepCmd string = "grep"

var filePath string = path.Join(home, shortcutFilename)
var termPath string = path.Join(home, searchTermFile)

const searchTermFile string = ".sack_searchterm"
const shortcutFilename string = ".sack_shortcuts"
