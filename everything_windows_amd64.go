// Package Everything calls Search Everything's dll
// I'm newbie in go and I'm just calling dll
// Search Everything : http://voidtools.com/
package Everything

import (
	"fmt"
	"syscall"
	"time"
	"unsafe"
)

const (
	EVERYTHING_OK                     = 0 // no error detected
	EVERYTHING_ERROR_MEMORY           = 1 // out of memory.
	EVERYTHING_ERROR_IPC              = 2 // Everything search client is not running
	EVERYTHING_ERROR_REGISTERCLASSEX  = 3 // unable to register window class.
	EVERYTHING_ERROR_CREATEWINDOW     = 4 // unable to create listening window
	EVERYTHING_ERROR_CREATETHREAD     = 5 // unable to create listening thread
	EVERYTHING_ERROR_INVALIDINDEX     = 6 // invalid index
	EVERYTHING_ERROR_INVALIDCALL      = 7 // invalid call
	EVERYTHING_ERROR_INVALIDREQUEST   = 8 // invalid request data, request data first.
	EVERYTHING_ERROR_INVALIDPARAMETER = 9 // bad parameter.
)

const (
	EVERYTHING_SORT_NAME_ASCENDING                   = 1
	EVERYTHING_SORT_NAME_DESCENDING                  = 2
	EVERYTHING_SORT_PATH_ASCENDING                   = 3
	EVERYTHING_SORT_PATH_DESCENDING                  = 4
	EVERYTHING_SORT_SIZE_ASCENDING                   = 5
	EVERYTHING_SORT_SIZE_DESCENDING                  = 6
	EVERYTHING_SORT_EXTENSION_ASCENDING              = 7
	EVERYTHING_SORT_EXTENSION_DESCENDING             = 8
	EVERYTHING_SORT_TYPE_NAME_ASCENDING              = 9
	EVERYTHING_SORT_TYPE_NAME_DESCENDING             = 10
	EVERYTHING_SORT_DATE_CREATED_ASCENDING           = 11
	EVERYTHING_SORT_DATE_CREATED_DESCENDING          = 12
	EVERYTHING_SORT_DATE_MODIFIED_ASCENDING          = 13
	EVERYTHING_SORT_DATE_MODIFIED_DESCENDING         = 14
	EVERYTHING_SORT_ATTRIBUTES_ASCENDING             = 15
	EVERYTHING_SORT_ATTRIBUTES_DESCENDING            = 16
	EVERYTHING_SORT_FILE_LIST_FILENAME_ASCENDING     = 17
	EVERYTHING_SORT_FILE_LIST_FILENAME_DESCENDING    = 18
	EVERYTHING_SORT_RUN_COUNT_ASCENDING              = 19
	EVERYTHING_SORT_RUN_COUNT_DESCENDING             = 20
	EVERYTHING_SORT_DATE_RECENTLY_CHANGED_ASCENDING  = 21
	EVERYTHING_SORT_DATE_RECENTLY_CHANGED_DESCENDING = 22
	EVERYTHING_SORT_DATE_ACCESSED_ASCENDING          = 23
	EVERYTHING_SORT_DATE_ACCESSED_DESCENDING         = 24
	EVERYTHING_SORT_DATE_RUN_ASCENDING               = 25
	EVERYTHING_SORT_DATE_RUN_DESCENDING              = 26
)

const (
	EVERYTHING_REQUEST_FILE_NAME                           = 0x00000001
	EVERYTHING_REQUEST_PATH                                = 0x00000002
	EVERYTHING_REQUEST_FULL_PATH_AND_FILE_NAME             = 0x00000004
	EVERYTHING_REQUEST_EXTENSION                           = 0x00000008
	EVERYTHING_REQUEST_SIZE                                = 0x00000010
	EVERYTHING_REQUEST_DATE_CREATED                        = 0x00000020
	EVERYTHING_REQUEST_DATE_MODIFIED                       = 0x00000040
	EVERYTHING_REQUEST_DATE_ACCESSED                       = 0x00000080
	EVERYTHING_REQUEST_ATTRIBUTES                          = 0x00000100
	EVERYTHING_REQUEST_FILE_LIST_FILE_NAME                 = 0x00000200
	EVERYTHING_REQUEST_RUN_COUNT                           = 0x00000400
	EVERYTHING_REQUEST_DATE_RUN                            = 0x00000800
	EVERYTHING_REQUEST_DATE_RECENTLY_CHANGED               = 0x00001000
	EVERYTHING_REQUEST_HIGHLIGHTED_FILE_NAME               = 0x00002000
	EVERYTHING_REQUEST_HIGHLIGHTED_PATH                    = 0x00004000
	EVERYTHING_REQUEST_HIGHLIGHTED_FULL_PATH_AND_FILE_NAME = 0x00008000
)

