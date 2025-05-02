package cfstring

import (
	"github.com/ebitengine/purego"

	"github.com/Fabianexe/purekeychain/internal/utility"
)

type CFString uintptr

func (c CFString) String() string {
	if c == 0 {
		return ""
	}
	length := cfStringGetLength(c)
	b := make([]uint16, length)

	cfStringGetCharacters(c, utility.CFRange{Length: uint64(length)}, b)
	ret := make([]rune, 0, length)
	for _, u := range b {
		ret = append(ret, rune(u))
	}

	return string(ret)
}

func Create(str string) CFString {
	return cfStringCreateWithCString(0, []byte(str), kCFStringEncodingUTF8)
}

// region C Code

var cfStringGetLength func(cfString CFString) int64
var cfStringGetCharacters func(cfString CFString, length utility.CFRange, target []uint16)
var cfStringCreateWithCString func(allocator uintptr, str []byte, encoding int) CFString

// https://github.com/opensource-apple/CF/blob/master/CFString.h#L129
const kCFStringEncodingUTF8 = 0x08000100

func init() {
	corefoundation, err := purego.Dlopen("/System/Library/Frameworks/CoreFoundation.framework/CoreFoundation", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	// extern CFIndex CFStringGetLength(CFStringRef theString);
	purego.RegisterLibFunc(&cfStringGetLength, corefoundation, "CFStringGetLength")
	// extern void CFStringGetCharacters(CFStringRef theString, CFRange range, UniChar * buffer);
	purego.RegisterLibFunc(&cfStringGetCharacters, corefoundation, "CFStringGetCharacters")
	// extern CFStringRef CFStringCreateWithCString(CFAllocatorRef alloc, const char * cStr, CFStringEncoding encoding);
	purego.RegisterLibFunc(&cfStringCreateWithCString, corefoundation, "CFStringCreateWithCString")
}

// endregion
