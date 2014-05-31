package sack

import (
	"fmt"
	"github.com/wsxiaoys/terminal/color"
	"strings"
)

func display() {
	lines := content()

	// Header
	color.Printf("@r[%2s]@{|} @b%5s@{|}  @g%s@{|}\n", "IDX", "Line", "Path")

	for i, line := range lines {
		li := strings.Split(line, " ")
		s := color.Sprintf("@r[%2d]@{|} @b%5s@{|}  @g%s@{|}", i, li[0], li[1])
		fmt.Println(s)
	}
}
