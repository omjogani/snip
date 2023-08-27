package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/omjogani/snip/decoder"
)

func main() {
	color.Cyan("Welcome to Snip")
	color.Cyan("--------------------------")
	if len(os.Args) < 2 {
		for {
			var command string
			fmt.Printf("> ")
			fmt.Scan(&command)
			if command == "exit" {
				color.Green("Snip: Bye... 👋")
				os.Exit(0)
			}
			decoder.CommandDecoder(command)
		}
	} else if len(os.Args) >= 2 {
		decoder.CommandDecoder(os.Args[1])
	}
}
