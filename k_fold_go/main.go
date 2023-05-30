package main

import (
	"fmt"
)

func main() {
	var X = [][]float64{
		{1.1, 2.2, 3.3, 4.4},
		{1.1, 2.2, 3.3, 4.4},
		{1.1, 2.2, 3.3, 4.4},
		{1.1, 2.2, 3.3, 4.4},
		{1.1, 2.2, 3.3, 4.4},
		{1.1, 2.2, 3.3, 4.4},
		{1.1, 2.2, 3.3, 4.4},
		{1.1, 2.2, 3.3, 4.4},
		{1.1, 2.2, 3.3, 4.4},
		{1.1, 2.2, 3.3, 4.4},
	}

	var y = []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("This is the main function for k-fold")

	returned1, returned2 := validateRow(X, y)
	fmt.Printf("\n validateRow(), returned %v, %v \n", returned1, returned2)

	kFoldCrossValidation(X, y, 5)

}

// k fold cross validation for X: 2D array, y: 1D array
func kFoldCrossValidation(X [][]float64, y []float64, k uint) {

	fmt.Printf("This had %v folds", k)

}

// Validate Row: raise error if input array lengths are not equal
// TODO validateRow must raise ValueError if len(x) != len(y)
func validateRow(X [][]float64, y []float64) (bool, error) {
	if len(X) != len(y) {
		fmt.Println("Raise Value Error")
		return false, nil

	}

	return true, nil

}
