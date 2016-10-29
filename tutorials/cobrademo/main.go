package main

import (
	"fmt"
	"os"

	"github.com/markTward/tutorials/cobrademo/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
