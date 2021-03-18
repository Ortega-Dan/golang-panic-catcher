package gpanic

import (
	"fmt"
	"os"
	"runtime/debug"
)

// Catches a panic (recovers from it) and prints it to stderr
func Catch() {
	if err := recover(); err != nil {
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, err)
		debug.PrintStack()
	}
}

// Catches a panic (recovers from it) and calls a given callback function,
// preserving a convenience relation through an arbitrary string identifier
// passed as the first argument
func CatchAndCall(relationId string, fcallback func(relationId string, errorMsg string, stacktrace string)) {
	if err := recover(); err != nil {
		errorMsg := fmt.Errorf("%v", err).Error()
		stackString := string(debug.Stack())

		fcallback(relationId, errorMsg, stackString)
	}
}

// Catches a panic (recovers from it) and prints it to a given file
// it can be used to print to stdout by passing os.Stdout as argument
func CatchToFile(file *os.File) {
	if err := recover(); err != nil {
		fmt.Fprintln(file, "")
		fmt.Fprintln(file, err)
		fmt.Fprintln(file, debug.Stack())
	}
}
