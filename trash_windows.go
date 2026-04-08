//go:build windows

package go2trash

import (
	"errors"
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

// Win32 constants for SHFileOperationW
const (
	foDelete          = 0x0003 // FO_DELETE
	fofAllowUndo      = 0x0040 // FOF_ALLOWUNDO — send to Recycle Bin instead of permanent delete
	fofNoConfirmation = 0x0010 // FOF_NOCONFIRMATION — don't show confirmation dialogs
	fofSilent         = 0x0004 // FOF_SILENT — don't show progress UI
	fofNoErrorUI      = 0x0400 // FOF_NOERRORUI — don't show error dialogs
)

var (
	shell32DLL         = windows.NewLazySystemDLL("Shell32.dll")
	shFileOperationWFn = shell32DLL.NewProc("SHFileOperationW")
)

// _SHFILEOPSTRUCTW mirrors the Win32 SHFILEOPSTRUCTW structure.
// See: https://learn.microsoft.com/en-us/windows/win32/api/shellapi/ns-shellapi-shfileopstructw
type _SHFILEOPSTRUCTW struct {
	hwnd                  uintptr
	wFunc                 uint32
	pFrom                 uintptr
	pTo                   uintptr
	fFlags                uint16
	fAnyOperationsAborted int32
	hNameMappings         uintptr
	lpszProgressTitle     uintptr
}

// MoveToTrash sends a file or directory to the Windows Recycle Bin.
//
// Uses the SHFileOperationW API from Shell32.dll with the FOF_ALLOWUNDO
// flag, which is the standard Win32 approach for Recycle Bin operations.
func MoveToTrash(filePath string) error {
	// SHFileOperationW expects double-null-terminated UTF-16 strings
	utf16Path, err := windows.UTF16FromString(filePath)
	if err != nil {
		return fmt.Errorf("go2trash: invalid path: %w", err)
	}
	// Append an extra null terminator (double-null terminated list)
	utf16Path = append(utf16Path, 0)

	param := &_SHFILEOPSTRUCTW{
		wFunc:  foDelete,
		pFrom:  uintptr(unsafe.Pointer(&utf16Path[0])),
		fFlags: fofAllowUndo | fofNoConfirmation | fofSilent | fofNoErrorUI,
	}

	ret, _, callErr := shFileOperationWFn.Call(uintptr(unsafe.Pointer(param)))
	if ret != 0 {
		if callErr != nil && !errors.Is(callErr, windows.ERROR_SUCCESS) {
			return fmt.Errorf("go2trash: SHFileOperationW failed (code %d): %w", ret, callErr)
		}
		return fmt.Errorf("go2trash: SHFileOperationW failed with code %d", ret)
	}
	return nil
}
