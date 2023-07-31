package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func PrintMe(outText string) {
	p(outText)
	var ind = s.Index(outText, " ")
	p(s.Replace(outText, " ", "", ind))
	p(s.ToUpper((outText)))
}

func main() {
	var text = "hello world"
	PrintMe((text))
}
