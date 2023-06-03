package main

import (
	"errors"
	"fmt"
)

// Enum-like values for method param in KFoldCrossValidation() to select implementation
type Method int

const (
	Sequential Method = iota
	Concurrent
	Parallel
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

// Serial execution of k fold cross validation
func kFoldLoop(X [][]float64, y []float64, k uint) {
	fmt.Println("Run loop implementation")

	var testStart, testEnd, trainStart, trainEnd int

	fold := int(len(X) / int(k))
	for i := 0; i < int(k); i++ {

		switch i {

		//first iteration
		case 0:
			fmt.Println("First iteration")
			//index 0 to the first fold
			testStart = 0
			testEnd = fold

			//first fold to end of the data
			trainStart = testEnd
			trainEnd = len(X)

			testDataX := X[testStart:testEnd]
			trainDataX := X[trainStart:trainEnd]

			testDataY := y[testStart:testEnd]
			trainDataY := y[trainStart:trainEnd]

			fmt.Println("***INDEXES***")
			fmt.Printf("\n testStart: %v, testEnd: %v \n", testStart, testEnd)
			fmt.Printf("\n trainStart: %v, trainEnd: %v \n", trainStart, trainEnd)

			//data for model
			fmt.Printf("\n Testdata X: %v \n Traindata X: %v \n", testDataX, trainDataX)
			fmt.Printf("\n Testdata Y: %v \n Traindata Y: %v \n", testDataY, trainDataY)

		//last iteration
		case int(k - 1):
			fmt.Println("Last iteration")

		default:
			fmt.Printf("Iteration: %v \n", i)

		}

	}

	fmt.Printf("DELETE LATER__use fold var:  %v \n", fold)
}

// k fold cross validation for X: 2D array, y: 1D array
// TODO: implement k fold
// Default method implementation: Serial for loop
func kFoldCrossValidation(X [][]float64, y []float64, k uint, method Method) {

	//row validation
	err := HandleRows(X, y)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	switch method {
	case Concurrent:
		fmt.Println("Run concurrent implementation")

	case Parallel:
		fmt.Println("Run parallel implementation")

	default:
		kFoldLoop(X, y, k)
	}

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

	var y = []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//var y = []float64{1, 2.5}

	fmt.Println("This is the main function for k-fold")

	kFoldCrossValidation(X, y, 10, Sequential)

}
