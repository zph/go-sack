package main

import (
	"fmt"
	"github.com/codegangsta/cli"
)

func shellInit(c *cli.Context) {
	sh := `
    sack=$(which sack)

    alias S="${sack} -s"
    alias F="${sack} -e"
    `

	fmt.Println(sh)
}

func shellEval(c *cli.Context) {
	sh := "eval \"$(sack init)\""
	fmt.Println(sh)
}

/*
// TODO: Add bash and zsh autocomplete
CREDIT: https://github.com/codegangsta/cli/blob/master/autocomplete/bash_autocomplete
    _cli_bash_autocomplete() {
        local cur prev opts base
        COMPREPLY=()
        cur="${COMP_WORDS[COMP_CWORD]}"
        prev="${COMP_WORDS[COMP_CWORD-1]}"
        opts=$( ${COMP_WORDS[@]:0:COMP_CWORD} --generate-bash-completion )
        COMPREPLY=( $(compgen -W "${opts}" -- ${cur}) )
        return 0
    }

    complete -F _cli_bash_autocomplete $PROG
*/
