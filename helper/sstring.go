package sstring

import (
    "bytes"
    "unicode"
)

/*
ToggleCase toggles string case.
*/
func ToggleCase(str string) string {
    buf := bytes.Buffer{}
    
    for _, r := range str {
        if unicode.IsUpper(r) {
            buf.WriteRune(unicode.ToLower(r))
        } else {
            buf.WriteRune(unicode.ToUpper(r))
        }
    }
    
    return buf.String()
}