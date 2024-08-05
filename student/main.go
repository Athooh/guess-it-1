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
			fmt.Println("Invalid input, please enter a number.")
			continue
		}
		dataSlice = append(dataSlice, num)

		if len(dataSlice) > 1 {
			lower, upper := mathfuncs.Range(dataSlice)
			fmt.Printf("%d %d\n", lower, upper)
		}
	}
}
