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
