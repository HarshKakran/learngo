package main

import "testing"

func TestDoMath(t *testing.T) {
	addition := doMath(11, 12, add)
	if addition != 23 {
		t.Errorf("Addition failed")
	}

	subtraction := doMath(12, 11, subtract)
	if subtraction != 1 {
		t.Errorf("Subtraction failed")
	}

}

func TestAdd(t *testing.T) {
	got := add(11, 12)
	want := 23

	if got != want {
		t.Errorf("Addition failed!!")
	}
}
