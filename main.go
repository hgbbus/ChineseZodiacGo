package main

import (
	"fmt"
)

func GetElement(stemIndex int) string {
	elements := []string{"Wood", "Fire", "Earth", "Metal", "Water"}
	return elements[stemIndex/2]
}

func GetYinYang(stemIndex int) string {
	if stemIndex%2 == 0 {
		return "Yang"
	}
	return "Yin"
}

func main() {
	fmt.Println("Welcome to the Chinese Zodiac Program!")

	for {
		fmt.Println()
		fmt.Println("Please select an option:")
		fmt.Println()
		fmt.Println(" 1. Generate and display the full 60-year cycle starting from a specified year.")
		fmt.Println(" 2. Look up the zodiac information for a specific Gregorian year.")
		fmt.Println(" 0. Exit")
		fmt.Println()
		fmt.Print("Enter your choice (1, 2, or 0): ")

		var choice int
		fmt.Scanln(&choice)
		fmt.Println()

		switch choice {
		case 1:
			fmt.Print("Enter the starting Gregorian year: ")
			var startYear int
			fmt.Scanln(&startYear)
			// generateCycle(startYear)
		case 2:
			fmt.Print("Enter the Gregorian year to look up: ")
			var year int
			fmt.Scanln(&year)
			// lookupYear(year)
		case 0:
			fmt.Println("Exiting the program.")
			return
		default:
			fmt.Println("Invalid choice. Please enter 1, 2, or 0.")
		}
	}
}

func Main() {
	main()
}