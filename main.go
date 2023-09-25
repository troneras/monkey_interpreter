package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	// Get the current user
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	// Print the welcome message
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")

	// Start the REPL
	repl.Start(os.Stdin, os.Stdout)
}
