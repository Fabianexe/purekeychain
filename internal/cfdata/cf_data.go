package cfdata

import (
	"github.com/ebitengine/purego"
)

type CFData uintptr

func (d CFData) String() string {
	length := cfDataGetLength(d)
	r := make([]byte, length)
	cfDataGetBytes(d, cfRange{0, uint64(length)}, r)

	return string(r)
}

type cfRange struct {
	location uint64
	length   uint64
}

// region C Code
var cfDataGetLength func(CFData) int
var cfDataGetBytes func(CFData, cfRange, []byte) int

func init() {
	corefoundation, err := purego.Dlopen("/System/Library/Frameworks/CoreFoundation.framework/CoreFoundation", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	// extern CFIndex CFDataGetLength(CFDataRef theData);
	purego.RegisterLibFunc(&cfDataGetLength, corefoundation, "CFDataGetLength")
	// extern void CFDataGetBytes(CFDataRef theData, CFRange range, UInt8 * buffer);
	purego.RegisterLibFunc(&cfDataGetBytes, corefoundation, "CFDataGetBytes")
}

// endregion
