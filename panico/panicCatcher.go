package panico

import (
	"fmt"
	"os"
	"runtime/debug"
)

func Catch() {
	if err := recover(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		debug.PrintStack()
	}
}
