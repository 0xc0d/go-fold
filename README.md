# go-fold
[![PkgGoDev](https://pkg.go.dev/badge/github.com/0xc0d/go-fold)](https://pkg.go.dev/github.com/0xc0d/go-fold@v1.0.0?tab=doc)

A Go implementation of fold command (unix) around io.Reader.

# Why
Folding a string is not a hassle but having a stream in io.Reader could be a little bit challenging.
The idea for this fold reader started when we encountered with an issue in streaming data to a socket
with a protocol like SMTP. The maximum line length for SMTP mail is 990 characters.

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
