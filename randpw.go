package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Empty map for checking if a character has been used more than twice
var uses = map[string]int{}

// Arrays for the characters
var alfabet = [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
var phonetic = [...]string{"alfa", "bravo", "charlie", "delta", "echo", "foxtrot", "golf", "hotel", "india", "juliet", "kilo", "lima", "mike", "november", "oscar", "papa", "quebec", "romeo", "sierra", "tango", "uniform", "victor", "whiskey", "x-ray", "yankee", "zulu"}
var numbers = [...]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

// Set our rng
var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// Main function for the password generator
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage is 'randpw [length]'")
		os.Exit(1)
	}

	// Check if the argument is a number
	length, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Not a valid number")
		os.Exit(1)
	}

	var password string
	var spelling string

	for i := 0; i < length; i++ {
		case_decider := r.Intn(3)

		switch case_decider {
		case 0:
			value := r.Intn(len(alfabet))
			uses[alfabet[value]]++

			if uses[alfabet[value]] > 2 {
				value = reroll(value, len(alfabet))
			}

			password += alfabet[value]
			spelling += phonetic[value]

		case 1:
			value := r.Intn(len(alfabet))
			uses[strings.ToUpper(alfabet[value])]++

			if uses[strings.ToUpper(alfabet[value])] > 2 {
				value = reroll(value, len(alfabet))
			}

			password += strings.ToUpper(alfabet[value])
			spelling += strings.ToUpper(phonetic[value])

		case 2:
			value := r.Intn(len(numbers))
			uses[numbers[value]]++

			if uses[alfabet[value]] > 2 {
				value = reroll(value, len(numbers))
			}

			password += numbers[value]
			spelling += numbers[value]
		}
		spelling += " "

	}

	fmt.Println(password)
	fmt.Println(spelling)
}

// Function to recursively check for a new character. Bound to take a while for very large passwords ¯\_(ツ)_/¯
func reroll(value int, length int) int {
	new_value := r.Intn(length)
	if new_value == value {
		new_value = reroll(new_value, length)
	}
	return new_value
}
