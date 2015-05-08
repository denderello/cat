package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	if len(os.Args) < 2 {
		syscall.Write(syscall.Stderr, []byte("filename parameter is missing.\n"))
		os.Exit(1)
	}

	cat(os.Args[1])
}

func cat(filename string) {
	fd, err := syscall.Open(filename, syscall.O_RDONLY, syscall.S_IRUSR|syscall.S_IRGRP|syscall.S_IROTH)
	if err != nil {
		msg := fmt.Sprintf("error on Open: %#v\n", err)
		if err == syscall.ENOENT {
			msg = fmt.Sprintf("'%s' does not exist.\n", filename)
		}
		syscall.Write(syscall.Stderr, []byte(msg))
		os.Exit(1)
	}

	stat := &syscall.Stat_t{}
	err = syscall.Fstat(fd, stat)
	if err != nil {
		syscall.Write(syscall.Stderr, []byte(fmt.Sprintf("error on Fstat: %#v\n", err)))
		os.Exit(1)
	}

	filetype := stat.Mode & syscall.S_IFMT
	if filetype != syscall.S_IFREG {
		syscall.Write(syscall.Stderr, []byte(fmt.Sprintf("'%s' is not a file.\n", filename)))
		os.Exit(1)
	}

	done := false
	readBuffer := make([]byte, 4000)
	for done != true {
		bytesRead, err := syscall.Read(fd, readBuffer)
		if err != nil {
			syscall.Write(syscall.Stderr, []byte(fmt.Sprintf("error on Read: %#v\n", err)))
			os.Exit(1)
		}

		if bytesRead == 0 {
			done = true
			continue
		}
		syscall.Write(syscall.Stdout, readBuffer)
	}

	err = syscall.Close(fd)
	if err != nil {
		syscall.Write(syscall.Stderr, []byte(fmt.Sprintf("error on Close: %#v\n", err)))
	}
}
