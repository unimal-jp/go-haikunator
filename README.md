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
	fmt.Println("[Hex].....", Hex(8))   // throbbing-shiken-ec706505
	fmt.Println("[Num].....", Num(8))   // long-kome-02652440
	fmt.Println("[Alpha]...", Alpha(8)) // yellow-do-rqXmCaPP
	fmt.Println("[Alnum]...", Alnum(8)) // lucky-mikan-jyP87IMY
}
```
