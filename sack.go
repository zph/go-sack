package main

// Taken straight from https://github.com/jmhodges/jsonpp/blob/master/jsonpp.go
import (
  // "bufio"
  // "bytes"
  "flag"
  "fmt"
  // "io"
  "os"
  "io/ioutil"
  "strings"
)

var newline = []uint8("\n")
var help = flag.Bool("help", false, "help")

func main() {
  flag.Parse()
  if *help {
    cmd := os.Args[0]

    // if cmd[0:2] == "./" {
    //   cmd = cmd[2:]
    // }
    fmt.Fprintf(os.Stderr, "Usage: "+cmd+" [file]"+"\n")
    fmt.Fprintf(os.Stderr, "   or: $COMMAND | "+cmd+"\n")
    os.Exit(0)
  }

  // filename := os.Getenv("SACK_SHORTCUTS_FILE")
  // if filename == "" {
  //   filename = "~/.sack_shortcuts"
  // }

  var exitStatus = 0
  filename := os.Args[1]

  // file, err := os.OpenFile(filename, os.O_RDONLY, 0)

  content, err := ioutil.ReadFile(filename)

  if err != nil {
    exitStatus = 1
      //Do something
  }

  lines := strings.Split(string(content), "\n")

  // p := []string{"", "", ""}
  l := strings.Split(lines[0], ":")

  type Line struct {
    filename string
    line_number string
    excerpt string
  }

  line := Line{l[0], l[1], l[2]}
  fmt.Println(line.excerpt)

      // status := processFile(file, indent)
      // if status > 0 {
      //   exitStatus = status
      // }
    // }
  // } else {
    // status := processFile(os.Stdin, indent)
    // if status > 0 {
      // exitStatus = status
    // }
  // }
  os.Exit(exitStatus)
}
