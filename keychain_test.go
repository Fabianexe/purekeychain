package puregokeychain_test

import (
	"fmt"

	puregokeychain "github.com/Fabianexe/purekeychain"
)

func Example() {
	s := puregokeychain.New("test_purego")
	err := s.SaveData("hallo", "welt")
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	login, password, err := s.LoadData()
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}
	fmt.Println("login:", login)
	fmt.Println("password:", password)

	err = s.UpdateData("bye", "to you")
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}

	login, password, err = s.LoadData()
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}
	fmt.Println("login2:", login)
	fmt.Println("password2:", password)

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
