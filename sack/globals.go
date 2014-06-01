package sack

import (
	"os"
)

var home string = os.Getenv("HOME")

const agCmd string = "ag"
const grepCmd string = "grep"

const shortcutFilename string = ".sack_shortcuts"
