package ostrash

import "C"

func TrashItem(filePath string) error {
	cPath := C.Cstring(filePath)

	return nil
}
