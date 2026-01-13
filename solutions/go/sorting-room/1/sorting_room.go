package sorting

import (
    "strconv"
    "fmt"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

type numberBoxContaining struct {
    num int
}

func (i numberBoxContaining) Number() int {
    return i.num
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
    fancy, ok := fnb.(FancyNumber)
    if !ok {
        return 0
    }
	val, err := strconv.Atoi(fancy.Value())
    if err != nil {
        return 0
    } 
    return val
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	fancy, ok := fnb.(FancyNumber)
    if !ok {
        return "This is a fancy box containing the number 0.0"
    }
    val, err := strconv.Atoi(fancy.Value())
    if err!=nil {
        return ""
    }
    return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(val))
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i any) string {
	switch v := i.(type) {
        case int: 
        	return DescribeNumber(float64(v))
        case float64: 
        	return DescribeNumber(v)
        case NumberBox: 
        	return DescribeNumberBox(v)
        case FancyNumberBox: 
        	return DescribeFancyNumberBox(v)
        default: 
        	return "Return to sender"
    }
}
