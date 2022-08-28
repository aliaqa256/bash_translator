package main

import (
	"fmt"
	"github.com/aliaqa256/bash_translator/cmd"
	"os"
)

func main() {
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
