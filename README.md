# go-fold
A Go implementation of fold command (unix) around io.Reader.


# Install
```
go get github.com/0xc0d/go-fold
```

# Use
```go
package main

import (
  "io"
  "os"
  "strings"
  
  "github.com/0xc0d/go-fold"
)

func main() {
  s := strings.Repeat("0", 20)
  r := strings.NewReader(s)
  foldReader := fold.NewReader(r, 7)
  
  io.Copy(os.Stdout, foldReader)
}

// Output:
// 0000000
// 0000000
// 000000

```
