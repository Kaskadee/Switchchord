// +build !windows

package main

import (
	"fmt"
	"strings"
)

// SetConsoleTitle sets the title of the console window.
// On UNIX systems, a special command will be written to stdout.
// On Windows systems, syscall will be used to call a native windows function.
func SetConsoleTitle(title string) error {
	fmt.Print(strings.Replace("\033]0;{title}\007", title, 1))
	return nil
}
