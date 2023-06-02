package main

import (
	"errors"
	"fmt"
)

type ValueError struct {
	rowCountX uint
	rowCountY uint
	Msg       string
}

// Error() returns the error message associated with the ValueError.
// It satisfies the error interface by implementing the Error() method.
// Returns string 'Msg'
func (e *ValueError) Error() string {
	return e.Msg
}

// ValidateRow() raises error if input array lengths are not equal
func ValidateRow(X [][]float64, y []float64) (bool, error) {
	if len(X) != len(y) {

		return false, &ValueError{
			Msg:       "Input arrays X and y must have the same number of rows",
			rowCountX: uint(len(X)),
			rowCountY: uint(len(y)),
		}

	}

	return true, nil

}

// Error wrapper for row validation
// Calls ValidateRow() on input arrays and returns relevenat error message.
// Returns ValueError or Unexpected error message.
// Returns nil for no error
func HandleRows(X [][]float64, y []float64) error {
	_, err := ValidateRow(X, y)
	if err != nil {
		var valErr *ValueError
		switch {
		case errors.As(err, &valErr):
			fmt.Printf("ValueError: X has %d rows, Y has %d rows: %s \n", valErr.rowCountX, valErr.rowCountY, valErr.Error())
		default:
			fmt.Printf("unexpected error: %s \n", err)
		}

		return fmt.Errorf("HandleRows(): failed: %w", err)
	}

	return nil
}

// k fold cross validation for X: 2D array, y: 1D array
// TODO: implement k fold
func kFoldCrossValidation(X [][]float64, y []float64, k uint) {

	fmt.Printf("This had %v folds", k)

}

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

	//var y = []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var y = []float64{1, 2.5}

	fmt.Println("This is the main function for k-fold")

	err := HandleRows(X, y)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	kFoldCrossValidation(X, y, 5)

}
