# go-fold
[![PkgGoDev](https://pkg.go.dev/badge/github.com/0xc0d/go-fold)](https://pkg.go.dev/github.com/0xc0d/go-fold?tab=doc)

A Go implementation of fold command (unix) around io.Reader.

# Why
Folding a string is not a hassle but a folding a stream in io.Reader 
could be a little bit challenging. The idea for this fold reader started 
when we have encountered with an issue in streaming data as an email
message. According to RFC 822:
> Lines of characters in the body MUST be limited to 998 characters, 
and SHOULD be limited to 78 characters, excluding the CRLF.

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
