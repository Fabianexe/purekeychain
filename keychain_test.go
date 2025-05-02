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
	// output: login: hallo
	// password: welt
}
