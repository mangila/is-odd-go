package iseven

import (
	"fmt"
	"testing"
)

func TestIsEven(t *testing.T) {
	numbersOdd := []int{1, 3, 5, 7, 9}
	numbersEven := []int{0, 2, 4, 6, 8}

	for _, v := range numbersOdd {
		if IsEven(v) {
			fmt.Println(v)
			t.Error("Number is not odd!")
		}
	}

	for _, v := range numbersEven {
		if !IsEven(v) {
			fmt.Println(v)
			t.Error("Number is not odd!")
		}
	}
}
