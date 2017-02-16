package main

import (
	"testing"
)

func TestArgline(t *testing.T) {
	tests := []struct {
		exe  string
		line string
		want string
	}{
		{`foo`, `foo/`, `/`},
		{`foo.exe`, `foo.exe/`, `/`},
		{`foo.exe`, `foo.exe /`, `/`},
	}

	for _, test := range tests {
		got := argline(test.exe, test.line)
		if test.want != got {
			t.Fatalf("want %v, but %v:", test.want, got)
		}
	}
}
