package main

import (
	"fmt"
	"path/filepath"

	Et "github.com/jof4002/Everything"
)

func main() {
	Et.SetSearch("*.go")
	Et.SetSort(Et.EVERYTHING_SORT_SIZE_DESCENDING)
	Et.SetRequestFlags(Et.EVERYTHING_REQUEST_FILE_NAME | Et.EVERYTHING_REQUEST_PATH | Et.EVERYTHING_REQUEST_SIZE | Et.EVERYTHING_REQUEST_DATE_MODIFIED)
	Et.Query(true)

	if Et.GetResultListSort() != Et.EVERYTHING_SORT_SIZE_DESCENDING {
		fmt.Println("Sorted Query failed :", Et.GetResultListSort())
	}

	num := Et.GetNumResults()
	if num > 100 {
		num = 100
	}
	for i := 0; i < num; i++ {
		path := Et.GetResultFullPathName(i)
		name := filepath.Base(path)
		size := Et.GetResultSize(i)
		tmod := Et.GetResultDateModified(i)
		fmt.Println(name, size, tmod.Format("2006-01-02 15:04"))
	}

}
