package security

import (
	"github.com/ebitengine/purego"
	"github.com/pkg/errors"

	"github.com/Fabianexe/purekeychain/internal/cfdata"
	"github.com/Fabianexe/purekeychain/internal/cfdictionary"
	"github.com/Fabianexe/purekeychain/internal/cfstring"
	"github.com/Fabianexe/purekeychain/internal/utility"
)

// AppendAccountData appends to a go map with C-Types the account data
func AppendAccountData(m map[uintptr]uintptr, account string, password string) {
	m[kSecValueData] = uintptr(cfdata.Create(password))
	m[kSecAttrAccount] = uintptr(cfstring.Create(account))
}

// AppendSearchData appends to a go map with C-Types the searching data
func AppendSearchData(m map[uintptr]uintptr, service string) {
	m[kSecClass] = kSecClassGenericPassword
	m[kSecAttrService] = uintptr(cfstring.Create(service))
}

// AppendReturnData appends to a go map with C-Types the return definition
func AppendReturnData(m map[uintptr]uintptr) {
	m[kSecReturnData] = kCFBooleanTrue
	m[kSecReturnAttributes] = kCFBooleanTrue
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

// Load loads an account given in a C-Dict o the keychain
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

// Update updates an account given in a C-Dict with the given values
func Update(search cfdictionary.CFDictionary, update cfdictionary.CFDictionary) error {
	status := secItemUpdate(search, update)

	if status == 0 {
		// Success
		return nil
	}

	// Make error human-readable
	s := secCopyErrorMessageString(status, nil)

	return errors.New(s.String())
}

// Delete deletes an account given in a C-Dict
func Delete(search cfdictionary.CFDictionary) error {
	status := secItemDelete(search)

	if status == 0 {
		// Success
		return nil
	}

	// Make error human-readable
	s := secCopyErrorMessageString(status, nil)

	return errors.New(s.String())
}

// region C Code
var secItemAdd func(attributes cfdictionary.CFDictionary, result *uintptr) (status int32)
var secItemCopyMatching func(attributes cfdictionary.CFDictionary, result *cfdictionary.CFDictionary) (status int32)
var secCopyErrorMessageString func(status int32, reserved *uintptr) cfstring.CFString
var secItemUpdate func(search cfdictionary.CFDictionary, update cfdictionary.CFDictionary) (status int32)
var secItemDelete func(search cfdictionary.CFDictionary) (status int32)

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

	// OSStatus SecItemUpdate(CFDictionaryRef query, CFDictionaryRef attributesToUpdate);
	purego.RegisterLibFunc(&secItemUpdate, security, "SecItemUpdate")

	// OSStatus SecItemDelete(CFDictionaryRef query);
	purego.RegisterLibFunc(&secItemDelete, security, "SecItemDelete")

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
