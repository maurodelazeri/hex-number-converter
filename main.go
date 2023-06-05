package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please input a decimal or hexadecimal number.")
		os.Exit(1)
	}

	input := os.Args[1]

	// If the number starts with 0x, we assume it's hexadecimal
	if len(input) > 2 && input[0:2] == "0x" {
		result := new(big.Int)
		_, success := result.SetString(input[2:], 16)
		if !success {
			fmt.Println("Failed to parse the hexadecimal number")
			os.Exit(1)
		}

		fmt.Printf("Hexadecimal %v is equal to decimal %v\n", input, result)
	} else {
		// Otherwise, we assume it's decimal
		result, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("Failed to parse the decimal number: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Decimal %v is equal to hexadecimal 0x%x\n", input, result)
	}
}
