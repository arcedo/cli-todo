package cli

import (
	"context"
	"fmt"
	"io"
	"text/tabwriter"

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
		IDs, filter, err := manageListArgs(args[2:])
		if err != nil {
			println(c.errOut, err)
			c.printUsage()
			return
		}
		tasks, err := c.taskService.List(ctx, IDs, filter)
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

func manageListArgs(args []string) ([]int, task.ListFilter, error) {
	if len(args) == 0 {
		return nil, task.Uncompleted, nil // default list uncompleted
	}
	IDs, err := validateIDs(args)
	if err != nil {
		switch args[0] {
		case "all":
			return nil, task.All, nil
		case "completed":
			return nil, task.Completed, nil
		case "uncompleted":
			return nil, task.Uncompleted, nil
		case "removed":
			return nil, task.Removed, nil
		default:
			return nil, task.All, fmt.Errorf("invalid arguments: %s", args[0])
		}
	}
	return IDs, task.IDs, nil
}

func printTasks(out io.Writer, tasks []task.Task) {
	if len(tasks) == 0 {
		println(out, "No tasks found")
		return
	}

	w := tabwriter.NewWriter(out, 0, 0, 2, ' ', 0)

	println(w, "ID\tStatus\tCreated At\tDescription")
	println(w, "------------------------------------------------")

	for _, t := range tasks {
		status := "·"
		if t.CompletedAt != nil {
			status = "✓"
		}

		printf(
			w,
			"%d\t%s\t%s\t%s\n",
			t.ID,
			status,
			t.CreatedAt.Format("02/01/2006 15:04"),
			t.Description,
		)
	}

	w.Flush()
}
