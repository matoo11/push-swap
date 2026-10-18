package checker

import (
	"bytes"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	cases := []struct {
		name       string
		args       []string
		stdin      string
		wantOut    string
		wantErr    string
		wantStatus int
	}{
		{"no args", nil, "", "", "", 0},
		{"invalid number", []string{"0 one 2 3"}, "", "", "Error\n", 1},
		{"duplicate", []string{"1 2 2 3"}, "", "", "Error\n", 1},
		{"invalid op", []string{"2 1"}, "sa\nxx\n", "", "Error\n", 1},
		{"audit KO", []string{"0 9 1 8 2 7 3 6 4 5"}, "sa\npb\nrrr\n\n", "KO\n", "", 0},
		{"audit OK", []string{"0 9 1 8 2"}, "pb\nra\npb\nra\nsa\nra\npa\npa\n\n", "OK\n", "", 0},
		{"sorted no ops", []string{"1 2 3"}, "", "OK\n", "", 0},
		{"unsorted no ops", []string{"2 1 3"}, "", "KO\n", "", 0},
		{"B not empty", []string{"1 2 3"}, "pb\n", "KO\n", "", 0},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var out, errOut bytes.Buffer
			status := Run(tc.args, strings.NewReader(tc.stdin), &out, &errOut)
			if status != tc.wantStatus || out.String() != tc.wantOut || errOut.String() != tc.wantErr {
				t.Fatalf("got status=%d out=%q err=%q", status, out.String(), errOut.String())
			}
		})
	}
}
