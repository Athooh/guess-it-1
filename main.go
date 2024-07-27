package main

import (
	"bufio"
	"fmt"

	"os"
	"strconv"

	"github.com/Athooh/guess-it-1/mathfuncs"
)

func main() {
	var dataSlice []float64

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()
		num, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Invalid input, please enter a valid number.")
			continue
		}
		dataSlice = append(dataSlice, num)

		if len(dataSlice) == 1 {
			fmt.Println(dataSlice[0])
			continue
		}

		prevData := dataSlice[:len(dataSlice)-1]

		med := mathfuncs.Median(prevData)
		stdDev := mathfuncs.StdDeviation(prevData)

		// Define the range using median and standard deviation
		lowerLimit := med - stdDev
		upperLimit := med + stdDev

		fmt.Println(num)
		fmt.Printf("%.0f %.0f\n", lowerLimit, upperLimit)
	}
}
