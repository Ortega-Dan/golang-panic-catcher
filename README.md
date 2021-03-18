# golang-panic-catcher

Usage:

```golang
import "github.com/Ortega-Dan/golang-panic-catcher/panic"


func myFunction(){
  defer panic.Catch()
  
  // ...
  // do your stuff with confidence
  // ...
}
```
