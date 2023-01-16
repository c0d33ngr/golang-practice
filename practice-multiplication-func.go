package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num, numOfTimes int

	fmt.Println("Enter number to generate multiplication table")
	fmt.Scanln(&num)

	fmt.Println("Enter number of times the table should stop")
	fmt.Scanln(&numOfTimes)

	multiplyFunc(num, numOfTimes)
	fmt.Println("Program run successfully")
}

func multiplyFunc(num, iteration int) {
	for i := 1; i <= iteration; i++ {
		fmt.Println(strconv.Itoa(num), "*", strconv.Itoa(i), "=", num * i)
	}
}
