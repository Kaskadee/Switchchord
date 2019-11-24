package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Input represents an input class, which is able to read user input from files.
type Input struct {
	scanner *bufio.Scanner
}

// DefaultInput returns an instance of the Input class, which is able to read from stdin.
var DefaultInput = NewInput(os.Stdin)

// ReadString reads a string from the scanner.
// Line-break characters like \r and \n are not included.
// Returns an error, if end-of-file has been reached.
func (input *Input) ReadString() (string, error) {
	// Check if there is still data left to read.
	if !input.scanner.Scan() {
		return "", fmt.Errorf("EOF reached: %w", input.scanner.Err())
	}

	// Read next scanner token without linebreak characters.
	return input.scanner.Text(), input.scanner.Err()
}

// ReadInteger reads an integer from the scanner.
// If the input is not valid, the function will retry until a valid integer has been entered or EOF has been reached.
// Returns an error, if end-of-file has been reached.
func (input *Input) ReadInteger() (int, error) {
	// Do until valid input.
	for {
		// Check if there is still data left to read.
		if !input.scanner.Scan() {
			return -1, fmt.Errorf("EOF reached")
		}

		// Read next scanner token without linebreak characters.
		str := input.scanner.Text()
		if input.scanner.Err() != nil {
			return -1, input.scanner.Err()
		}

		// Try to cast string to integer.
		result, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("invalid input. try again!")
			continue
		}

		return result, nil
	}
}

// WaitForInput blocks until the next line-break character is read without returning the input.
func (input *Input) WaitForInput() {
	_ = input.scanner.Scan()
}

// NewInput creates a new instance of the function class and returns the Input pointer.
func NewInput(f *os.File) *Input {
	return &Input{scanner: bufio.NewScanner(f)}
}
