# go-haikunator

Heroku like word generator.

# Usage

```go
package main

import (
	"fmt"
	"github.com/unimal-jp/go-haikunator"
)

func main() {
	fmt.Println("[Hex].....", haikunator.Hex(8))   // throbbing-shiken-ec706505
	fmt.Println("[Num].....", haikunator.Num(8))   // long-kome-02652440
	fmt.Println("[Alpha]...", haikunator.Alpha(8)) // yellow-do-rqXmCaPP
	fmt.Println("[Alnum]...", haikunator.Alnum(8)) // lucky-mikan-jyP87IMY
}
```
