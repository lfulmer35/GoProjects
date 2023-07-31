package main

import "fmt"

func PrintMe(outText string) {
	fmt.Println(outText)
}

func main() {
	var text = "hello world"
	PrintMe((text))
}
