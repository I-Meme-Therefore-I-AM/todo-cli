package util

import (
	"fmt"
	"os"
	"syscall"
)

func LoadFile(file string) (*os.File, error) {
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)

	if err != nil {
		return nil, fmt.Errorf("failed to open file for reading: %s\n", err)
	}

	// Exclusive lock obtained on the file descriptor
	if err = syscall.Flock(int(f.Fd()), syscall.LOCK_EX); err != nil {
		return nil, err
	}
	return f, nil
}

func Close(f *os.File) error {
	syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
	return f.Close()
}
