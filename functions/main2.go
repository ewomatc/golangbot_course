package main2

import (
	"fmt"
)

func rectProps(length, width float64)(float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

// using named return values
func rectProps(length, width float64) (area, perimeter float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return //no explicit return value. It return the area and perimiter declared after the parameters automatically
}

func main() {
 	area, perimeter := rectProps(10.8, 5.6)
	fmt.Printf("Area %f Perimeter %f", area, perimeter)	
}

// using blank identifier '_'
func main() {
	area, _ := rectProps(10.8, 5.6) // perimeter is discarded, the code uses only the area
	fmt.Printf("Area %f ", area)
}
GO
