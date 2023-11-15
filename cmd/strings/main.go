package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = "résumé"
	var indexed = myString[1]

	fmt.Println(myString)
	fmt.Printf("%v, %T", indexed, indexed)

	for i, v := range myString {
		fmt.Println(i, v)
	}

	// it is better to use rune array for the string rather than the string,
	var myString2 = []rune("résumé")

	// String Builder
	var strSlice = []string{"h", "a", "r", "s", "h"}
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Printf("\n%v\n%v", catStr, myString2)

}
