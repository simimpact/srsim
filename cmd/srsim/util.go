package main

import (
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

var (
	title = color.New(color.FgMagenta, color.Bold).PrintlnFunc()
	item  = color.New(color.FgYellow).SprintFunc()
)

func LogFile(path string) string {
	return filepath.Join(path, "log.gz")
}

func ResultFile(path string) string {
	return filepath.Join(path, "result.gz")
}

func LogExists(path string) bool {
	_, err := os.Stat(LogFile(path))
	return err == nil
}

func ResultExists(path string) bool {
	_, err := os.Stat(ResultFile(path))
	return err == nil
}

func concatMultipleSlices[T any](slices ...[]T) []T {
	var totalLen int

	for _, s := range slices {
		totalLen += len(s)
	}

	result := make([]T, totalLen)

	var i int

	for _, s := range slices {
		i += copy(result[i:], s)
	}

	return result
}
