package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func runeLiterals() {
	s := "After a backslash, certain single character escapes represent special values\nn is a line feed or new line \n\t t is a tab"
	fmt.Println(s)
}

func stringLiterals() {
	s := `After a backslash, certain single character escapes represent
    special values
  n is a line feed or new line
    t is a tab`
	fmt.Println(s)
}

func concatination() {
	s := "Can you hear me?"
	s += "\nHear me screamin'?"
	fmt.Println(s)
}

func convertIntToString() {
	var i int = 1
	var s string = " egg"
	intToString := strconv.Itoa(i)
	var breakfast string = intToString + s
	fmt.Println(breakfast)
}

func largeStringConcatenation() {
	var buffer bytes.Buffer

	for i := 0; i < 500; i++ {
		buffer.WriteString("z")
	}

	fmt.Println(buffer.String())
}

func getNumberOfBytes() {
	s := "hello"
	fmt.Println(len(s))
}

func printByteValue() {
	s := "hello"
	fmt.Println(s[0])
}

func showCharacterOrBinary() {
	s := "hello"
	fmt.Printf("%q\n", s[0])
	// outputs 'h;

	fmt.Printf("%b\n", s[0])
	// outputs 1101000
}

func lower() {
	fmt.Println(strings.ToLower("VERY IMPORTANT MESSAGE"))
}

func substring() {
	fmt.Println(strings.Index("surface", "face"))
	fmt.Println(strings.Index("moon", "aer"))
}

func trim() {
	fmt.Println(strings.TrimSpace("   I don't need all this space     "))
}

func main() {
	runeLiterals()
	stringLiterals()
	concatination()
	convertIntToString()
	largeStringConcatenation()
	getNumberOfBytes()
	printByteValue()
	showCharacterOrBinary()
	lower()
	substring()
	trim()
}
