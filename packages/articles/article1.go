package artist

import (
	"fmt"
	"strings"
)

var Name string = "Samuel"

func printNumber(val int) {
	fmt.Println("Current Number", val)
}

func Add(val ...int) int {
	sum := 0
	for _, v := range val{
		printNumber(v)
		sum += v
	}
	return sum
}

func MakeExcited(val string) string{
	return strings.ToUpper(val)
}