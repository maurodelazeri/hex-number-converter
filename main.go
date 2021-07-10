package main

import (
	"fmt"
	"strconv"
	"strings"
)

func hexToInt(hexStr string) {
	cleaned := strings.Replace(hexStr, "0x", "", -1)
	result, _ := strconv.ParseUint(cleaned, 16, 64)
	fmt.Println(uint64(result))
}

func intToHex(val int64) {
	fmt.Println("0x" + strconv.FormatInt(val, 16))
}

func main() {
	var i string
	fmt.Print("Enter the hex or the number to convert to hex: ")
	fmt.Scanf("%s", &i)
	if strings.Contains(i, "0x") {
		hexToInt(i)
	} else {
		val, _ := strconv.ParseInt(i, 10, 64)
		intToHex(val)
	}
}
