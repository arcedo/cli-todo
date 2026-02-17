// Package cli handles all the commands with heach service
package cli

import (
	"fmt"

	"arcedo/cli-todo/internal/task"
)

func Run(args []string, taskService *task.Service) {
	if len(args) < 2 {
		printUsage()
		return
	}

	// I did a separated function so that in the future
	// we want to add other entities is handled easly
	runTask(args, taskService)
}

func runTask(args []string, s *task.Service) {
	switch args[1] {
	case "new":
		return
	case "list":
		return
	}
}

func printUsage() {
	fmt.Println("some usage here!")
}
