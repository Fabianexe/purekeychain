package utility

import (
	"unsafe"

	"github.com/ebitengine/purego"
)

// deRef dereferences a C pointer
func deRef(in uintptr) uintptr {
	x := (*uintptr)(unsafe.Pointer(in))

	return *x
}

// Load loads a C library
func Load(handle uintptr, name string) uintptr {
	ret, err := purego.Dlsym(handle, name)
	if err != nil {
		panic(err)
	}

	return deRef(ret)
}
