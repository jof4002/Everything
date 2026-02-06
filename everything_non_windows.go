//go:build !windows

package everything

import (
	"fmt"
	"time"
)

// Error message for unsupported operations on non-Windows platforms.
var errUnsupported = fmt.Errorf("Everything SDK is not supported on this platform")

// FileInfo resembles os.FileInfo but for non-Windows stubs.
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

// Walk is a stub for non-Windows platforms.
func Walk(root string, walkFn WalkFunc) error {
	return errUnsupported
}

// GetVersionString is a stub for non-Windows platforms.
func GetVersionString() (ver string) {
	return "0.0.0 (unsupported)"
}

// GetAndClearLastError is a stub for non-Windows platforms.
func GetAndClearLastError() error {
	return errUnsupported
}

// SetSearch is a stub for non-Windows platforms.
func SetSearch(str string) error {
	return errUnsupported
}

// SetMatchPath is a stub for non-Windows platforms.
func SetMatchPath(bEnable bool) {
	// No-op
}

// SetMatchCase is a stub for non-Windows platforms.
func SetMatchCase(bEnable bool) {
	// No-op
}

// SetMatchWholeWord is a stub for non-Windows platforms.
func SetMatchWholeWord(bEnable bool) {
	// No-op
}

// SetRegex is a stub for non-Windows platforms.
func SetRegex(bEnable bool) {
	// No-op
}

// SetMax is a stub for non-Windows platforms.
func SetMax(max int) {
	// No-op
}

// SetOffset is a stub for non-Windows platforms.
func SetOffset(offset int) {
	// No-op
}

// SetReplyWindow is a stub for non-Windows platforms.
func SetReplyWindow(hWnd uintptr) {
	// No-op
}

// SetReplyID is a stub for non-Windows platforms.
func SetReplyID(id int) {
	// No-op
}

// SetSort is a stub for non-Windows platforms.
func SetSort(sortMode int) {
	// No-op
}

// SetRequestFlags is a stub for non-Windows platforms.
func SetRequestFlags(flags int) {
	// No-op
}

// GetMatchPath is a stub for non-Windows platforms.
func GetMatchPath() bool {
	return false
}

// GetMatchCase is a stub for non-Windows platforms.
func GetMatchCase() bool {
	return false
}

// GetMatchWholeWord is a stub for non-Windows platforms.
func GetMatchWholeWord() bool {
	return false
}

// GetRegex is a stub for non-Windows platforms.
func GetRegex() bool {
	return false
}

// GetMax is a stub for non-Windows platforms.
func GetMax() int {
	return 0
}

// GetOffset is a stub for non-Windows platforms.
func GetOffset() int {
	return 0
}

// GetSearch is a stub for non-Windows platforms.
func GetSearch() (query string) {
	return ""
}

// GetLastError is a stub for non-Windows platforms.
func GetLastError() int {
	return -1 // Indicates an error
}

// GetReplyWindow is a stub for non-Windows platforms.
func GetReplyWindow() uintptr {
	return 0
}

// GetReplyID is a stub for non-Windows platforms.
func GetReplyID() int {
	return 0
}

// GetSort is a stub for non-Windows platforms.
func GetSort() (ret int) {
	return 0
}

// GetRequestFlags is a stub for non-Windows platforms.
func GetRequestFlags() int {
	return 0
}

// Query is a stub for non-Windows platforms.
func Query(bWait bool) (bool, error) {
	return false, errUnsupported
}

// IsQueryReply is a stub for non-Windows platforms.
func IsQueryReply(message uint32, wParam, lParam uintptr, dwId int) bool {
	return false
}

// SortResultsByPath is a stub for non-Windows platforms.
func SortResultsByPath() {
	// No-op
}

// GetNumFileResults is a stub for non-Windows platforms.
func GetNumFileResults() int {
	return 0
}

// GetNumFolderResults is a stub for non-Windows platforms.
func GetNumFolderResults() int {
	return 0
}

// GetNumResults is a stub for non-Windows platforms.
func GetNumResults() (ret int) {
	return 0
}

// GetTotFileResults is a stub for non-Windows platforms.
func GetTotFileResults() int {
	return 0
}

// GetTotFolderResults is a stub for non-Windows platforms.
func GetTotFolderResults() int {
	return 0
}

// GetTotResults is a stub for non-Windows platforms.
func GetTotResults() int {
	return 0
}

// IsVolumeResult is a stub for non-Windows platforms.
func IsVolumeResult(index int) bool {
	return false
}

// IsFolderResult is a stub for non-Windows platforms.
func IsFolderResult(index int) (ret bool) {
	return false
}

// IsFileResult is a stub for non-Windows platforms.
func IsFileResult(index int) (ret bool) {
	return false
}

