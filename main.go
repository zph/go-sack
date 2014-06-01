package main

import (
	"github.com/zph/go-sack/sack"
)

/*
TODO:
- Add ability to specify alternate ag flags
- Make it use current dir for search if os.Args()[1] is absent
- Add term printing colors
- Improve columnar layout of printed text
*/

func main() {
	sack.Execute()
}
