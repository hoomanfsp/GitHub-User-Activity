package main

import (
	"fmt"
	"github-activity/process"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a variable.")
		return
	}

	// Access the variable
	user_name := os.Args[1]
	process.Start(user_name)
}
