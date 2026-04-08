//go:build darwin

package go2trash

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation

#import <Foundation/Foundation.h>
#include <stdlib.h>

// moveToTrash moves a file to the macOS Trash using NSFileManager.
// Returns 0 on success, -1 on failure.
static int moveToTrash(const char* path) {
	@autoreleasepool {
		NSFileManager *fm = [NSFileManager defaultManager];
		NSString *nsPath = [NSString stringWithUTF8String:path];
		NSURL *url = [NSURL fileURLWithPath:nsPath];
		NSError *error = nil;
		BOOL ok = [fm trashItemAtURL:url
		            resultingItemURL:nil
		                       error:&error];
		if (!ok) {
			return -1;
		}
		return 0;
	}
}
*/
import "C"

import (
	"fmt"
	"unsafe"
)

// MoveToTrash sends a file or directory to the macOS Trash.
//
// Uses NSFileManager's trashItemAtURL:resultingItemURL:error: method,
// which is the official Apple API.
func MoveToTrash(filePath string) error {
	cPath := C.CString(filePath)
	defer C.free(unsafe.Pointer(cPath))

	result := C.moveToTrash(cPath)
	if result != 0 {
		return fmt.Errorf("go2trash: failed to trash file: %s", filePath)
	}
	return nil
}
