// Main --> Basics: Marshlling and Umarshalling of JSON
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Fname string `json:"firstName"`
	Lname string `json:"lastName"`
	Age   int    `json:"age"`
}

type Contact struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type CombinedData struct {
	Person  `json:"person"`
	Contact `json:"contact"`
}

func main() {
	p1 := Person{"Harsh", "Kakran", 21}
	con1 := Contact{"harsh42@gmail.com", "123456"}
	combined1 := CombinedData{p1, con1}

	p2 := Person{"Ram", "Suri", 22}
	con2 := Contact{"ram42@gmail.com", "223456"}
	combined2 := CombinedData{p2, con2}

	combinedData := []CombinedData{combined1, combined2}

	// JsonData Marshaling --> Converting struct to JSON data
	jsonData, err := json.Marshal(combinedData)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	fmt.Printf("%v\t%T\n", string(jsonData), jsonData)

	// Unmarshaling the JsonData --> Converting JSON data to struct
	var CombinedDataReceiver []CombinedData
	err = json.Unmarshal(jsonData, &CombinedDataReceiver)
	if err != nil {
		log.Fatalf("Error while unmarshalling --> %s", err)
	}

	fmt.Printf("------------------------------\n%v\t%T", CombinedDataReceiver, CombinedDataReceiver)

}
