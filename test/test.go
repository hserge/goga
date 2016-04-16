package main

import (
    "fmt"
    "github.com/hserge/goga/helper/string"
)

func main() {
    s:=string.ToggleCase("Ghoper")
    fmt.Println("Converted string is : ", s)
}