package main

import "testing"

func TestSquare(t *testing.T) {
	result := square(10)

	if result != 100 {
		t.Errorf("Error - Expect Square of %d is %d but got %d", 10, 100, result)
	}

}
