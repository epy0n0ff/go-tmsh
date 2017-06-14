package main

import (
	"fmt"
	"os"

	"github.com/shiftky/go-tmsh/cmd/tmsh/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
