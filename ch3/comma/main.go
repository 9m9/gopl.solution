package main

import (
	"bytes"
	"fmt"
)

// exercise 3.10

func comma(num string) string {
	if len(num) <= 3 {
		return num
	}

	var buf bytes.Buffer
	var j = 0

	for i := len(num) % 3; j < len(num); i += 3 {
		buf.WriteString(num[j:i])
		if i != len(num) {
			buf.WriteByte(',')
		}
		j = i
	}

	return buf.String()
}

func main() {
	fmt.Println(comma("34"))
	fmt.Println(comma("1234"))
	fmt.Println(comma("12345678"))
}
