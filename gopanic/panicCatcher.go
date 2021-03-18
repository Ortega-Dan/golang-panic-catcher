package gopanic

import (
	"fmt"
	"os"
	"runtime/debug"
)

// Catches a panic (recovers from it) and prints it to stderr
func Catch() {
	if err := recover(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		debug.PrintStack()
	}
}
