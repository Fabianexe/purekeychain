package cfdictionary

import (
	"github.com/ebitengine/purego"

	"github.com/Fabianexe/purekeychain/internal/utility"
)

type CFDictionary uintptr

// CreateDictionary creates a C-dict from a go map of C types
func CreateDictionary[K ~uintptr, V ~uintptr, M ~map[K]V](in M) CFDictionary {
	length := len(in)
	keys := make([]uintptr, 0, length)
	values := make([]uintptr, 0, length)
	for k, v := range in {
		keys = append(keys, uintptr(k))
		values = append(values, uintptr(v))
	}

	ptr := cfDictionaryCreate(0, &keys[0], &values[0], length, kCFTypeDictionaryKeyCallBacks, kCFTypeDictionaryValueCallBacks)

	return ptr
}

// region C Code
var cfDictionaryCreate func(allocator uintptr, keys *uintptr, values *uintptr, length int, keyCallBack, valueVallBack uintptr) CFDictionary

var (
	kCFTypeDictionaryKeyCallBacks   uintptr
	kCFTypeDictionaryValueCallBacks uintptr
)

func init() {
	corefoundation, err := purego.Dlopen("/System/Library/Frameworks/CoreFoundation.framework/CoreFoundation", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	// extern CFDictionaryRef CFDictionaryCreate(CFAllocatorRef allocator, const void * * keys, const void * * values, CFIndex numValues, const CFDictionaryKeyCallBacks * keyCallBacks, const CFDictionaryValueCallBacks * valueCallBacks);
	purego.RegisterLibFunc(&cfDictionaryCreate, corefoundation, "CFDictionaryCreate")

	load := func(name string) uintptr {
		return utility.Load(corefoundation, name)
	}
	kCFTypeDictionaryKeyCallBacks = load("kCFTypeDictionaryKeyCallBacks")
	kCFTypeDictionaryValueCallBacks = load("kCFTypeDictionaryValueCallBacks")
}

// endregion C
