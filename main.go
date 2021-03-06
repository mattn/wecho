package main

import (
	"fmt"
	"os"
	"strings"
	"syscall"
	"unsafe"
)

func argline(exe, s string) string {
	if len(s) == 0 {
		return ""
	}
	if strings.HasPrefix(s, `"`) && len(s) >= len(exe)+2 {
		s = s[1+len(exe)+1:]
	} else if len(s) >= len(exe) {
		s = s[len(exe):]
	}
	if len(s) > 0 && s[0] == ' ' {
		s = s[1:]
		if len(s) > 0 && s[0] == ' ' {
			s = s[1:]
		}
	}
	return s
}

func main() {
	p := syscall.GetCommandLine()
	s := syscall.UTF16ToString(*(*[]uint16)(unsafe.Pointer(&p)))
	fmt.Println(argline(os.Args[0], s))
}
