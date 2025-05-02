package cfdata

import (
	"github.com/ebitengine/purego"

	"github.com/Fabianexe/purekeychain/internal/utility"
)

type CFData uintptr

func (d CFData) String() string {
	length := cfDataGetLength(d)
	r := make([]byte, length)
	cfDataGetBytes(d, utility.CFRange{Length: uint64(length)}, r)

	return string(r)
}

func Create(data string) CFData {
	bytes := []byte(data)

	return cfDataCreate(0, bytes, len(bytes))
}

// region C Code
var cfDataGetLength func(CFData) int
var cfDataGetBytes func(CFData, utility.CFRange, []byte) int
var cfDataCreate func(uintptr, []byte, int) CFData

func init() {
	corefoundation, err := purego.Dlopen("/System/Library/Frameworks/CoreFoundation.framework/CoreFoundation", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	// extern CFIndex CFDataGetLength(CFDataRef theData);
	purego.RegisterLibFunc(&cfDataGetLength, corefoundation, "CFDataGetLength")
	// extern void CFDataGetBytes(CFDataRef theData, CFRange range, UInt8 * buffer);
	purego.RegisterLibFunc(&cfDataGetBytes, corefoundation, "CFDataGetBytes")
	// extern CFDataRef CFDataCreate(CFAllocatorRef allocator, const UInt8 * bytes, CFIndex length);
	purego.RegisterLibFunc(&cfDataCreate, corefoundation, "CFDataCreate")
}

// endregion
