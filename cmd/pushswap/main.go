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

// func Solve(values []int) ([]string, error) {
//     if len(values) == 0 || isSorted(values) {
//         return []string{}, nil
//     }

//     state := createStacks(values)
//     operations := []string{}

//     applyOperation := func(op string) {
//         apply(&state, op)
//         operations = append(operations, op)
//     }

//     if len(values) <= 6 {
//         smallOps, err := sortSmallStack(state)
//         if err != nil {
//             return nil, err
//         }
//         for _, op := range smallOps {
//             applyOperation(op)
//         }
//     } else {
//         chunkSort(&state, applyOperation)
//     }

//     if !isSorted(state.StackA) || len(state.StackB) != 0 {
//         return nil, fmt.Errorf("error: list not sorted")
//     }

//     return operations, nil
// }

