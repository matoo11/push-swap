package checker

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"pshswp/shared"
)

// Run executes the checker logic. Returns exit code.
func Run(args []string, stdin io.Reader, stdout, stderr io.Writer) int {
	// Parse numbers
	nums, err := shared.ParseArgs(args)
	if err != nil {
		if len(nums) == 0 && len(args) == 0 {
			return 0
		}
		fmt.Fprintln(stderr, "Error")
		return 1
	}
	if len(nums) == 0 {
		// no output per spec
		return 0
	}

	stacks := shared.Stacks{StackA: append([]int(nil), nums...), StackB: []int{}}

	// Read ops from stdin
	scanner := bufio.NewScanner(stdin)
	// allow longer lines than default 64K? Not necessary for short ops
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			// tolerate blank lines (common with echo -e "...\n")
			continue
		}
		if _, ok := shared.ValidOps[line]; !ok {
			fmt.Fprintln(stderr, "Error")
			return 1
		}
		if err := shared.Apply(&stacks, line); err != nil {
			fmt.Fprintln(stderr, "Error")
			return 1
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(stderr, "Error")
		return 1
	}

	if len(stacks.StackB) == 0 && shared.IsAscending(stacks.StackA) {
		fmt.Fprintln(stdout, "OK")
	} else {
		fmt.Fprintln(stdout, "KO")
	}
	return 0
}

// Main is provided for convenience if used as a package executable
func Main() {
	code := Run(os.Args[1:], os.Stdin, os.Stdout, os.Stderr)
	os.Exit(code)
}
