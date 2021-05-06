# golang-panic-catcher

Usage:

```golang
import "github.com/Ortega-Dan/golang-panic-catcher/gpanic"


func myFunction(){
  
  // This brings the message (msg) and stack trace (trc)
  // of the the panic, if a panic occurred.
  defer gpanic.CatchAndDo(func(msg, trc string) {
    // ...
    // do your catch logic
    // ...

  })
  
  // ...
  // do your stuff with confidence
  // ...
}
```

It can be used to return values from the parent function by using Golang 'named return values'.

It also offers:
* CatchAndCall (for sending the stacktrace to a callback by keeping some arbitrary relation by a received id)
* CatchToFile (for printing to a file or stdout).
* GetPanicInfo (which just returns the message and stacktrace of a panic object recovered by the user).
* Catch (simply catches a panic [recovers from it] and prints it to stderr)

:)