// write search state
var callSetSearch *syscall.LazyProc
var callSetMatchPath *syscall.LazyProc
var callSetMatchCase *syscall.LazyProc
var callSetRegex *syscall.LazyProc
var callSetSort *syscall.LazyProc
var callSetRequestFlags *syscall.LazyProc

// read search state
var callGetSort *syscall.LazyProc

// execute query
var callQuery *syscall.LazyProc

// query reply

// write result state

// read result state
var callGetNumResults *syscall.LazyProc

var callGetResultFullPathName *syscall.LazyProc
var callIsFolderResult *syscall.LazyProc
var callIsFileResult *syscall.LazyProc

var callGetResultListSort *syscall.LazyProc
var callGetResultSize *syscall.LazyProc
var callGetResultDateModified *syscall.LazyProc

// reset state and free any allocated memory
var callReset *syscall.LazyProc

func init() {
	mod := syscall.NewLazyDLL("Everything64.dll")
	if mod != nil {
		callSetSearch = mod.NewProc("Everything_SetSearchW")
		callSetMatchPath = mod.NewProc("Everything_SetMatchPath")
		callSetMatchCase = mod.NewProc("Everything_SetMatchCase")
		callSetRegex = mod.NewProc("Everything_SetRegex")
		callSetSort = mod.NewProc("Everything_SetSort")
		callSetRequestFlags = mod.NewProc("Everything_SetRequestFlags")
		callGetSort = mod.NewProc("Everything_GetSort")
		callQuery = mod.NewProc("Everything_QueryW")
		callGetNumResults = mod.NewProc("Everything_GetNumResults")
		callGetResultFullPathName = mod.NewProc("Everything_GetResultFullPathNameW")
		callIsFolderResult = mod.NewProc("Everything_IsFolderResult")
		callIsFileResult = mod.NewProc("Everything_IsFileResult")
		callGetResultListSort = mod.NewProc("Everything_GetResultListSort")
		callGetResultSize = mod.NewProc("Everything_GetResultSize")
		callGetResultDateModified = mod.NewProc("Everything_GetResultDateModified")
		callReset = mod.NewProc("Everything_Reset")
	}
}

// FileInfo resemble os.FileInfo
type FileInfo struct {
	name    string
	size    int64
	modTime time.Time
	isDir   bool
}

func (fi *FileInfo) Name() string       { return fi.name }
func (fi *FileInfo) Size() int64        { return fi.size }
func (fi *FileInfo) ModTime() time.Time { return fi.modTime }
func (fi *FileInfo) IsDir() bool        { return fi.isDir }

// WalkFunc is the type of the function called for each file or directory visited by Walk.
type WalkFunc func(path string, info FileInfo, err error) error

// Walk calling walkFn for each file or directory in queried resulr
func Walk(root string, walkFn WalkFunc) error {
	SetSearch(root)
	SetRequestFlags(EVERYTHING_REQUEST_FILE_NAME | EVERYTHING_REQUEST_PATH | EVERYTHING_REQUEST_SIZE | EVERYTHING_REQUEST_DATE_MODIFIED)
	Query(true)
	num := GetNumResults()
	for i := 0; i < num; i++ {
		var fi FileInfo
		fi.name = GetResultFullPathName(i)
		fi.size = GetResultSize(i)
		fi.modTime = GetResultDateModified(i)
		fi.isDir = IsFolderResult(i)
		err := walkFn(fi.name, fi, nil)
		if err != nil {
			return err
		}
	}
	return nil

}

// GetVersionString print ver
func GetVersionString() (ver string) {
	mod := syscall.NewLazyDLL("Everything64.dll")
	if mod != nil {
		fmajor := mod.NewProc("Everything_GetMajorVersion")
		fminor := mod.NewProc("Everything_GetMinorVersion")
		frevision := mod.NewProc("Everything_GetRevision")
		p1, _, _ := fmajor.Call()
		p2, _, _ := fminor.Call()
		p3, _, _ := frevision.Call()
		ver = fmt.Sprintf("%d.%d.%d", int(p1), int(p2), int(p3))
	}
	return
}

