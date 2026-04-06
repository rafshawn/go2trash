
# go-2-trash
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