// GetResultFileName is a stub for non-Windows platforms.
func GetResultFileName(index int) (name string) {
	return ""
}

// GetResultPath is a stub for non-Windows platforms.
func GetResultPath(index int) (path string) {
	return ""
}

// GetResultFullPathName is a stub for non-Windows platforms.
func GetResultFullPathName(index int) (path string) {
	return ""
}

// GetResultListSort is a stub for non-Windows platforms.
func GetResultListSort() (mode int) {
	return 0
}

// GetResultListRequestFlags is a stub for non-Windows platforms.
func GetResultListRequestFlags() int {
	return 0
}

// GetResultExtension is a stub for non-Windows platforms.
func GetResultExtension(index int) (ext string) {
	return ""
}

// GetResultSize is a stub for non-Windows platforms.
func GetResultSize(index int) (size int64) {
	return 0
}

// GetResultDateCreated is a stub for non-Windows platforms.
func GetResultDateCreated(index int) (t time.Time) {
	return time.Time{}
}

// GetResultDateModified is a stub for non-Windows platforms.
func GetResultDateModified(index int) (t time.Time) {
	return time.Time{}
}

// GetResultDateAccessed is a stub for non-Windows platforms.
func GetResultDateAccessed(index int) (t time.Time) {
	return time.Time{}
}

// GetResultAttributes is a stub for non-Windows platforms.
func GetResultAttributes(index int) int {
	return 0
}

// GetResultFileListFileName is a stub for non-Windows platforms.
func GetResultFileListFileName(index int) (name string) {
	return ""
}

// GetResultRunCount is a stub for non-Windows platforms.
func GetResultRunCount(index int) int {
	return 0
}

// GetResultDateRun is a stub for non-Windows platforms.
func GetResultDateRun(index int) (t time.Time) {
	return time.Time{}
}

// GetResultDateRecentlyChanged is a stub for non-Windows platforms.
func GetResultDateRecentlyChanged(index int) (t time.Time) {
	return time.Time{}
}

// GetResultHighlightedFileName is a stub for non-Windows platforms.
func GetResultHighlightedFileName(index int) (name string) {
	return ""
}

// GetResultHighlightedPath is a stub for non-Windows platforms.
func GetResultHighlightedPath(index int) (path string) {
	return ""
}

// GetResultHighlightedFullPathAndFileName is a stub for non-Windows platforms.
func GetResultHighlightedFullPathAndFileName(index int) (path string) {
	return ""
}

// Reset is a stub for non-Windows platforms.
func Reset() {
	// No-op
}

// CleanUp is a stub for non-Windows platforms.
func CleanUp() {
	// No-op
}

// GetMajorVersion is a stub for non-Windows platforms.
func GetMajorVersion() int {
	return 0
}

// GetMinorVersion is a stub for non-Windows platforms.
func GetMinorVersion() int {
	return 0
}

// GetRevision is a stub for non-Windows platforms.
func GetRevision() int {
	return 0
}

// GetBuildNumber is a stub for non-Windows platforms.
func GetBuildNumber() int {
	return 0
}

// Exit is a stub for non-Windows platforms.
func Exit() bool {
	return false
}

// IsDBLoaded is a stub for non-Windows platforms.
func IsDBLoaded() bool {
	return false
}

// IsAdmin is a stub for non-Windows platforms.
func IsAdmin() bool {
	return false
}

// IsAppData is a stub for non-Windows platforms.
func IsAppData() bool {
	return false
}

// RebuildDB is a stub for non-Windows platforms.
func RebuildDB() bool {
	return false
}

// UpdateAllFolderIndexes is a stub for non-Windows platforms.
func UpdateAllFolderIndexes() bool {
	return false
}

// SaveDB is a stub for non-Windows platforms.
func SaveDB() bool {
	return false
}

// SaveRunHistory is a stub for non-Windows platforms.
func SaveRunHistory() bool {
	return false
}

// DeleteRunHistory is a stub for non-Windows platforms.
func DeleteRunHistory() bool {
	return false
}

// GetTargetMachine is a stub for non-Windows platforms.
func GetTargetMachine() int {
	return 0
}

// IsFastSort is a stub for non-Windows platforms.
func IsFastSort(sortType int) bool {
	return false
}

// IsFileInfoIndexed is a stub for non-Windows platforms.
func IsFileInfoIndexed(fileInfoType int) bool {
	return false
}

// GetRunCountFromFileName is a stub for non-Windows platforms.
func GetRunCountFromFileName(filename string) int {
	return 0
}

// SetRunCountFromFileName is a stub for non-Windows platforms.
func SetRunCountFromFileName(filename string, runCount int) bool {
	return false
}

// IncRunCountFromFileName is a stub for non-Windows platforms.
func IncRunCountFromFileName(filename string) int {
	return 0
}