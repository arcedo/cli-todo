package cli

import (
	"fmt"
	"io"
	"strconv"
)

func validateIDs(ss []string) (ids []int, err error) {
	for _, s := range ss {
		id, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("failed to parse ID %v: %w", s, err)
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func println(out io.Writer, a ...any) {
	_, _ = fmt.Fprintln(out, a...)
}

func printf(out io.Writer, format string, a ...any) {
	_, _ = fmt.Fprintf(out, format, a...)
}
