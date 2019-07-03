# Everything
Everything SDK caller for golang

You need Everything64.dll from Everything SDK ( http://voidtools.com/downloads/ )

```go
package main

import (
	"fmt"
	"path/filepath"

	Et "github.com/jof4002/Everything"
)

func main() {

	Et.Walk("*.go", func(path string, info Et.FileInfo, err error) error {
		name := filepath.Base(path)
		size := info.Size()
		tmod := info.ModTime()
		fmt.Println(name, size, tmod.Format("2006-01-02 15:04"))
		return nil
	})
}
```

#### Windows x86 and 64bit

`everything_windows_amd64.go`   works at `GOARCH=amd64`.

`everything_windows_386.go`  works at `GOARCH=386`.

These will be processed automatically at `go run` and `go build`.

By the way,You can use `go env` to check the `GOOS` and `GOARCH`.

