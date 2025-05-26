package purekeychain_test

import (
	"fmt"

	"github.com/Fabianexe/purekeychain"
)

func Example() {
	// Create a new service to interact with keychain
	s := purekeychain.New("test_purego")

	// Save login and password to keychain
	err := s.Save("hallo", "welt")
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	// Load login and password from keychain
	login, password, err := s.Load()
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}
	fmt.Println("login:", login)
	fmt.Println("password:", password)

	// Update login and password in keychain
	err = s.Update("bye", "to you")
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}

	// Load login and password from keychain
	login, password, err = s.Load()
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}
	fmt.Println("login2:", login)
	fmt.Println("password2:", password)

	// Delete service from keychain
	err = s.Delete()
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}

	// output:
	// login: hallo
	// password: welt
	// login2: bye
	// password2: to you
}
