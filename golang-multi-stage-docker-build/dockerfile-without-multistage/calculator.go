package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("Hi All, This is a calculator app created using GO (golang)")
	fmt.Println("Type 'exit' to quit or 'help' for instructions.")
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		fmt.Println("Enter calculation (Example: 1 + 2 -> works with or without spaces): ")
		fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "exit" {
			fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
			fmt.Println("Good Bye")
			fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
			break
		}

		if text == "help" {
			fmt.Println("Instructions:")
			fmt.Println("- Enter a calculation like: 5 + 3, 6*2, 7 / 0")
			fmt.Println("- Allowed operators: +, -, *, /")
			fmt.Println("- You can add or skip spaces: 4+5 or 4 + 5 both work")
			fmt.Println("- Type 'exit' to quit")
			continue
		}

		// Make spacing flexible
		text = strings.ReplaceAll(text, " ", "")
		var operator string
		for _, op := range []string{"+", "-", "*", "/"} {
			if strings.Contains(text, op) {
				operator = op
				break
			}
		}

		if operator == "" {
			fmt.Println("Invalid input. No valid operator found.")
			continue
		}

		parts := strings.Split(text, operator)
		if len(parts) != 2 {
			fmt.Println("Invalid format. Please follow 'number operator number' format.")
			continue
		}

		left, err1 := strconv.ParseFloat(parts[0], 64)
		right, err2 := strconv.ParseFloat(parts[1], 64)
		if err1 != nil || err2 != nil {
			fmt.Println("Invalid numbers. Please try again.")
			continue
		}

		var result float64
		switch operator {
		case "+":
			result = left + right
		case "-":
			result = left - right
		case "*":
			result = left * right
		case "/":
			if right == 0 {
				fmt.Println("Error: Division by zero is not allowed.")
				continue
			}
			result = left / right
		default:
			fmt.Println("Unknown operator. Use +, -, *, or /.")
			continue
		}

		fmt.Printf("Result: %.2f\n", result)
	}
}

