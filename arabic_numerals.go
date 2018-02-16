package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func convert_arabic_to_roman(arg int)(string) {
base_ten := []int{1000,100,10,1}

roman_map_1 := []string{
1 : "I",
10 : "X",
100 : "C",
1000 : "M",
}
roman_map_2 := []string{
1 : "V",
10 : "L",
100 : "D",
1000 : "MMMMM",
}

if arg < 0 || arg > 4000 {
return "ROMAN_OUT_OF_RANGE"
}

roman_slice := []string{}
x := ""

for _, f := range base_ten {
digit, map1, map2 := int(arg / f), roman_map_1[f], roman_map_2[f]

switch digit {
case 1: roman_slice = append(roman_slice, string(map1) )
case 2: roman_slice = append(roman_slice, string(map1) + string(map1))
case 3: roman_slice = append(roman_slice, string(map1) + string(map1) + string(map1) )
case 4: roman_slice = append(roman_slice, string(map1) + string(map2) )
case 5: roman_slice = append(roman_slice, string(map2) )
case 6: roman_slice = append(roman_slice, string(map2) + string(map1) )
case 7: roman_slice = append(roman_slice, string(map2) + string(map1)+ string(map1) )
case 8: roman_slice = append(roman_slice, string(map2) + string(map1)+ string(map1)+ string(map1) )
case 9: roman_slice = append(roman_slice, string(map1) + string(x)  )
}

arg -= digit * f
x = map1
}
roman_number := ""
for _, e := range roman_slice {
roman_number += e
}
fmt.Printf("Your number is: %v\n", roman_number)
return roman_number
}

func main(){

	fmt.Print("Enter numeral: ")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		numeral, _ := strconv.Atoi(input)
		convert_arabic_to_roman(numeral)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
