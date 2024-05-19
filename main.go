package main

// Logical test:
// You are given a message encoded using the following mapping:
//
// 'A' -> 1
// 'B' -> 2
// ...
// 'Z' -> 26
//
// Write a function or algorithm that takes a string of digits and returns the number of ways it can be decoded back
// into its original message.
//
// For example:
//
// - Given the input "12", the possible decodings are "AB" and "L", so the output should be 2.
// - For the input "226", the possible decodings are "BZ", "VF", and "BBF", making the output 3.
// - With the input "0", there are no valid decodings, resulting in an output of 0.
//
// Your solution should efficiently handle larger inputs as well.

import (
	"fmt"
	"log"
)

func parseCodedString(s string) int {
	if len(s) == 0 {
		log.Println("String is empty")
		return 0
	}

	if s[0] == '0' {
		log.Println("String can't begin with 0")
		return 0
	}

	prev1, prev2 := 1, 1

	// starting from indexes 1 -> 2 and 0 -> 2
	for i := 2; i <= len(s); i++ {
		current := 0

		oneDigit := s[i-1 : i]
		log.Println("One is", oneDigit, i-1, "to", i)
		if oneDigit >= "1" && oneDigit <= "9" {
			current += prev1
		}

		twoDigits := s[i-2 : i]
		log.Println("Two is", twoDigits, i-2, "to", i)

		if twoDigits == "00" {
			log.Println("No decoding for 00 sequence")
			return 0
		}

		if twoDigits >= "10" && twoDigits <= "26" {
			current += prev2
		}

		prev2 = prev1
		prev1 = current
	}

	return prev1
}

func main() {
	fmt.Println(parseCodedString("12"))           // Output: 2
	fmt.Println(parseCodedString("226"))          // Output: 3
	fmt.Println(parseCodedString("0"))            // Output: 0
	fmt.Println(parseCodedString("226123450034")) // Output: 0
	fmt.Println(parseCodedString("2261234534"))   // Output: 9
	fmt.Println(parseCodedString("226134"))       // Output: 6
	fmt.Println(parseCodedString("99999"))        // Output: 1
	fmt.Println(parseCodedString("101010"))       // Output: 1
}
