package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func formatPhoneNumber(input string, zerolead bool) string {
	var result strings.Builder
	//check long size
	if len(input) > 10 {
		return input
	}

	if zerolead {
		// Add "-" at positions 3 and 7
		if len(input) > 6 {
			result.WriteString(input[:3] + "-" + input[3:6] + "-" + input[6:])
		} else if len(input) > 3 {
			result.WriteString(input[:3] + "-" + input[3:])
		} else {
			result.WriteString(input)
		}
	} else {
		// Add "-" at positions 2 and 6
		if len(input) > 5 {
			result.WriteString(input[:2] + "-" + input[2:5] + "-" + input[5:])
		} else if len(input) > 2 {
			result.WriteString(input[:2] + "-" + input[2:])
		} else {
			result.WriteString(input)
		}
	}

	//check suffix (handle case "000-123-" Expected "000-123")
	resultString := result.String()
	if strings.HasSuffix(result.String(), "-") {
		resultString = resultString[:len(resultString)-1]
	}
	return resultString
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	i := 0
	for {
		i = i + 1 // count round of program
		fmt.Printf("############################# %d #############################\n", i)
		// Read input from the user
		fmt.Print("Enter telephone number (or type 'exit' to quit): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input) // removed the leading and trailing spaces from the original string

		if input == "exit" { // exit command line
			break
		}

		// Divide case 1 and case 2 based on 0 lead
		zerolead := true
		if strings.HasPrefix(input, "0") { // 0 Prefix zerolead = true
			zerolead = true
		} else { // 1 Prefix zerolead = 0
			zerolead = false
		}

		result := formatPhoneNumber(input, zerolead)
		fmt.Println("Formatted number:", result)

	}
}
