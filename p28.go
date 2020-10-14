package main

import (
	"os"
	"path/filepath"
	"regexp"
)

const BUF_SIZE = 1000

func find(path string) <-chan string {
	out := make(chan string, BUF_SIZE)
	go func() {
		filepath.Walk(path, func(file string, info os.FileInfo, err error) error {
			out <- file
			return nil
		})
		close(out)
	}()
	return out
}

func grep(pattern string, in <-chan string) <-chan string {
	out := make(chan string, cap(in))
	go func() {
		regex, err := regexp.Compile(pattern)
		if err != nil {
			panic(err)
		}

		for file := range in {
			if regex.MatchString(file) {
				out <- file
			}
		}

		close(out)
	}()
	return out
}
