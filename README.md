# golang-panic-catcher

Usage:

```golang
import "github.com/Ortega-Dan/golang-panic-catcher/gpanic"


func myFunction(){
  defer gpanic.Catch()
  
  // ...
  // do your stuff with confidence
  // ...
}
```

It also offers CatchToFile (for printing to file or STDOUT),\
CatchToCall (for sending the stacktrace to a callback by keeping some arbitrary relation by a received id), and or to set return values (by setting parent function named return values) for inline callbacks thanks to closures,\
and GetPanicInfo (which just returns the message and stacktrace of a panic object recovered by the user)



:)
