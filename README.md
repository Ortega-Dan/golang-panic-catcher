# golang-panic-catcher

Usage:

```golang
import "github.com/Ortega-Dan/golang-panic-catcher/gopanic"


func myFunction(){
  defer gopanic.Catch()
  
  // ...
  // do your stuff with confidence
  // ...
}
```
