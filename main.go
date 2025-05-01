package main

import (
	"fmt"
)

func main() {
	err := SaveAccount("test_purego", "hallo", "welt")
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}
}
