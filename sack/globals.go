package sack

import (
	"os"
)

var home string = os.Getenv("HOME")

const agCmd string = "ag"
const flags string = "-i"

const shortcutFilename string = ".sack_shortcuts"
