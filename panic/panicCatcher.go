package panic

import (
	"fmt"
	"runtime/debug"
)

func Catch() {
	if err := recover(); err != nil {
		fmt.Println(err)
		debug.PrintStack()
	}
}
