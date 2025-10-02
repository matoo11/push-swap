package pushswap

import (
	"fmt"
	"strconv"
	"strings"

	"pshswp/shared"
)

// Solve produces a list of operations to sort values ascending (top first).
func Solve(values []int) ([]string, error) {
	n := len(values)
	if n == 0 || shared.IsAscending(values) {
		return []string{}, nil
	}

	state := shared.Stacks{StackA: append([]int(nil), values...), StackB: []int{}}
	operations := make([]string, 0, n*10)
	applyAndRecord := func(operation string) {
		if err := shared.Apply(&state, operation); err != nil {
			panic(err)
		}
		operations = append(operations, operation)
	}

	if n <= 6 {
		smallOps, err := solveSmallOpt(state)
		if err != nil {
			return nil, err
		}
		for _, operation := range smallOps {
			applyAndRecord(operation)
		}
	} else {
		chunkSort(&state, applyAndRecord)
	}

	if !shared.IsAscending(state.StackA) || len(state.StackB) != 0 {
		return nil, fmt.Errorf("internal error: not sorted")
	}
	return operations, nil
}

func bringIndexToTop(stacks *shared.Stacks, index int, applyOperation func(string)) {
	if index <= 0 {
		return
	}
	length := len(stacks.StackA)
	if index == 1 {
		applyOperation("sa")
		return
	}
	if index <= length/2 {
		for rotation := 0; rotation < index; rotation++ {
			applyOperation("ra")
		}
	} else {
		for rotation := 0; rotation < length-index; rotation++ {
			applyOperation("rra")
		}
	}
}

func chunkSort(stacks *shared.Stacks, applyOperation func(string)) {
	stacks.StackA = shared.CompressToRanks(stacks.StackA)
	n := len(stacks.StackA)
	if n == 0 {
		return
	}

	chunk := chooseChunkSize(n)
	nextLimit := chunk
	for len(stacks.StackA) > 0 {
		index := findIndexBelow(stacks.StackA, nextLimit)
		if index == -1 {
			nextLimit += chunk
			if nextLimit > n {
				nextLimit = n
			}
			continue
		}
		bringIndexToTop(stacks, index, applyOperation)
		applyOperation("pb")
		if len(stacks.StackB) > 1 {
			threshold := nextLimit - chunk/2
			if threshold < 0 {
				threshold = 0
			}
			if stacks.StackB[0] < threshold {
				applyOperation("rb")
			} else if stacks.StackB[0] < stacks.StackB[1] {
				applyOperation("sb")
			}
		}
	}

	for len(stacks.StackB) > 0 {
		index := indexOfMax(stacks.StackB)
		bringIndexToTopB(stacks, index, applyOperation)
		applyOperation("pa")
	}
}

func chooseChunkSize(n int) int {
	switch {
	case n <= 20:
		return 5
	case n <= 100:
		return 20
	case n <= 500:
		return 30
	default:
		return int(float64(n)*0.08) + 1
	}
}

func findIndexBelow(values []int, limit int) int {
	for index, value := range values {
		if value < limit {
			return index
		}
	}
	return -1
}

var allOps = []string{"sa", "sb", "ss", "pa", "pb", "ra", "rb", "rr", "rra", "rrb", "rrr"}

func solveSmallOpt(start shared.Stacks) ([]string, error) {
	total := len(start.StackA) + len(start.StackB)
	if total > 6 {
		return nil, fmt.Errorf("unsupported size for exhaustive search")
	}
	start = start.Clone()
	type node struct {
		stacks shared.Stacks
		log    []string
	}
	queue := []node{{stacks: start, log: nil}}
	visited := map[string]struct{}{serialize(start): {}}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if len(current.stacks.StackB) == 0 && shared.IsAscending(current.stacks.StackA) {
			return current.log, nil
		}
		for _, operation := range allOps {
			next := current.stacks.Clone()
			if err := shared.Apply(&next, operation); err != nil {
				return nil, err
			}
			key := serialize(next)
			if _, ok := visited[key]; ok {
				continue
			}
			visited[key] = struct{}{}
			opSequence := append(append([]string{}, current.log...), operation)
			queue = append(queue, node{stacks: next, log: opSequence})
		}
	}
	return nil, fmt.Errorf("no solution found")
}

func serialize(stacks shared.Stacks) string {
	var builder strings.Builder
	builder.WriteString("A:")
	for index, value := range stacks.StackA {
		if index > 0 {
			builder.WriteByte(',')
		}
		builder.WriteString(strconv.Itoa(value))
	}
	builder.WriteString("|B:")
	for index, value := range stacks.StackB {
		if index > 0 {
			builder.WriteByte(',')
		}
		builder.WriteString(strconv.Itoa(value))
	}
	return builder.String()
}

func indexOfMax(values []int) int {
	if len(values) == 0 {
		return -1
	}
	maxIndex := 0
	maxValue := values[0]
	for index := 1; index < len(values); index++ {
		if values[index] > maxValue {
			maxValue = values[index]
			maxIndex = index
		}
	}
	return maxIndex
}

func bringIndexToTopB(stacks *shared.Stacks, index int, applyOperation func(string)) {
	if index <= 0 {
		return
	}
	length := len(stacks.StackB)
	if index == 1 {
		applyOperation("sb")
		return
	}
	if index <= length/2 {
		for rotation := 0; rotation < index; rotation++ {
			applyOperation("rb")
		}
	} else {
		for rotation := 0; rotation < length-index; rotation++ {
			applyOperation("rrb")
		}
	}
}
