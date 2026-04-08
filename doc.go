// Package go2trash provides cross-platform OS trash integration for Go.
//
// It moves files to the native operating system trash/recycle bin so that
// users can restore them later, instead of permanently deleting with os.Remove.
//
// # Platform Support
//
// macOS: Uses NSFileManager.trashItemAtURL (native Cocoa API via CGO).
// Supports Finder "Put Back" functionality. No security prompts required.
//
// Linux: Follows the FreeDesktop.org Trash Specification 1.0.
// Files appear in GNOME Files, Dolphin, and other compliant file managers.
// Creates .trashinfo metadata files for recovery.
//
// Windows: Uses SHFileOperationW via golang.org/x/sys/windows.
// Pure syscall implementation with no CGO dependency. Sends files to the
// Recycle Bin with no confirmation dialogs shown.
//
// # Basic Usage
//
//	err := go2trash.MoveToTrash("/path/to/file.txt")
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// # Index
//
//   - [MoveToTrash] -- sends a file or directory to the OS trash
//
// # Testing
//
// Run tests with: go test -v ./...
//
// Cross-platform CI runs tests on macOS, Linux, and Windows via GitHub Actions.
package go2trash
