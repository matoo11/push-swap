package main

import (
	"bufio"
	"fmt"
	"os"

	"pshswp/internal/pushswap"
	"pshswp/shared"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}

	nums, err := shared.ParseArgs(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		return
	}
	if len(nums) == 0 {
		return
	}

	ops, err := pushswap.Solve(nums)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		return
	}
	w := bufio.NewWriter(os.Stdout)
	for _, op := range ops {
		fmt.Fprintln(w, op)
	}
	w.Flush()
}
