# go2trash

[![Go Reference](https://pkg.go.dev/badge/github.com/rafshawn/go2trash.svg)](https://pkg.go.dev/github.com/rafshawn/go2trash)
[![Go Report Card](https://goreportcard.com/badge/github.com/rafshawn/go2trash)](https://goreportcard.com/report/github.com/rafshawn/go2trash)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENCE)

Cross-platform OS trash integration for Go. Moves files to the native operating system trash/recycle bin so users can restore them later, instead of permanently deleting with `os.Remove`.

## Platform Support

| Platform | Implementation |
|----------|---------------|
| Linux | [FreeDesktop.org Trash Specification 1.0](https://specifications.freedesktop.org/trash/latest/) |
| macOS | [`NSFileManager.trashItem`](https://developer.apple.com/documentation/foundation/filemanager/trashitem(at:resultingitemurl:)) (macOS 10.8+) |
| Windows | [`SHFileOperationW`](https://learn.microsoft.com/en-us/windows/win32/api/shellapi/ns-shellapi-shfileopstructw) via [`golang.org/x/sys/windows`](https://pkg.go.dev/golang.org/x/sys/windows#AdjustTokenPrivileges) |

## Installation

```bash
go get github.com/rafshawn/go2trash
```

## Usage

```go
package main

import (
    "log"

    go2trash "github.com/rafshawn/go2trash"
)

func main() {
    err := go2trash.MoveToTrash("/path/to/file.txt")
    if err != nil {
        log.Fatal(err)
    }
}
```

## Overview

### Linux
Trashed files are sent to user's (or root's) trash (`~/.local/share/Trash/`), following the FreeDesktop.org Trash Specification[^1]. A `.trashinfo` file is created with the original path and deletion timestamp.

```
~/.local/share/Trash/
    files/
        file.txt
    info/
        file.txt.trashinfo
```

The `.trashinfo` file contains metadata to allow file restoration:

```
[Trash Info]
Path=/absolute/path/to/file.txt
DeletionDate=20060102T03:04:05
```

### macOS
> [!IMPORTANT]
> macOS builds require a Mac due to CGO and the Objective-C SDK dependency.

File operations are entirely handled by [`NSFileManager`](https://developer.apple.com/documentation/foundation/filemanager/) through [CGO](https://go.dev/wiki/cgo) and Objective-C. The `trashItemAtURL` method moves files to the Trash without showing security prompts, and supports the Finder "Put Back" feature. GO just waits for a `return 0` on success or a `return -1` on failure.

This means:
- The file appears in Finder's Trash with full "Put Back" support
- No security prompts (*"App would like to control Finder"*)
- No `osascript`/`AppleScript` subprocess (*native performance*)
- Works in headless environments

### Windows

The Windows implementation uses `SHFileOperationW` from `Shell32.dll` via `golang.org/x/sys/windows` to send files to the Recycle Bin. It is a pure syscall implementation (no CGO or MinGW is dependency required).

The `FOF_ALLOWUNDO` flag enables trash functionality, while additional flags suppress confirmation dialogs and progress UI for silent operation.

## API

```go
func MoveToTrash(filePath string) error
```

`MoveToTrash` sends a file or directory to the OS trash. Returns an error if the file does not exist or the operation fails.

## Testing

Run tests locally:

```bash
go test -v ./...
```

The project includes GitHub Actions that run tests on macOS, Linux, and Windows on every push and pull request.

## License

MIT License. See [LICENCE](LICENCE) for details.

## Resources
- [Calling C from Go](https://ericchiang.github.io/post/cgo/) (Eric Chiang)