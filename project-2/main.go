package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int16 = 255
	fmt.Println(intNum)

	var floatNum32 float32 = 1231.2
	fmt.Println(floatNum32)

	var myString string = "γ"
	fmt.Println(len(myString))                    // Counts the number of bytes
	fmt.Println(utf8.RuneCountInString(myString)) // Rune stores chars outside of basic ascii chars

	var intNum2 int
	fmt.Println(intNum2)

	const myConst string = "const string" // Consts need to be initialized
	fmt.Println(myConst)

}
