package security

import (
	"github.com/ebitengine/purego"
	"github.com/pkg/errors"

	"github.com/Fabianexe/purekeychain/internal/cfdictionary"
	"github.com/Fabianexe/purekeychain/internal/cfstring"
	"github.com/Fabianexe/purekeychain/internal/utility"
)

// CreateSaveDict creates a go map with C-Types for saving data
func CreateSaveDict(service string, account string, password string) map[uintptr]uintptr {
	ret := make(map[uintptr]uintptr, 10)
	ret[kSecClass] = kSecClassGenericPassword
	ret[kSecAttrService] = uintptr(cfstring.Create(service))
	ret[kSecAttrAccount] = uintptr(cfstring.Create(account))
	ret[kSecValueData] = uintptr(cfstring.Create(password))

	return ret
}

// Save saves an account given in a C-Dict o the keychain
func Save(cDict cfdictionary.CFDictionary) error {
	status := secItemAdd(cDict, nil)

	if status == 0 {
		// Success
		return nil
	}

	// Make error human-readable
	s := secCopyErrorMessageString(status, nil)

	return errors.New(s.String())
}

// CreateLoadDict creates a go map with C-Types for loading data
func CreateLoadDict(service string) map[uintptr]uintptr {
	ret := make(map[uintptr]uintptr, 10)
	ret[kSecClass] = kSecClassGenericPassword
	ret[kSecAttrService] = uintptr(cfstring.Create(service))
	ret[kSecReturnData] = kCFBooleanTrue
	ret[kSecReturnAttributes] = kCFBooleanTrue

	return ret
}

// Save saves an account given in a C-Dict o the keychain
func Load(cDict cfdictionary.CFDictionary) (cfdictionary.CFDictionary, error) {
	var result cfdictionary.CFDictionary
	status := secItemCopyMatching(cDict, &result)

	if status == 0 {
		// Success
		return result, nil
	}

	// Make error human-readable
	s := secCopyErrorMessageString(status, nil)

	return 0, errors.New(s.String())
}

// region C Code
var secItemAdd func(attributes cfdictionary.CFDictionary, result *uintptr) (status int32)
var secItemCopyMatching func(attributes cfdictionary.CFDictionary, result *cfdictionary.CFDictionary) (status int32)
var secCopyErrorMessageString func(status int32, reserved *uintptr) cfstring.CFString

var (
	kSecClass                uintptr
	kSecClassGenericPassword uintptr
	kSecAttrService          uintptr
	kSecAttrAccount          uintptr
	kSecValueData            uintptr
	kSecReturnData           uintptr
	kSecReturnAttributes     uintptr
	kCFBooleanTrue           uintptr
)

func init() {
	security, err := purego.Dlopen("/System/Library/Frameworks/Security.framework/Security", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	// OSStatus SecItemAdd(CFDictionaryRef attributes, CFTypeRef * result);
	purego.RegisterLibFunc(&secItemAdd, security, "SecItemAdd")

	// CFStringRef SecCopyErrorMessageString(OSStatus status, void * reserved);
	purego.RegisterLibFunc(&secCopyErrorMessageString, security, "SecCopyErrorMessageString")

	// OSStatus SecItemCopyMatching(CFDictionaryRef query, CFTypeRef * result);
	purego.RegisterLibFunc(&secItemCopyMatching, security, "SecItemCopyMatching")

	load := func(name string) uintptr {
		return utility.Load(security, name)
	}

	kSecClass = load("kSecClass")
	kSecClassGenericPassword = load("kSecClassGenericPassword")
	kSecAttrService = load("kSecAttrService")
	kSecAttrAccount = load("kSecAttrAccount")
	kSecValueData = load("kSecValueData")
	kSecReturnData = load("kSecReturnData")
	kSecReturnAttributes = load("kSecReturnAttributes")
	kCFBooleanTrue = load("kCFBooleanTrue")
}

// endregion
