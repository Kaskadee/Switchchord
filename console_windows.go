// +build windows

package main

import (
	"syscall"
	"unsafe"
)

// Sets the title of the console window.
// On UNIX systems, a special command will be written to stdout.
// On Windows systems, syscall will be used to call a native windows function.
func SetConsoleTitle(title string) error {
	// Load kernel library which contains function to change console title.
	handle, err := syscall.LoadLibrary("kernel32.dll")
	if err != nil {
		return err
	}

	// Free library handle after we are finished.
	defer syscall.FreeLibrary(handle)

	// Get address of SetConsoleTitleW function.
	proc, err := syscall.GetProcAddress(handle, "SetConsoleTitleW")
	if err != nil {
		return err
	}

	// Convert golang string to UTF16 pointer.
	utfPtr, err := syscall.UTF16PtrFromString(title)
	if err != nil {
		return err
	}

	// Convert UTF16 pointer to unsafe pointer.
	titlePtr := uintptr(unsafe.Pointer(utfPtr))
	result, _, errno := syscall.Syscall(proc, 1, titlePtr, 0, 0)

	// Check if function call was successful.
	if errno != 0 {
		err = errno
		return err
	}

	// If result is zero, function call was not successful.
	if result == 0 {
		return syscall.GetLastError()
	}

	return nil
}
