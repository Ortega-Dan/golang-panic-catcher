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

It also offers CatchToFile (for printing to file or STDOUT)\
and CatchToCall (for sending the stacktrace to a callback by keeping some arbitrary relation by a received id)



:)
