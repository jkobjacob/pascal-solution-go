package main

import (
	"fmt"
)

func generatePascalValues(numRows uint64, startValue uint64) [][]uint64 {
	var arr = make([][]uint64, numRows)
	var i, j uint64
	for i = 0; i < numRows; i++ {
		arr[i] = make([]uint64, i+1)
		arr[i][0], arr[i][i] = startValue, startValue
		for j = 1; j < i; j++ {
			arr[i][j] = (arr[i-1][j] + arr[i-1][j-1])
		}
	}
	return arr
}

func prettyPrintPascalValues(pascalValues [][]uint64) {
	for _, row := range pascalValues {
		fmt.Println(row)
	}
}

func main() {
	var rows uint64
	fmt.Println("Enter the number of pascal values you need")
	fmt.Scan(&rows)

	var startValue uint64
	fmt.Println("Enter the start value, which is 7 in your case. Please enter 7 or any")
	fmt.Scan(&startValue)

	prettyPrintPascalValues(generatePascalValues(rows, startValue))
}

