package main

import (
    "fmt"
    "github.com/hserge/helper/string"
)

func main() {
    s:=string.ToggleCase("Ghoper")
    fmt.Println("Converted string is : ", s)
}