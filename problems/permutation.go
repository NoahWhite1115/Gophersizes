package main

import (
	"os"
	"fmt"
	"strconv"
)

func main() {
	input1 := os.Args[1]
	input2 := os.Args[2]

	fmt.Printf(strconv.FormatBool(hashMapPerm(input1, input2)))
}

func hashMapPerm(input1 string, input2 string) (bool){
	arr := make([]int, 128)

	for _, runeval := range input1{
		arr[runeval] += 1
	}

	for _, runeval := range input2{
		arr[runeval] -= 1

		if arr[runeval] < 0{
			return false
		}
	}

	for _, val := range arr{
		if val != 0{
			return false
		}
	}

	return true 
}