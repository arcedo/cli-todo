package cli

import (
	"context"
	"io"

	"arcedo/cli-todo/internal/task"
)

func (c *CLI) runTask(ctx context.Context, args []string) {
	switch args[1] {
	case "new":
		tasks, err := c.taskService.Create(ctx, args[2:])
		if err != nil {
			println(c.errOut, err)
			return
		}
		printTasks(c.out, tasks)

	case "remove":
		ids, err := validateIDs(args[2:])
		if err != nil {
			println(c.errOut, err)
			return
		}
		affected, err := c.taskService.Delete(ctx, ids)
		if err != nil {
			println(c.errOut, err)
			return
		}
		printf(c.out, "%v of %v tasks successfully deleted\n", affected, len(ids))

	case "list":
		tasks, err := c.taskService.List(ctx, nil, task.Uncompleted)
		if err != nil {
			println(c.errOut, err)
			return
		}
		printTasks(c.out, tasks)

	case "complete":
		ids, err := validateIDs(args[2:])
		if err != nil {
			println(c.errOut, err)
			return
		}
		affected, err := c.taskService.Complete(ctx, ids)
		if err != nil {
			println(c.errOut, err)
			return
		}
		printf(c.out, "%v of %v tasks successfully completed\n", affected, len(ids))

	default:
		c.printUsage()
	}
}

func printTasks(out io.Writer, tasks []task.Task) {
	if len(tasks) == 0 {
		println(out, "No tasks found")
		return
	}
	for _, t := range tasks {
		status := " "
		if t.CompletedAt != nil {
			status = "x"
		}
		printf(out, "[%s] %d - %s\n", status, t.ID, t.Description)
	}
}
