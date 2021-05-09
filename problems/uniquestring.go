package main
	
import (
    "fmt"
    "os"
	"strconv"
)

func main() {
	input := os.Args[1]
	//run := os.Args[2]

	fmt.Printf(strconv.FormatBool(inplace1(input)))
	fmt.Printf("\n")
	fmt.Printf(strconv.FormatBool(ds1ascii(input)))
	fmt.Printf("\n")
	fmt.Printf(strconv.FormatBool(ds2unicode(input)))
}

func ds1ascii(input string) (bool) {
	arr := make([]bool, 128)

	for _, runeval := range input{

		if arr[runeval] == false {
			arr[runeval] = true
		} else {
			return false
		} 
	}
	return true 
}

func ds2unicode(input string) (bool) {
	arr := make([]byte, 18000)

	for _, runeval := range input{
		mod := runeval % 8
		var mask byte = 1 << mod
		var zero byte = 0
		index := runeval / 8

		if arr[index] & mask == zero {
			arr[index] = arr[index] | mask
		} else {
			return false
		} 
	}
	return true 
}

func inplace1(input string) (bool) {
	for i := 0; i < len(input); i++ {
		for j := i+1; j < len(input); j++ {
			if input[i] == input[j]{
				return false 
			}
		}
	}
	return true
}