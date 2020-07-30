package tools

import (
	"fmt"
	"os"
)

func Error(format string, a ...interface{}) {
	fmt.Fprint(os.Stderr, fmt.Sprintf(format, a...)+"\n")
}

func Die(format string, a ...interface{}) {
	Error(format, a...)
	os.Exit(1)
}
