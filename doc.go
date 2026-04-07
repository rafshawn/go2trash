// `go2trash` provides cross-platform OS trash integration for Go.
//
// It moves files to the native operating system trash/recycle bin so that
// users can restore them later, instead of permanently deleting with `os.Remove`.
//
// Platform support:
//
//   - macOS: Uses NSFileManager.trashItemAtURL (native Cocoa API via CGO).
//     Supports Finder "Put Back" functionality. No security prompts.
//
//   - Linux: Follows the FreeDesktop.org Trash Specification 1.0.
//     Files appear in GNOME Files, Dolphin, and other compliant file managers.
//
//   - Windows:
//
// Basic usage:
//
//	err := ostrash.MoveToTrash("/path/to/file.txt")
//	if err != nil {
//	    log.Fatal(err)
//	}
package go2trash
