package main

import (
	"errors"
	"fmt"
)

func main() {
	var printVal string = "Hey!"
	printMe(printVal)

	var res, rem, err = intDiv(11, 2)
	/*
		if err != nil {
			fmt.Println(err.Error())
		} else if remainder == 0 {
			fmt.Println("THe result of the interger div is %v", res)
		} else {
			fmt.printf("The result of dicision is %v with remainder %v", res, rem)
		}
	*/

	switch {
	case err != nil:
		fmt.Println(err.Error())
	case rem == 0:
		fmt.Printf("The result of the integer div is %v", res)
	default:
		fmt.Printf("The result of the div is %v and the remainder is %v", res, rem)
	}

	switch rem {
	case 0:
		fmt.Printf("Completely divisible!")
	default:
		fmt.Printf("Not Divisible")
	}
}

func printMe(printVal string) {
	fmt.Println("I am executed, ", printVal)
}

func intDiv(num int, den int) (int, int, error) {
	var err error
	if den == 0 {
		err = errors.New("can't divide by zero")
		return 0, 0, err
	}
	return num / den, num % den, err
}
