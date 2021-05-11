package main

import (
	"os"
	"fmt"
	"strconv"
)

func main() {
	input1 := os.Args[1]

	fmt.Printf(strconv.FormatBool(hashMapPalPerm(input1)))
}

func hashMapPalPerm(input1 string) (bool){
	arr := make([]int, 128) //assuming ascii (in reality this should be letters only, so size 26)

	for _, runeval := range input1{
		arr[runeval] += 1
	}

	odds := false
	for _, val := range arr{
		if val % 2 == 1 {
			if odds {
				return false
			}
			
			odds = true
		}
	}

	return true 
}