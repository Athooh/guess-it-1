package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadInput() []float64 {
	scanner := bufio.NewScanner(os.Stdin)
	var dataSlice []float64

	for scanner.Scan() {
		input := scanner.Text()
		num, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Invalid input, please enter a number.")
			continue
		}
		dataSlice = append(dataSlice, num)
	}

	return dataSlice
}
