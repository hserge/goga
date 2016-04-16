package main

import (
    "fmt"
    "github.com/hserge/goga/helper/sstring"
)

func main() {
    s:=string.ToggleCase("Ghoper")
    fmt.Println("Converted string is : ", s)
}