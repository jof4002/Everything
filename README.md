# Everything
Everything SDK caller for golang

This Go package provides an interface to the Everything search engine SDK by voidtools. It allows Go applications to perform fast file and folder searches on Windows systems.

**Requirements:**
- Windows OS (For other OS, the package builds but returns an error at runtime).
- `Everything64.dll` (for 64-bit) or `Everything32.dll` (for 32-bit) in your executable's directory.

**Usage Example:**

```go
package main

import (
	"fmt"
	"log" // Added for error handling
	"path/filepath"

	Et "github.com/jof4002/Everything"
)

func main() {
	// Check if the platform is supported
	err := Et.SetSearch("test")
	if err != nil {
		// Check for the specific unsupported platform error
		fmt.Printf("Notice: %v\n", err)
		return
	}	

	if err := Et.SetSearch("*.go"); err != nil {
		log.Fatalf("Error setting search: %v", err)
	}
	Et.SetRequestFlags(Et.EVERYTHING_REQUEST_FILE_NAME | Et.EVERYTHING_REQUEST_PATH | Et.EVERYTHING_REQUEST_SIZE | Et.EVERYTHING_REQUEST_DATE_MODIFIED)
	
	ok, err := Et.Query(true)
	if err != nil {
		log.Fatalf("Error querying Everything: %v", err)
	}
	if !ok {
		log.Fatal("Everything query failed without returning a specific error.")
	}

	numResults := Et.GetNumResults()
	fmt.Printf("Found %d results:\n", numResults)

	for i := 0; i < numResults; i++ {
		name := filepath.Base(Et.GetResultFullPathName(i))
		size := Et.GetResultSize(i)
		tmod := Et.GetResultDateModified(i)
		fmt.Printf("File: %s, Size: %d, Modified: %s\n", name, size, tmod.Format("2006-01-02 15:04"))
	}
}
```

By the way, you can use `go env` to check the `GOOS` and `GOARCH`.