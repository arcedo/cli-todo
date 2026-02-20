// Package cli handles all the commands with heach service
package cli

import (
	"context"
	"io"

	"arcedo/cli-todo/internal/task"
)

type CLI struct {
	taskService *task.Service
	out         io.Writer
	errOut      io.Writer
}

func New(taskService *task.Service, out io.Writer, errOut io.Writer) *CLI {
	return &CLI{
		taskService,
		out,
		errOut,
	}
}

func (c *CLI) Run(ctx context.Context, args []string) {
	if len(args) < 2 {
		c.printUsage()
		return
	}

	// I did a separated function so that in the future
	// we want to add other entities is handled easly
	c.runTask(ctx, args)
}

func (c *CLI) printUsage() {
	println(c.out, "Usage: some usage here!")
}
