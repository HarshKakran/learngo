package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int
	fmt.Println(intNum)

	var floatNum float64
	// var floatNum float32

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	// Typecasting
	var result float32 = floatNum32 + float32(intNum32)

	fmt.Println(result)

	var myString string = "Hello World"
	var myString2 string = "Hello" + " World"
	var myString3 string = `Hello
world`
	fmt.Println(len(myString))                    // -> This will give the length in bytes
	fmt.Println(utf8.RuneCountInString(myString)) // -> This will give the character length of the string
	// Rune is the data type in GO, which represent characters
	var myRune rune = 'a'
	fmt.Println(myRune)

	var myBool bool = false
	fmt.Println(myBool)

	myVar := "Hey"
	var1, var2 := 1, 2
	var var3, var4 int = 3, 4

	fmt.Println(myString2, myString3, floatNum, myVar, var1, var2, var3, var4)

	const myConst string = "const value, can't change, hence immutable"
	const pi float32 = 3.1412

	fmt.Println(myConst, pi)
}
