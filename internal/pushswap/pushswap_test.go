package pushswap

import (
	"math/rand"
	"testing"

	"pshswp/shared"
)

// verify replays ops on values and reports whether the result is sorted with B empty.
func verify(t *testing.T, values []int, ops []string) {
	t.Helper()
	s := shared.Stacks{StackA: append([]int(nil), values...), StackB: []int{}}
	for _, op := range ops {
		if err := shared.Apply(&s, op); err != nil {
			t.Fatalf("invalid op %q", op)
		}
	}
	if len(s.StackB) != 0 || !shared.IsAscending(s.StackA) {
		t.Fatalf("not sorted for %v: A=%v B=%v", values, s.StackA, s.StackB)
	}
}

func permutations(values []int) [][]int {
	if len(values) <= 1 {
		return [][]int{append([]int(nil), values...)}
	}
	var result [][]int
	for i := range values {
		rest := append(append([]int(nil), values[:i]...), values[i+1:]...)
		for _, p := range permutations(rest) {
			result = append(result, append([]int{values[i]}, p...))
		}
	}
	return result
}

func randomDistinct(r *rand.Rand, n int) []int {
	seen := map[int]bool{}
	values := make([]int, 0, n)
	for len(values) < n {
		v := r.Intn(200000) - 100000
		if !seen[v] {
			seen[v] = true
			values = append(values, v)
		}
	}
	return values
}

func TestSortedInputProducesNothing(t *testing.T) {
	for _, values := range [][]int{{}, {5}, {0, 1, 2, 3, 4, 5}, {-3, 7, 42}} {
		ops, err := Solve(values)
		if err != nil || len(ops) != 0 {
			t.Fatalf("%v: got %v, %v", values, ops, err)
		}
	}
}

func TestAuditExample(t *testing.T) {
	values := []int{2, 1, 3, 6, 5, 8}
	ops, err := Solve(values)
	if err != nil {
		t.Fatal(err)
	}
	verify(t, values, ops)
	if len(ops) >= 9 {
		t.Fatalf("got %d ops, want < 9", len(ops))
	}
}

func TestAllPermutationsUpToSix(t *testing.T) {
	limits := map[int]int{2: 1, 3: 2, 4: 8, 5: 11, 6: 15}
	for n := 2; n <= 6; n++ {
		base := make([]int, n)
		for i := range base {
			base[i] = i * 3
		}
		for _, values := range permutations(base) {
			ops, err := Solve(values)
			if err != nil {
				t.Fatal(err)
			}
			verify(t, values, ops)
			if len(ops) > limits[n] {
				t.Fatalf("%v: %d ops, limit %d", values, len(ops), limits[n])
			}
		}
	}
}

func TestRandomHundred(t *testing.T) {
	r := rand.New(rand.NewSource(1))
	for i := 0; i < 500; i++ {
		values := randomDistinct(r, 100)
		ops, err := Solve(values)
		if err != nil {
			t.Fatal(err)
		}
		verify(t, values, ops)
		if len(ops) >= 700 {
			t.Fatalf("run %d: %d ops, want < 700", i, len(ops))
		}
	}
}

func TestRandomOtherSizes(t *testing.T) {
	r := rand.New(rand.NewSource(2))
	for _, n := range []int{7, 10, 20, 50, 500} {
		values := randomDistinct(r, n)
		ops, err := Solve(values)
		if err != nil {
			t.Fatal(err)
		}
		verify(t, values, ops)
	}
}
