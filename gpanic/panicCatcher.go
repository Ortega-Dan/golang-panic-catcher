package gpanic

import (
	"fmt"
	"os"
	"runtime/debug"
)

// Catches a panic (recovers from it) and calls the given callback function,
// providing msg (error message) and trc (stack trace).
// This is intended mainly for usage withing closures deferment, in order to use the
// context of the caller in the callback fucntion.
// It can also be used to set named return values of the parent function (set the values to return), when used with an inline callback function, thanks to closures.
func CatchAndDo(fdo func(msg, trc string)) {
	if err := recover(); err != nil {
		msg, trc := GetPanicInfo(err)
		fdo(msg, trc)
	}
}

// Catches a panic (recovers from it) and calls the given callback function, preserving a convenience relation through an arbitrary string identifier passed as the first argument (id).
// It can also be used to set named return values of the parent function (set the values to return), when used with an inline callback function, thanks to closures.
func CatchAndCall(id string, fcall func(id string, msg string, trc string)) {
	if err := recover(); err != nil {
		errorMsg, stackString := GetPanicInfo(err)
		fcall(id, errorMsg, stackString)
	}
}

// Get the info from the given panic object
// (do not provide something different than a panic to this function)
func GetPanicInfo(panicErr interface{}) (msg string, trc string) {
	msg = fmt.Errorf("%v", panicErr).Error()
	trc = string(debug.Stack())
	return msg, trc
}

// Catches a panic (recovers from it) and prints it to stderr
func Catch() {
	if err := recover(); err != nil {
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, err)
		debug.PrintStack()
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
