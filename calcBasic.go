package mycalculator

import (
	"fmt"
	"strconv"
)

func parseString(operator string) (int, error) {
	result, err := strconv.Atoi(operator)
	return result, err
}

func operate(operation string, value1 int, value2 int) {
	switch operation {
	case "+":
		fmt.Println("Add: ", value1+value2)
	case "-":
		fmt.Println("Sub: ", value1-value2)
	case "*":
		fmt.Println("Mult: ", value1*value2)
	case "/":
		fmt.Println("Div: ", value1/value2)
	default:
		fmt.Println(operation, "operation is not supported!")

	}

}
