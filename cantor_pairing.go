package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"time"
)

func cantor_pair(x, y int) (int, error) { // for encoding two numbers
	if x < 0 || y < 0 {
		return 0, errors.New("both numbers must be positive integers")
	}
	return (x+y)*(x+y+1)/2 + y, nil
}
func iverse_cantor(z int) (int, int, error) { // for decoding a encoded number to 2 orginal numbers
	if z < 0 {
		return 0, 0, errors.New("input must be a postive integer")
	}
	w := int((math.Sqrt(float64(8*z+1)) - 1) / 2)
	t := (w*w + w) / 2
	y := z - t
	x := w - y
	return x, y, nil
}
func input(p string) int { // input pormpts to obligate user to enter a positive integer
	for {
		var input string
		fmt.Print(p)
		fmt.Scanln(&input)
		value, err := strconv.Atoi(input)
		if err == nil && value >= 0 {
			return value
		}
		fmt.Println("ivalid input please enter postive value :)")
		fmt.Println("---------------------------------------------------------------")
	}
}
func loading(m string, delay time.Duration) { //animation
	for _, char := range m {
		fmt.Print(string(char))
		time.Sleep(delay)
	}
	fmt.Println()
}
func main() {
	loading("------------------Cantor Pairing Function Program------------------", time.Millisecond*1)
	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("(1) Encode two numbers")
		fmt.Println("(2) Decode a number")
		fmt.Println("(3) Exit")
		var choice string
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			loading("------------------------------cantor pairing------------------------------", time.Millisecond*10)
			x := input("enter the first positive integer: ")
			y := input("enter the second positive integer: ")
			z, err := cantor_pair(x, y)
			if err != nil {
				fmt.Println("error:", err)
			} else {
				fmt.Printf("encoded result is : z = %d\n", z)
			}
		case "2":
			loading("------------------------------The inverse of cantor pairing----------------------------", time.Millisecond*10)
			z := input("enter the encoded number: ")
			x, y, err := iverse_cantor(z)
			if err != nil {
				fmt.Println("error:", err)
			} else {
				fmt.Printf("decoded numbers are : x = %d, y = %d\n", x, y)
			}
		case "3":
			fmt.Println("-----------------exiting the program!-----------------")
			loading("__________________Goodbye!__________________", time.Millisecond*10)
			return
		default:
			fmt.Println("invalid select from 1 to 3 ")
		}
	}
}
