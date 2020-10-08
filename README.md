# uninames

GO haikunator. Heroku like word generator

# Usage

```go
package main

import (
  "fmt"
  "github.com/unimal-jp/uninames"
)


func main() {
	fmt.Println("[Hex].....", Hex(8))
	fmt.Println("[Num].....", Num(8))
	fmt.Println("[Alpha]...", Alpha(8))
	fmt.Println("[Alnum]...", Alnum(8))
}
```
