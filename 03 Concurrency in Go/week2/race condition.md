```go
package main
import "fmt"

x := 1
func f1(){
  x++
}

func f2(){
  fmt.print(x)
}
```

If these two functions are excuted concurrently, the race condition would occur because these two functions are access the same resource and the output depends on non-deterministic ordering. 

Race condition is a condition when two or more thread can access the same resource at the same time. The output of the program in race condition became non-deterministic. The behavior of the same program could output two difference things.

