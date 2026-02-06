//go:build windows

// Search Everything : http://voidtools.com/
// https://www.voidtools.com/support/everything/sdk/
package everything

import (
	"fmt"
	"time"
	"unsafe"

	"golang.org/x/sys/windows"
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

var everythingErrorMap = map[int]string{
	EVERYTHING_OK:                     "no error detected",
	EVERYTHING_ERROR_MEMORY:           "out of memory",
	EVERYTHING_ERROR_IPC:              "Everything search client is not running",
	EVERYTHING_ERROR_REGISTERCLASSEX:  "unable to register window class",
	EVERYTHING_ERROR_CREATEWINDOW:     "unable to create listening window",
	EVERYTHING_ERROR_CREATETHREAD:     "unable to create listening thread",
	EVERYTHING_ERROR_INVALIDINDEX:     "invalid index",
	EVERYTHING_ERROR_INVALIDCALL:      "invalid call",
	EVERYTHING_ERROR_INVALIDREQUEST:   "invalid request data, request data first",
	EVERYTHING_ERROR_INVALIDPARAMETER: "bad parameter",
}

// GetAndClearLastError retrieves the last error code from Everything and returns it as a Go error.
// The error is cleared after retrieval.
func GetAndClearLastError() error {
	errCode := GetLastError()
	if errCode == EVERYTHING_OK {
		return nil
	}
	if errMsg, ok := everythingErrorMap[errCode]; ok {
		return fmt.Errorf("Everything error %d: %s", errCode, errMsg)
	}
	return fmt.Errorf("Everything unknown error %d", errCode)
}

// manipulate search state
var Everything_SetSearch *windows.LazyProc
var Everything_SetMatchPath *windows.LazyProc
var Everything_SetMatchCase *windows.LazyProc
var Everything_SetMatchWholeWord *windows.LazyProc
var Everything_SetRegex *windows.LazyProc
var Everything_SetMax *windows.LazyProc
var Everything_SetOffset *windows.LazyProc
var Everything_SetReplyWindow *windows.LazyProc
var Everything_SetReplyID *windows.LazyProc
var Everything_SetSort *windows.LazyProc
var Everything_SetRequestFlags *windows.LazyProc

// read search state
var Everything_GetSearch *windows.LazyProc
var Everything_GetMatchPath *windows.LazyProc
var Everything_GetMatchCase *windows.LazyProc
var Everything_GetMatchWholeWord *windows.LazyProc
var Everything_GetRegex *windows.LazyProc
var Everything_GetMax *windows.LazyProc
var Everything_GetOffset *windows.LazyProc
var Everything_GetReplyWindow *windows.LazyProc
var Everything_GetReplyID *windows.LazyProc
var Everything_GetLastError *windows.LazyProc
var Everything_GetSort *windows.LazyProc
var Everything_GetRequestFlags *windows.LazyProc

// execute query
var Everything_Query *windows.LazyProc

// query reply
var Everything_IsQueryReply *windows.LazyProc

// write result state
var Everything_SortResultsByPath *windows.LazyProc

// read result state
var Everything_GetNumFileResults *windows.LazyProc
var Everything_GetNumFolderResults *windows.LazyProc
var Everything_GetNumResults *windows.LazyProc
var Everything_GetTotFileResults *windows.LazyProc
var Everything_GetTotFolderResults *windows.LazyProc
var Everything_GetTotResults *windows.LazyProc
var Everything_IsVolumeResult *windows.LazyProc
var Everything_IsFolderResult *windows.LazyProc
var Everything_IsFileResult *windows.LazyProc
var Everything_GetResultFileName *windows.LazyProc // For Everything_GetResultFileNameW
var Everything_GetResultPath *windows.LazyProc     // For Everything_GetResultPathW
var Everything_GetResultFullPathName *windows.LazyProc
var Everything_GetResultListSort *windows.LazyProc
var Everything_GetResultListRequestFlags *windows.LazyProc
var Everything_GetResultExtension *windows.LazyProc // For Everything_GetResultExtensionW
var Everything_GetResultSize *windows.LazyProc
var Everything_GetResultDateCreated *windows.LazyProc
var Everything_GetResultDateModified *windows.LazyProc
var Everything_GetResultDateAccessed *windows.LazyProc
var Everything_GetResultAttributes *windows.LazyProc
var Everything_GetResultFileListFileName *windows.LazyProc // For Everything_GetResultFileListFileNameW
var Everything_GetResultRunCount *windows.LazyProc
var Everything_GetResultDateRun *windows.LazyProc
var Everything_GetResultDateRecentlyChanged *windows.LazyProc
var Everything_GetResultHighlightedFileName *windows.LazyProc // For Everything_GetResultHighlightedFileNameW
var Everything_GetResultHighlightedPath *windows.LazyProc     // For Everything_GetResultHighlightedPathW
var Everything_GetResultHighlightedFullPathAndFileName *windows.LazyProc

// reset state and free any allocated memory
var Everything_Reset *windows.LazyProc
var Everything_CleanUp *windows.LazyProc

var Everything_GetMajorVersion *windows.LazyProc
var Everything_GetMinorVersion *windows.LazyProc
var Everything_GetRevision *windows.LazyProc
var Everything_GetBuildNumber *windows.LazyProc
var Everything_Exit *windows.LazyProc
var Everything_MSIExitAndStopService *windows.LazyProc
var Everything_MSIStartService *windows.LazyProc
var Everything_IsDBLoaded *windows.LazyProc
var Everything_IsAdmin *windows.LazyProc
var Everything_IsAppData *windows.LazyProc
var Everything_RebuildDB *windows.LazyProc
var Everything_UpdateAllFolderIndexes *windows.LazyProc
var Everything_SaveDB *windows.LazyProc
var Everything_SaveRunHistory *windows.LazyProc
var Everything_DeleteRunHistory *windows.LazyProc
var Everything_GetTargetMachine *windows.LazyProc
var Everything_IsFastSort *windows.LazyProc
var Everything_IsFileInfoIndexed *windows.LazyProc

var Everything_GetRunCountFromFileName *windows.LazyProc // For Everything_GetRunCountFromFileNameW
var Everything_SetRunCountFromFileName *windows.LazyProc // For Everything_SetRunCountFromFileNameW
var Everything_IncRunCountFromFileName *windows.LazyProc // For Everything_IncRunCountFromFileNameW

func init() {
	mod := windows.NewLazyDLL(everythingDLLName)
	if mod != nil {
		// write search state
		Everything_SetSearch = mod.NewProc("Everything_SetSearchW")
		Everything_SetMatchPath = mod.NewProc("Everything_SetMatchPath")
		Everything_SetMatchCase = mod.NewProc("Everything_SetMatchCase")
		Everything_SetMatchWholeWord = mod.NewProc("Everything_SetMatchWholeWord")
		Everything_SetRegex = mod.NewProc("Everything_SetRegex")
		Everything_SetMax = mod.NewProc("Everything_SetMax")
		Everything_SetOffset = mod.NewProc("Everything_SetOffset")
		Everything_SetReplyWindow = mod.NewProc("Everything_SetReplyWindow")
		Everything_SetReplyID = mod.NewProc("Everything_SetReplyID")
		Everything_SetSort = mod.NewProc("Everything_SetSort")
		Everything_SetRequestFlags = mod.NewProc("Everything_SetRequestFlags")

		// read search state
		Everything_GetSearch = mod.NewProc("Everything_GetSearchW")
		Everything_GetMatchPath = mod.NewProc("Everything_GetMatchPath")
		Everything_GetMatchCase = mod.NewProc("Everything_GetMatchCase")
		Everything_GetMatchWholeWord = mod.NewProc("Everything_GetMatchWholeWord")
		Everything_GetRegex = mod.NewProc("Everything_GetRegex")
		Everything_GetMax = mod.NewProc("Everything_GetMax")
		Everything_GetOffset = mod.NewProc("Everything_GetOffset")
		Everything_GetReplyWindow = mod.NewProc("Everything_GetReplyWindow")
		Everything_GetReplyID = mod.NewProc("Everything_GetReplyID")
		Everything_GetLastError = mod.NewProc("Everything_GetLastError")
		Everything_GetSort = mod.NewProc("Everything_GetSort")
		Everything_GetRequestFlags = mod.NewProc("Everything_GetRequestFlags")

		// execute query
		Everything_Query = mod.NewProc("Everything_QueryW")

		// query reply
		Everything_IsQueryReply = mod.NewProc("Everything_IsQueryReply") // FIX: Corrected mapping

		// write result state
		Everything_SortResultsByPath = mod.NewProc("Everything_SortResultsByPath")

		// read result state
		Everything_GetNumFileResults = mod.NewProc("Everything_GetNumFileResults")
		Everything_GetNumFolderResults = mod.NewProc("Everything_GetNumFolderResults")
		Everything_GetNumResults = mod.NewProc("Everything_GetNumResults")
		Everything_GetTotFileResults = mod.NewProc("Everything_GetTotFileResults")
		Everything_GetTotFolderResults = mod.NewProc("Everything_GetTotFolderResults")
		Everything_GetTotResults = mod.NewProc("Everything_GetTotResults")
		Everything_IsVolumeResult = mod.NewProc("Everything_IsVolumeResult")
		Everything_IsFolderResult = mod.NewProc("Everything_IsFolderResult")
		Everything_IsFileResult = mod.NewProc("Everything_IsFileResult")
		Everything_GetResultFileName = mod.NewProc("Everything_GetResultFileNameW")
		Everything_GetResultPath = mod.NewProc("Everything_GetResultPathW")
		Everything_GetResultFullPathName = mod.NewProc("Everything_GetResultFullPathNameW")
		Everything_GetResultListSort = mod.NewProc("Everything_GetResultListSort")
		Everything_GetResultListRequestFlags = mod.NewProc("Everything_GetResultListRequestFlags")
		Everything_GetResultExtension = mod.NewProc("Everything_GetResultExtensionW")
		Everything_GetResultSize = mod.NewProc("Everything_GetResultSize")
		Everything_GetResultDateCreated = mod.NewProc("Everything_GetResultDateCreated")
		Everything_GetResultDateModified = mod.NewProc("Everything_GetResultDateModified")
		Everything_GetResultDateAccessed = mod.NewProc("Everything_GetResultDateAccessed")
		Everything_GetResultAttributes = mod.NewProc("Everything_GetResultAttributes") // FIX: Corrected mapping
		Everything_GetResultFileListFileName = mod.NewProc("Everything_GetResultFileListFileNameW")
		Everything_GetResultRunCount = mod.NewProc("Everything_GetResultRunCount")
		Everything_GetResultDateRun = mod.NewProc("Everything_GetResultDateRun")
		Everything_GetResultDateRecentlyChanged = mod.NewProc("Everything_GetResultDateRecentlyChanged")
		Everything_GetResultHighlightedFileName = mod.NewProc("Everything_GetResultHighlightedFileNameW")
		Everything_GetResultHighlightedPath = mod.NewProc("Everything_GetResultHighlightedPathW")
		Everything_GetResultHighlightedFullPathAndFileName = mod.NewProc("Everything_GetResultHighlightedFullPathAndFileNameW")

		// reset state and free any allocated memory
		Everything_Reset = mod.NewProc("Everything_Reset")
		Everything_CleanUp = mod.NewProc("Everything_CleanUp")

		Everything_GetMajorVersion = mod.NewProc("Everything_GetMajorVersion")
		Everything_GetMinorVersion = mod.NewProc("Everything_GetMinorVersion")
		Everything_GetRevision = mod.NewProc("Everything_GetRevision")
		Everything_GetBuildNumber = mod.NewProc("Everything_GetBuildNumber")
		Everything_Exit = mod.NewProc("Everything_Exit")
		Everything_MSIExitAndStopService = mod.NewProc("Everything_MSIExitAndStopService")
		Everything_MSIStartService = mod.NewProc("Everything_MSIStartService")
		Everything_IsDBLoaded = mod.NewProc("Everything_IsDBLoaded")
		Everything_IsAdmin = mod.NewProc("Everything_IsAdmin")
		Everything_IsAppData = mod.NewProc("Everything_IsAppData")
		Everything_RebuildDB = mod.NewProc("Everything_RebuildDB")
		Everything_UpdateAllFolderIndexes = mod.NewProc("Everything_UpdateAllFolderIndexes")
		Everything_SaveDB = mod.NewProc("Everything_SaveDB")
		Everything_SaveRunHistory = mod.NewProc("Everything_SaveRunHistory")
		Everything_DeleteRunHistory = mod.NewProc("Everything_DeleteRunHistory")
		Everything_GetTargetMachine = mod.NewProc("Everything_GetTargetMachine")
		Everything_IsFastSort = mod.NewProc("Everything_IsFastSort")
		Everything_IsFileInfoIndexed = mod.NewProc("Everything_IsFileInfoIndexed")

		Everything_GetRunCountFromFileName = mod.NewProc("Everything_GetRunCountFromFileNameW")
		Everything_SetRunCountFromFileName = mod.NewProc("Everything_SetRunCountFromFileNameW")
		Everything_IncRunCountFromFileName = mod.NewProc("Everything_IncRunCountFromFileNameW")
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
	if err := SetSearch(root); err != nil {
		return err
	}
	SetRequestFlags(EVERYTHING_REQUEST_FILE_NAME | EVERYTHING_REQUEST_PATH | EVERYTHING_REQUEST_SIZE | EVERYTHING_REQUEST_DATE_MODIFIED)
	ret, err := Query(true)
	if err != nil {
		return err
	}
	if !ret {
		return fmt.Errorf("Everything Query failed for root: %s", root)
	}

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
	if Everything_GetMajorVersion != nil {
		p1, _, _ := Everything_GetMajorVersion.Call()
		p2, _, _ := Everything_GetMinorVersion.Call()
		p3, _, _ := Everything_GetRevision.Call()
		ver = fmt.Sprintf("%d.%d.%d", int(p1), int(p2), int(p3))
	}
	return
}

// SetSearch void Everything_SetSearchW(LPCWSTR lpString);
func SetSearch(str string) error {
	if Everything_SetSearch != nil {
		Everything_SetSearch.Call(uintptr(unsafe.Pointer(windows.StringToUTF16Ptr(str))))
		return GetAndClearLastError()
	}
	return fmt.Errorf("Everything_SetSearch not loaded")
}

// SetMatchPath void Everything_SetMatchPath(BOOL bEnable);
func SetMatchPath(bEnable bool) {
	if Everything_SetMatchPath != nil {
		var param int
		if bEnable {
			param = 1
		}
		Everything_SetMatchPath.Call(uintptr(param))
	}
}

// utf16FromDLL converts a UTF-16 null-terminated string pointer returned
// from a Windows DLL into a Go string.
//
// The pointer is owned by the DLL (not Go memory) and is assumed to remain
// valid only until the next DLL call. The string is copied immediately.
func stringFromDLL(ptr uintptr) string {
	if ptr == 0 {
		return ""
	}
	return windows.UTF16PtrToString((*uint16)(unsafe.Pointer(ptr)))
}

// SetMatchCase void Everything_SetMatchCase(BOOL bEnable);
func SetMatchCase(bEnable bool) {
	if Everything_SetMatchCase != nil {
		var param int
		if bEnable {
			param = 1
		}
		Everything_SetMatchCase.Call(uintptr(param))
	}
}

// SetMatchWholeWord void Everything_SetMatchWholeWord(BOOL bEnable);
func SetMatchWholeWord(bEnable bool) {
	if Everything_SetMatchWholeWord != nil {
		var param int
		if bEnable {
			param = 1
		}
		Everything_SetMatchWholeWord.Call(uintptr(param))
	}
}

// SetRegex void Everything_SetRegex(BOOL bEnable);
func SetRegex(bEnable bool) {
	if Everything_SetRegex != nil {
		var param int
		if bEnable {
			param = 1
		}
		Everything_SetRegex.Call(uintptr(param))
	}
}

// SetMax void Everything_SetMax(DWORD dwMax);
func SetMax(max int) {
	if Everything_SetMax != nil {
		Everything_SetMax.Call(uintptr(max))
	}
}

// SetOffset void Everything_SetOffset(DWORD dwOffset);
func SetOffset(offset int) {
	if Everything_SetOffset != nil {
		Everything_SetOffset.Call(uintptr(offset))
	}
}

// SetReplyWindow void Everything_SetReplyWindow(HWND hWnd);
func SetReplyWindow(hWnd windows.HWND) {
	if Everything_SetReplyWindow != nil {
		Everything_SetReplyWindow.Call(uintptr(hWnd))
	}
}

// SetReplyID void Everything_SetReplyID(DWORD dwId);
func SetReplyID(id int) {
	if Everything_SetReplyID != nil {
		Everything_SetReplyID.Call(uintptr(id))
	}
}

// SetSort void Everything_SetSort(DWORD dwSort); // Everything 1.4.1
func SetSort(sortMode int) {
	if Everything_SetSort != nil {
		Everything_SetSort.Call(uintptr(sortMode))
	}
}

// SetRequestFlags void Everything_SetRequestFlags(DWORD dwRequestFlags); // Everything 1.4.1
func SetRequestFlags(flags int) {
	if Everything_SetRequestFlags != nil {
		Everything_SetRequestFlags.Call(uintptr(flags))
	}
}

// GetMatchPath BOOL Everything_GetMatchPath(void);
func GetMatchPath() bool {
	if Everything_GetMatchPath != nil {
		r, _, _ := Everything_GetMatchPath.Call()
		return r != 0
	}
	return false
}

// GetMatchCase BOOL Everything_GetMatchCase(void);
func GetMatchCase() bool {
	if Everything_GetMatchCase != nil {
		r, _, _ := Everything_GetMatchCase.Call()
		return r != 0
	}
	return false
}

// GetMatchWholeWord BOOL Everything_GetMatchWholeWord(void);
func GetMatchWholeWord() bool {
	if Everything_GetMatchWholeWord != nil {
		r, _, _ := Everything_GetMatchWholeWord.Call()
		return r != 0
	}
	return false
}

// GetRegex BOOL Everything_GetRegex(void);
func GetRegex() bool {
	if Everything_GetRegex != nil {
		r, _, _ := Everything_GetRegex.Call()
		return r != 0
	}
	return false
}

// GetMax DWORD Everything_GetMax(void);
func GetMax() int {
	if Everything_GetMax != nil {
		r, _, _ := Everything_GetMax.Call()
		return int(r)
	}
	return 0
}

// GetOffset DWORD Everything_GetOffset(void);
func GetOffset() int {
	if Everything_GetOffset != nil {
		r, _, _ := Everything_GetOffset.Call()
		return int(r)
	}
	return 0
}

// GetSearch LPCWSTR Everything_GetSearchW(void);
func GetSearch() (query string) {
	if Everything_GetSearch != nil {
		// The Everything_GetSearchW function returns a pointer to the internal search buffer.
		// We need to copy this into a Go string.
		r, _, _ := Everything_GetSearch.Call()
		if r != 0 {
			query = stringFromDLL(r)
		}
	}
	return
}

// GetLastError DWORD Everything_GetLastError(void);
func GetLastError() int {
	if Everything_GetLastError != nil {
		r, _, _ := Everything_GetLastError.Call()
		return int(r)
	}
	return EVERYTHING_ERROR_INVALIDCALL // Return a reasonable error if the proc is not loaded
}

// GetReplyWindow HWND Everything_GetReplyWindow(void);
func GetReplyWindow() windows.HWND {
	if Everything_GetReplyWindow != nil {
		r, _, _ := Everything_GetReplyWindow.Call()
		return windows.HWND(r)
	}
	return 0
}

// GetReplyID DWORD Everything_GetReplyID(void);
func GetReplyID() int {
	if Everything_GetReplyID != nil {
		r, _, _ := Everything_GetReplyID.Call()
		return int(r)
	}
	return 0
}

// GetSort DWORD Everything_GetSort(void); // Everything 1.4.1
func GetSort() (ret int) {
	if Everything_GetSort != nil {
		r, _, _ := Everything_GetSort.Call()
		ret = int(r)
	}
	return
}

// GetRequestFlags DWORD Everything_GetRequestFlags(void); // Everything 1.4.1
func GetRequestFlags() int {
	if Everything_GetRequestFlags != nil {
		r, _, _ := Everything_GetRequestFlags.Call()
		return int(r)
	}
	return 0
}

// Query BOOL Everything_QueryW(BOOL bWait);
func Query(bWait bool) (bool, error) {
	if Everything_Query != nil {
		var param int
		if bWait {
			param = 1
		}
		r, _, _ := Everything_Query.Call(uintptr(param))
		err := GetAndClearLastError()
		return r != 0, err
	}
	return false, fmt.Errorf("Everything_Query not loaded")
}

// IsQueryReply BOOL Everything_IsQueryReply(UINT message,WPARAM wParam,LPARAM lParam,DWORD dwId);
func IsQueryReply(message uint32, wParam, lParam uintptr, dwId int) bool {
	if Everything_IsQueryReply != nil {
		r, _, _ := Everything_IsQueryReply.Call(uintptr(message), wParam, lParam, uintptr(dwId))
		return r != 0
	}
	return false
}

// SortResultsByPath void Everything_SortResultsByPath(void);
func SortResultsByPath() {
	if Everything_SortResultsByPath != nil {
		Everything_SortResultsByPath.Call()
	}
}

// GetNumFileResults DWORD Everything_GetNumFileResults(void);
func GetNumFileResults() int {
	if Everything_GetNumFileResults != nil {
		r, _, _ := Everything_GetNumFileResults.Call()
		return int(r)
	}
	return 0
}

// GetNumFolderResults DWORD Everything_GetNumFolderResults(void);
func GetNumFolderResults() int {
	if Everything_GetNumFolderResults != nil {
		r, _, _ := Everything_GetNumFolderResults.Call()
		return int(r)
	}
	return 0
}

// GetNumResults DWORD Everything_GetNumResults(void);
func GetNumResults() (ret int) {
	if Everything_GetNumResults != nil {
		r, _, _ := Everything_GetNumResults.Call()
		ret = int(r)
	}
	return
}

// GetTotFileResults DWORD Everything_GetTotFileResults(void);
func GetTotFileResults() int {
	if Everything_GetTotFileResults != nil {
		r, _, _ := Everything_GetTotFileResults.Call()
		return int(r)
	}
	return 0
}

// GetTotFolderResults DWORD Everything_GetTotFolderResults(void);
func GetTotFolderResults() int {
	if Everything_GetTotFolderResults != nil {
		r, _, _ := Everything_GetTotFolderResults.Call()
		return int(r)
	}
	return 0
}

// GetTotResults DWORD Everything_GetTotResults(void);
func GetTotResults() int {
	if Everything_GetTotResults != nil {
		r, _, _ := Everything_GetTotResults.Call()
		return int(r)
	}
	return 0
}

// IsVolumeResult BOOL Everything_IsVolumeResult(DWORD dwIndex);
func IsVolumeResult(index int) bool {
	if Everything_IsVolumeResult != nil {
		r, _, _ := Everything_IsVolumeResult.Call(uintptr(index))
		return r != 0
	}
	return false
}

// IsFolderResult BOOL Everything_IsFolderResult(DWORD dwIndex);
func IsFolderResult(index int) (ret bool) {
	if Everything_IsFolderResult != nil {
		r, _, _ := Everything_IsFolderResult.Call(uintptr(index))
		ret = r != 0
	}
	return
}

// IsFileResult BOOL Everything_IsFileResult(DWORD dwIndex);
func IsFileResult(index int) (ret bool) {
	if Everything_IsFileResult != nil {
		r, _, _ := Everything_IsFileResult.Call(uintptr(index))
		ret = r != 0
	}
	return
}

// GetResultFileName LPCWSTR Everything_GetResultFileNameW(DWORD dwIndex);
func GetResultFileName(index int) (name string) {
	if Everything_GetResultFileName != nil {
		r, _, _ := Everything_GetResultFileName.Call(uintptr(index))
		if r != 0 {
			name = stringFromDLL(r)
		}
	}
	return
}

// GetResultPath LPCWSTR Everything_GetResultPathW(DWORD dwIndex);
func GetResultPath(index int) (path string) {
	if Everything_GetResultPath != nil {
		r, _, _ := Everything_GetResultPath.Call(uintptr(index))
		if r != 0 {
			path = stringFromDLL(r)
		}
	}
	return
}

// GetResultFullPathName DWORD Everything_GetResultFullPathNameW(DWORD dwIndex,LPWSTR wbuf,DWORD wbuf_size_in_wchars);
func GetResultFullPathName(index int) (path string) {
	if Everything_GetResultFullPathName != nil {
		var pathbuf = make([]uint16, 1024)
		Everything_GetResultFullPathName.Call(uintptr(index), uintptr(unsafe.Pointer(&pathbuf[0])), 1023) // bufsize-1
		path = windows.UTF16ToString(pathbuf)
	}
	return
}

// GetResultListSort DWORD Everything_GetResultListSort(void); // Everything 1.4.1
func GetResultListSort() (mode int) {
	if Everything_GetResultListSort != nil {
		r, _, _ := Everything_GetResultListSort.Call()
		mode = int(r)
	}
	return
}

// GetResultListRequestFlags DWORD Everything_GetResultListRequestFlags(void); // Everything 1.4.1
func GetResultListRequestFlags() int {
	if Everything_GetResultListRequestFlags != nil {
		r, _, _ := Everything_GetResultListRequestFlags.Call()
		return int(r)
	}
	return 0
}

// GetResultExtension LPCWSTR Everything_GetResultExtensionW(DWORD dwIndex); // Everything 1.4.1
func GetResultExtension(index int) (ext string) {
	if Everything_GetResultExtension != nil {
		r, _, _ := Everything_GetResultExtension.Call(uintptr(index))
		if r != 0 {
			ext = stringFromDLL(r)
		}
	}
	return
}

// GetResultSize BOOL Everything_GetResultSize(DWORD dwIndex,LARGE_INTEGER *lpSize); // Everything 1.4.1
func GetResultSize(index int) (size int64) {
	if Everything_GetResultSize != nil {
		Everything_GetResultSize.Call(uintptr(index), uintptr(unsafe.Pointer(&size)))
	}
	return
}

// GetResultDateCreated BOOL Everything_GetResultDateCreated(DWORD dwIndex,FILETIME *lpDateCreated); // Everything 1.4.1
func GetResultDateCreated(index int) (t time.Time) {
	if Everything_GetResultDateCreated != nil {
		var ft windows.Filetime
		Everything_GetResultDateCreated.Call(uintptr(index), uintptr(unsafe.Pointer(&ft)))
		t = time.Unix(0, ft.Nanoseconds())
	}
	return
}

// GetResultDateModified BOOL Everything_GetResultDateModified(DWORD dwIndex,FILETIME *lpDateModified); // Everything 1.4.1
func GetResultDateModified(index int) (t time.Time) {
	if Everything_GetResultDateModified != nil {
		var ft windows.Filetime
		Everything_GetResultDateModified.Call(uintptr(index), uintptr(unsafe.Pointer(&ft)))
		t = time.Unix(0, ft.Nanoseconds())
	}
	return
}

// GetResultDateAccessed BOOL Everything_GetResultDateAccessed(DWORD dwIndex,FILETIME *lpDateAccessed); // Everything 1.4.1
func GetResultDateAccessed(index int) (t time.Time) {
	if Everything_GetResultDateAccessed != nil {
		var ft windows.Filetime
		Everything_GetResultDateAccessed.Call(uintptr(index), uintptr(unsafe.Pointer(&ft)))
		t = time.Unix(0, ft.Nanoseconds())
	}
	return
}

// GetResultAttributes DWORD Everything_GetResultAttributes(DWORD dwIndex); // Everything 1.4.1
func GetResultAttributes(index int) int {
	if Everything_GetResultAttributes != nil {
		r, _, _ := Everything_GetResultAttributes.Call(uintptr(index))
		return int(r)
	}
	return 0
}

// GetResultFileListFileName LPCWSTR Everything_GetResultFileListFileNameW(DWORD dwIndex); // Everything 1.4.1
func GetResultFileListFileName(index int) (name string) {
	if Everything_GetResultFileListFileName != nil {
		r, _, _ := Everything_GetResultFileListFileName.Call(uintptr(index))
		if r != 0 {
			name = stringFromDLL(r)
		}
	}
	return
}

// GetResultRunCount DWORD Everything_GetResultRunCount(DWORD dwIndex); // Everything 1.4.1
func GetResultRunCount(index int) int {
	if Everything_GetResultRunCount != nil {
		r, _, _ := Everything_GetResultRunCount.Call(uintptr(index))
		return int(r)
	}
	return 0
}

// GetResultDateRun BOOL Everything_GetResultDateRun(DWORD dwIndex,FILETIME *lpDateRun);
func GetResultDateRun(index int) (t time.Time) {
	if Everything_GetResultDateRun != nil {
		var ft windows.Filetime
		Everything_GetResultDateRun.Call(uintptr(index), uintptr(unsafe.Pointer(&ft)))
		t = time.Unix(0, ft.Nanoseconds())
	}
	return
}

// GetResultDateRecentlyChanged BOOL Everything_GetResultDateRecentlyChanged(DWORD dwIndex,FILETIME *lpDateRecentlyChanged);
func GetResultDateRecentlyChanged(index int) (t time.Time) {
	if Everything_GetResultDateRecentlyChanged != nil {
		var ft windows.Filetime
		Everything_GetResultDateRecentlyChanged.Call(uintptr(index), uintptr(unsafe.Pointer(&ft)))
		t = time.Unix(0, ft.Nanoseconds())
	}
	return
}

// GetResultHighlightedFileName LPCWSTR Everything_GetResultHighlightedFileNameW(DWORD dwIndex); // Everything 1.4.1
func GetResultHighlightedFileName(index int) (name string) {
	if Everything_GetResultHighlightedFileName != nil {
		r, _, _ := Everything_GetResultHighlightedFileName.Call(uintptr(index))
		if r != 0 {
			name = stringFromDLL(r)
		}
	}
	return
}

// GetResultHighlightedPath LPCWSTR Everything_GetResultHighlightedPathW(DWORD dwIndex); // Everything 1.4.1
func GetResultHighlightedPath(index int) (path string) {
	if Everything_GetResultHighlightedPath != nil {
		r, _, _ := Everything_GetResultHighlightedPath.Call(uintptr(index))
		if r != 0 {
			path = stringFromDLL(r)
		}
	}
	return
}

// GetResultHighlightedFullPathAndFileName LPCWSTR Everything_GetResultHighlightedFullPathAndFileNameW(DWORD dwIndex); // Everything 1.4.1
func GetResultHighlightedFullPathAndFileName(index int) (path string) {
	if Everything_GetResultHighlightedFullPathAndFileName != nil {
		r, _, _ := Everything_GetResultHighlightedFullPathAndFileName.Call(uintptr(index))
		if r != 0 {
			path = stringFromDLL(r)
		}
	}
	return
}

// Reset void Everything_Reset(void);
func Reset() {
	if Everything_Reset != nil {
		Everything_Reset.Call()
	}
}

// CleanUp void Everything_CleanUp(void);
func CleanUp() {
	if Everything_CleanUp != nil {
		Everything_CleanUp.Call()
	}
}

// GetMajorVersion DWORD Everything_GetMajorVersion(void);
func GetMajorVersion() int {
	if Everything_GetMajorVersion != nil {
		r, _, _ := Everything_GetMajorVersion.Call()
		return int(r)
	}
	return 0
}

// GetMinorVersion DWORD Everything_GetMinorVersion(void);
func GetMinorVersion() int {
	if Everything_GetMinorVersion != nil {
		r, _, _ := Everything_GetMinorVersion.Call()
		return int(r)
	}
	return 0
}

// GetRevision DWORD Everything_GetRevision(void);
func GetRevision() int {
	if Everything_GetRevision != nil {
		r, _, _ := Everything_GetRevision.Call()
		return int(r)
	}
	return 0
}

// GetBuildNumber DWORD Everything_GetBuildNumber(void);
func GetBuildNumber() int {
	if Everything_GetBuildNumber != nil {
		r, _, _ := Everything_GetBuildNumber.Call()
		return int(r)
	}
	return 0
}

// Exit BOOL Everything_Exit(void);
func Exit() bool {
	if Everything_Exit != nil {
		r, _, _ := Everything_Exit.Call()
		return r != 0
	}
	return false
}

// Everything_IsDBLoaded BOOL Everything_IsDBLoaded(void); // Everything 1.4.1
func IsDBLoaded() bool {
	if Everything_IsDBLoaded != nil {
		r, _, _ := Everything_IsDBLoaded.Call()
		return r != 0
	}
	return false
}

// Everything_IsAdmin BOOL Everything_IsAdmin(void); // Everything 1.4.1
func IsAdmin() bool {
	if Everything_IsAdmin != nil {
		r, _, _ := Everything_IsAdmin.Call()
		return r != 0
	}
	return false
}

// Everything_IsAppData BOOL Everything_IsAppData(void); // Everything 1.4.1
func IsAppData() bool {
	if Everything_IsAppData != nil {
		r, _, _ := Everything_IsAppData.Call()
		return r != 0
	}
	return false
}

// Everything_RebuildDB BOOL Everything_RebuildDB(void); // Everything 1.4.1
func RebuildDB() bool {
	if Everything_RebuildDB != nil {
		r, _, _ := Everything_RebuildDB.Call()
		return r != 0
	}
	return false
}

// Everything_UpdateAllFolderIndexes BOOL Everything_UpdateAllFolderIndexes(void); // Everything 1.4.1
func UpdateAllFolderIndexes() bool {
	if Everything_UpdateAllFolderIndexes != nil {
		r, _, _ := Everything_UpdateAllFolderIndexes.Call()
		return r != 0
	}
	return false
}

// Everything_SaveDB BOOL Everything_SaveDB(void); // Everything 1.4.1
func SaveDB() bool {
	if Everything_SaveDB != nil {
		r, _, _ := Everything_SaveDB.Call()
		return r != 0
	}
	return false
}

// Everything_SaveRunHistory BOOL Everything_SaveRunHistory(void); // Everything 1.4.1
func SaveRunHistory() bool {
	if Everything_SaveRunHistory != nil {
		r, _, _ := Everything_SaveRunHistory.Call()
		return r != 0
	}
	return false
}

// Everything_DeleteRunHistory BOOL Everything_DeleteRunHistory(void); // Everything 1.4.1
func DeleteRunHistory() bool {
	if Everything_DeleteRunHistory != nil {
		r, _, _ := Everything_DeleteRunHistory.Call()
		return r != 0
	}
	return false
}

// Everything_GetTargetMachine DWORD Everything_GetTargetMachine(void); // Everything 1.4.1
func GetTargetMachine() int {
	if Everything_GetTargetMachine != nil {
		r, _, _ := Everything_GetTargetMachine.Call()
		return int(r)
	}
	return 0
}

// Everything_IsFastSort BOOL Everything_IsFastSort(DWORD sortType); // Everything 1.4.1.859
func IsFastSort(sortType int) bool {
	if Everything_IsFastSort != nil {
		r, _, _ := Everything_IsFastSort.Call(uintptr(sortType))
		return r != 0
	}
	return false
}

// Everything_IsFileInfoIndexed BOOL Everything_IsFileInfoIndexed(DWORD fileInfoType); // Everything 1.4.1.859
func IsFileInfoIndexed(fileInfoType int) bool {
	if Everything_IsFileInfoIndexed != nil {
		r, _, _ := Everything_IsFileInfoIndexed.Call(uintptr(fileInfoType))
		return r != 0
	}
	return false
}

// Everything_GetRunCountFromFileNameW DWORD Everything_GetRunCountFromFileNameW(LPCWSTR lpFileName); // Everything 1.4.1
func GetRunCountFromFileName(filename string) int {
	if Everything_GetRunCountFromFileName != nil {
		r, _, _ := Everything_GetRunCountFromFileName.Call(uintptr(unsafe.Pointer(windows.StringToUTF16Ptr(filename))))
		return int(r)
	}
	return 0
}

// Everything_SetRunCountFromFileNameW BOOL Everything_SetRunCountFromFileNameW(LPCWSTR lpFileName,DWORD dwRunCount); // Everything 1.4.1
func SetRunCountFromFileName(filename string, runCount int) bool {
	if Everything_SetRunCountFromFileName != nil {
		r, _, _ := Everything_SetRunCountFromFileName.Call(uintptr(unsafe.Pointer(windows.StringToUTF16Ptr(filename))), uintptr(runCount))
		return r != 0
	}
	return false
}

// Everything_IncRunCountFromFileNameW DWORD Everything_IncRunCountFromFileNameW(LPCWSTR lpFileName); // Everything 1.4.1
func IncRunCountFromFileName(filename string) int {
	if Everything_IncRunCountFromFileName != nil {
		r, _, _ := Everything_IncRunCountFromFileName.Call(uintptr(unsafe.Pointer(windows.StringToUTF16Ptr(filename))))
		return int(r)
	}
	return 0
}
