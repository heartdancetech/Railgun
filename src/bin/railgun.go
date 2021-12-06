package main

import (
	"github.com/railgun-project/railgun/cmd"
	"os"
)

func main() {
	err := cmd.App.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
