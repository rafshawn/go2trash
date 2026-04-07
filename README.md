
# go2trash
`go2trash` provides cross-platform OS trash integration for Go.

It moves files to the native OS trash/bin and serves as an alternative to `os.Remove`.

## Basic usage
```go
err := ostrash.TrashItem("/path/to/file.txt")
if err != nil {
    log.Fatal(err)
}
```

## Logic
1. Call function `TrashItem` with path as a string
2. On compile time, check build (OS)
3. Check if exists:
   - file and its absolute path
   - trash directory
4. Move file to trash directory
5. Write `.trashinfo` for recovery information

### Linux
Trashed files are sent to the user's (or root's) home trash (`~/.local/share/Trash/`). Deleting a file creates a `.trashinfo` file to store the original file path and date of deletion. Files are sent to a `/files` subdir, and `.trashinfo` in a `/info` subdir. This is the standard for all major Linux environments (GNOME, KDE, etc.)[^1].

```
~/.local/share/Trash/
    files/
        photo.jpg
    info/
        photo.jpg.trashinfo
```

`TrashItem()` sends a file or directory to the trash and creates a `.trashinfo` metadata file so that files can be displayed and restored.

### macOS
Anything to do with Apple's file management is handled by the `NSFileManager`[^2], including moving files to trash[^3]. So Go just lets `NSFileManager` handle the action, and waits for a return value. If `RETURN 0`, success, else if `RETURN -1`, fail.

It doesn't look like there's a way to build macOS apps without a mac, since the required SDK to understand Objective-C is only available on Macs.

### Windows
[^4]

## Resources
- [Calling C from Go](https://ericchiang.github.io/post/cgo/) (Eric Chiang)

[^1]: **Linux**: [Trash Specifications v1.0 (FreeDesktop.org)](https://specifications.freedesktop.org/trash/latest/)
[^2]: **macOS**: [Apple FileManager docs](https://developer.apple.com/documentation/foundation/filemanager/)
[^3]: NSFileManager [`trashItem()` method](https://developer.apple.com/documentation/foundation/filemanager/trashitem(at:resultingitemurl:))
[^4]: **Windows**:
- [Go Wiki: `cgo` Documentation](https://go.dev/wiki/cgo)