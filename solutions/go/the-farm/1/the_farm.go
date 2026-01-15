package thefarm

import (
    "errors"
    "fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(cal FodderCalculator, cows int) (float64, error) {
    food, err1 := cal.FodderAmount(cows)
    if err1 != nil {
        return 0.0, err1
    }
    factor, err2 := cal.FatteningFactor()
    if err2 != nil {
        return 0.0, err2
    }
    return food*factor/float64(cows), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(cal FodderCalculator, cows int) (float64, error) {
    if cows > 0 {
        return DivideFood(cal, cows)
    } else {
        return 0.0, errors.New("invalid number of cows")
    }
}

// TODO: define the 'ValidateNumberOfCows' function
type InvalidCowsError struct {
  cows int
  message string
}

func (e *InvalidCowsError) Error() string {
    return fmt.Sprintf("%v cows are invalid: %s", e.cows, e.message)
}

func ValidateNumberOfCows(cows int) (error) {
    if cows < 0 {
        return &InvalidCowsError {
            cows: cows,
            message: "there are no negative cows",
        }
    } else if cows == 0 {
        return &InvalidCowsError {
            cows: cows,
            message: "no cows don't need food",
        }
    } else {
        return nil
    }
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
