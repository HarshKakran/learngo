package main

import (
	"fmt"
	"log"
)

// custom error
type norgatreMathError struct {
	lat  string
	long string
	err  error
}

func (n norgatreMathError) Error() string {
	return fmt.Sprintf("a norgate math error occurred: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Fatal(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// return 0, errors.New("sqrt of negative not allowed")
		// return 0, fmt.Errorf("sqrt of %v not allowed", f)
		nme := fmt.Errorf("sqrt of negative number not allowed")
		return 0, norgatreMathError{"50.3 N", "99.3 W", nme}
	}
	return f * f, nil
}
