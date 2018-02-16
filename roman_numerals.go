package main

import (
	"strings"
	"fmt"
	"bufio"
	"os"
)


func convert_roman_to_arabic(str_numeral string) int{
	convert_to_arabic := map[string]int{
		"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000,
	}

	str_array := strings.Split(str_numeral, "")
	final_digit := 1000
	number := 0
	for _, i := range str_array {
		value, ok := convert_to_arabic[i]
		if ok {
			if final_digit < value {
				number -= 2 * final_digit
			}
			final_digit = value
			number += final_digit
		}

	}
	fmt.Printf("Your number is: %v", number)
	return number
}

func main(){
	//reading a roman numeral
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Roman numeral: ")
	str_numeral, _ := reader.ReadString('\n')
	convert_roman_to_arabic(strings.ToUpper(str_numeral))
}
