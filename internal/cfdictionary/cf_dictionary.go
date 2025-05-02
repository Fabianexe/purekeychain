package cfdictionary

import (
	"github.com/ebitengine/purego"

	"github.com/Fabianexe/purekeychain/internal/utility"
)

type CFDictionary uintptr

// Create creates a C-dict from a go map of C types
func Create[K ~uintptr, V ~uintptr, M ~map[K]V](in M) CFDictionary {
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

// ToMap create a go map of C types from the C-dict
func ToMap[CK ~uintptr, CV ~uintptr, K comparable, V any](in CFDictionary, keyTransformer func(CK) K, valueTransformer func(CV) V) map[K]V {
	length := cfDictionaryGetCount(in)
	keys := make([]uintptr, length)
	values := make([]uintptr, length)
	cfDictionaryGetKeysAndValues(in, &keys[0], &values[0])
	ret := make(map[K]V, length)
	for i := range length {
		key := keyTransformer(CK(keys[i]))
		value := valueTransformer(CV(values[i]))
		ret[key] = value
	}

	return ret
}

// region C Code
var cfDictionaryCreate func(allocator uintptr, keys *uintptr, values *uintptr, length int, keyCallBack, valueVallBack uintptr) CFDictionary
var cfDictionaryGetCount func(CFDictionary) int
var cfDictionaryGetKeysAndValues func(dict CFDictionary, keys *uintptr, values *uintptr)

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
	// extern CFIndex CFDictionaryGetCount(CFDictionaryRef theDict);
	purego.RegisterLibFunc(&cfDictionaryGetCount, corefoundation, "CFDictionaryGetCount")
	// extern void CFDictionaryGetKeysAndValues(CFDictionaryRef theDict, const void * * keys, const void * * values);
	purego.RegisterLibFunc(&cfDictionaryGetKeysAndValues, corefoundation, "CFDictionaryGetKeysAndValues")

	load := func(name string) uintptr {
		return utility.Load(corefoundation, name)
	}
	kCFTypeDictionaryKeyCallBacks = load("kCFTypeDictionaryKeyCallBacks")
	kCFTypeDictionaryValueCallBacks = load("kCFTypeDictionaryValueCallBacks")
}

// endregion C
