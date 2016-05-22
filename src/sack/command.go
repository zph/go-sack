package main

import (
	. "github.com/tj/go-debug"
)

/*
TODO:
- Add ability to specify alternate ag flags
- Make it use current dir for search if os.Args()[1] is absent
*/

var debug = Debug("sack:main")

func main() {
	debug("starting main")
	execute()
}
