package main

import (
	"fmt"
)

func main() {
	var intArr = [3]int32{1, 2, 3} // array declaration
	intArr[1] = 123
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])

	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Println(intSlice)
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)

	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice3)

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 8, "Sarah": 45}
	fmt.Println(myMap2["Adam"])
	fmt.Println(myMap2["Jason"]) // will get the default value of the value data type, here it is 0
	var age, ok = myMap2["Jason"]
	delete(myMap2, "Adam")
	if ok {
		fmt.Printf("The age is %v", age)
	} else {
		fmt.Printf("Invalid Name")
	}

	for name, age := range myMap2 {
		fmt.Printf("Name: %v\t Age: %v\n", name, age)
	}

	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v\n", i, v)

		//while loop
		var i int = 0
		for {
			if i >= 10 {
				break
			}
			fmt.Printf("%v\n", i)
			i += 1
		}

		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}
}
