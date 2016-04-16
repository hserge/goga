package main

import (
	"fmt"

	helper "github.com/hserge/goga/helper"
)

func main() {
	s := helper.ToggleCase("Ghoper")
	fmt.Println("Converted string is : ", s)
}
