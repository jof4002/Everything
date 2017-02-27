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
