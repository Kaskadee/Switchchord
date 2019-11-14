// +build !windows

package main

import (
	"fmt"
	"strings"
)

func SetConsoleTitle(title string) error {
	fmt.Print(strings.Replace("\033]0;{title}\007", title, 1))
	return nil
}