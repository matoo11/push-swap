package shared

import (
	"reflect"
	"testing"
)

func TestParseArgs(t *testing.T) {
	cases := []struct {
		name    string
		args    []string
		want    []int
		wantErr bool
	}{
		{"empty", []string{}, []int{}, false},
		{"blank string", []string{"   "}, []int{}, false},
		{"single string", []string{"2 1 3"}, []int{2, 1, 3}, false},
		{"multiple args", []string{"2", "1", "3"}, []int{2, 1, 3}, false},
		{"negatives", []string{"-3 0 -1"}, []int{-3, 0, -1}, false},
		{"int32 bounds", []string{"-2147483648 2147483647"}, []int{-2147483648, 2147483647}, false},
		{"not a number", []string{"0 one 2 3"}, nil, true},
		{"duplicate", []string{"1 2 2 3"}, nil, true},
		{"overflow", []string{"2147483648"}, nil, true},
		{"underflow", []string{"-2147483649"}, nil, true},
		{"float", []string{"1.5"}, nil, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ParseArgs(tc.args)
			if (err != nil) != tc.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tc.wantErr)
			}
			if !tc.wantErr && !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestApply(t *testing.T) {
	cases := []struct {
		op           string
		wantA, wantB []int
	}{
		{"sa", []int{2, 1, 3}, []int{4, 5, 6}},
		{"sb", []int{1, 2, 3}, []int{5, 4, 6}},
		{"ss", []int{2, 1, 3}, []int{5, 4, 6}},
		{"pa", []int{4, 1, 2, 3}, []int{5, 6}},
		{"pb", []int{2, 3}, []int{1, 4, 5, 6}},
		{"ra", []int{2, 3, 1}, []int{4, 5, 6}},
		{"rb", []int{1, 2, 3}, []int{5, 6, 4}},
		{"rr", []int{2, 3, 1}, []int{5, 6, 4}},
		{"rra", []int{3, 1, 2}, []int{4, 5, 6}},
		{"rrb", []int{1, 2, 3}, []int{6, 4, 5}},
		{"rrr", []int{3, 1, 2}, []int{6, 4, 5}},
	}
	for _, tc := range cases {
		t.Run(tc.op, func(t *testing.T) {
			s := Stacks{StackA: []int{1, 2, 3}, StackB: []int{4, 5, 6}}
			if err := Apply(&s, tc.op); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(s.StackA, tc.wantA) || !reflect.DeepEqual(s.StackB, tc.wantB) {
				t.Fatalf("got A=%v B=%v, want A=%v B=%v", s.StackA, s.StackB, tc.wantA, tc.wantB)
			}
		})
	}
}

func TestApplyOnEmptyStacks(t *testing.T) {
	for op := range ValidOps {
		s := Stacks{StackA: []int{}, StackB: []int{}}
		if err := Apply(&s, op); err != nil {
			t.Fatalf("%s: %v", op, err)
		}
		if len(s.StackA) != 0 || len(s.StackB) != 0 {
			t.Fatalf("%s changed empty stacks", op)
		}
	}
}

func TestApplyInvalid(t *testing.T) {
	s := Stacks{StackA: []int{1}}
	if err := Apply(&s, "xx"); err != ErrInvalidOp {
		t.Fatalf("got %v, want ErrInvalidOp", err)
	}
}

func TestCompressToRanks(t *testing.T) {
	got := CompressToRanks([]int{40, -5, 100, 7})
	want := []int{2, 0, 3, 1}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}