// SetSearch void Everything_SetSearchW(LPCWSTR lpString);
func SetSearch(str string) {
	if callSetSearch != nil {
		callSetSearch.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(str))))
	}
}

// SetMatchPath void Everything_SetMatchPath(BOOL bEnable);
func SetMatchPath(bEnable bool) {
	if callSetMatchPath != nil {
		var param int
		if bEnable {
			param = 1
		}
		callSetMatchPath.Call(uintptr(param))
	}
}

// SetMatchCase void Everything_SetMatchCase(BOOL bEnable);
func SetMatchCase(bEnable bool) {
	if callSetMatchCase != nil {
		var param int
		if bEnable {
			param = 1
		}
		callSetMatchCase.Call(uintptr(param))
	}
}

// SetRegex void Everything_SetRegex(BOOL bEnable);
func SetRegex(bEnable bool) {
	if callSetRegex != nil {
		var param int
		if bEnable {
			param = 1
		}
		callSetRegex.Call(uintptr(param))
	}
}

// SetSort void Everything_SetSort(DWORD dwSort); // Everything 1.4.1
func SetSort(sortMode int) {
	if callSetSort != nil {
		callSetSort.Call(uintptr(sortMode))
	}
}

// SetRequestFlags void Everything_SetRequestFlags(DWORD dwRequestFlags); // Everything 1.4.1
func SetRequestFlags(flags int) {
	if callSetRequestFlags != nil {
		callSetRequestFlags.Call(uintptr(flags))
	}
}

// GetSort DWORD Everything_GetSort(void); // Everything 1.4.1
func GetSort() (ret int) {
	if callGetSort != nil {
		r, _, _ := callGetSort.Call()
		ret = int(r)
	}
	return
}

// Query BOOL Everything_QueryW(BOOL bWait);
func Query(bWait bool) (ret bool) {
	if callQuery != nil {
		var param int
		if bWait {
			param = 1
		}
		r, _, _ := callQuery.Call(uintptr(param))
		ret = r != 0
	}
	return
}

// GetNumResults DWORD Everything_GetNumResults(void);
func GetNumResults() (ret int) {
	if callGetNumResults != nil {
		r, _, _ := callGetNumResults.Call()
		ret = int(r)
	}
	return
}

// GetResultFullPathName DWORD Everything_GetResultFullPathNameW(DWORD dwIndex,LPWSTR wbuf,DWORD wbuf_size_in_wchars);
func GetResultFullPathName(index int) (path string) {
	if callGetResultFullPathName != nil {
		var pathbuf []uint16
		pathbuf = make([]uint16, 1024)
		callGetResultFullPathName.Call(uintptr(index), uintptr(unsafe.Pointer(&pathbuf[0])), 1023) // bufsize-1
		path = syscall.UTF16ToString(pathbuf)
	}
	return
}

// IsFolderResult BOOL Everything_IsFolderResult(DWORD dwIndex);
func IsFolderResult(index int) (ret bool) {
	if callIsFolderResult != nil {
		r, _, _ := callIsFolderResult.Call(uintptr(index))
		ret = r != 0
	}
	return
}

// IsFileResult BOOL Everything_IsFileResult(DWORD dwIndex);
func IsFileResult(index int) (ret bool) {
	if callIsFileResult != nil {
		r, _, _ := callIsFileResult.Call(uintptr(index))
		ret = r != 0
	}
	return
}

// GetResultListSort DWORD Everything_GetResultListSort(void); // Everything 1.4.1
func GetResultListSort() (mode int) {
	if callGetResultListSort != nil {
		r, _, _ := callGetResultListSort.Call()
		mode = int(r)
	}
	return
}

// GetResultSize BOOL Everything_GetResultSize(DWORD dwIndex,LARGE_INTEGER *lpSize); // Everything 1.4.1
func GetResultSize(index int) (size int64) {
	if callGetResultSize != nil {
		callGetResultSize.Call(uintptr(index), uintptr(unsafe.Pointer(&size)))
	}
	return
}

// GetResultDateModified BOOL Everything_GetResultDateModified(DWORD dwIndex,FILETIME *lpDateModified); // Everything 1.4.1
func GetResultDateModified(index int) (t time.Time) {
	if callGetResultDateModified != nil {
		var ft syscall.Filetime
		callGetResultDateModified.Call(uintptr(index), uintptr(unsafe.Pointer(&ft)))
		t = time.Unix(0, ft.Nanoseconds())
	}
	return
}

// Reset void Everything_Reset(void);
func Reset() {
	if callReset != nil {
		callReset.Call()
	}
}
