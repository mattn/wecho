package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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
			t.Fatalf("want %q, but %q:", test.want, got)
		}
	}
}

func TestRun(t *testing.T) {
	dir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)
	_, err = exec.Command("go", "build", "-o", filepath.Join(dir, "wecho.exe"), "main.go").CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}
	wd, _ := os.Getwd()
	defer os.Chdir(wd)

	os.Chdir(dir)

	tests := []struct {
		cmd  string
		want string
	}{
		{`wecho/`, `/`},
		{`wecho /`, `/`},
		{`wecho.exe/`, `/`},
		{`wecho.exe /`, `/`},
	}

	for _, test := range tests {
		b, err := exec.Command("cmd", "/c", test.cmd).CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}
		got := strings.Split(string(b), "\n")[0]
		if test.want != got {
			t.Fatalf("want %q, but %q:", test.want, got)
		}
	}

}
