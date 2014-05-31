package sack

import (
	"os"
)

var home string = os.Getenv("HOME")

const shortcutFilename string = ".sack_shortcuts"
