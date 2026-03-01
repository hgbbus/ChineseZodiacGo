package main

import (
	"fmt"
)

var (
	Stems = [...]string{"Jia", "Yi", "Bing", "Ding", "Wu", "Ji", "Geng", "Xin", "Ren", "Gui"}
	Elements = [...]string{"Wood", "Fire", "Earth", "Metal", "Water"}
	Branches = [...]string{"Zi", "Chou", "Yin", "Mao", "Chen", "Si", "Wu", "Wei", "Shen", "You", "Xu", "Hai"}
	Animals = [...]string{"Rat", "Ox", "Tiger", "Rabbit", "Dragon", "Snake", "Horse", "Goat", "Monkey", "Rooster", "Dog", "Pig"}
)

func GetElement(stemIndex int) string {
	return Elements[stemIndex/2]
}

func GetYinYang(stemIndex int) string {
	if stemIndex%2 == 0 {
		return "Yang"
	}
	return "Yin"
}

func mod(a, b int) int {
	return (a % b + b) % b
}

func GetZodiacInfo(year int) (string, string, string, string, string) {
	stemIndex := mod(year - 1984, 10)
	branchIndex := mod(year - 1984, 12)
	
	stem := Stems[stemIndex]
	branch := Branches[branchIndex]
	element := GetElement(stemIndex)
	yinYang := GetYinYang(stemIndex)
	animal := Animals[branchIndex]

	return stem, branch, element, yinYang, animal
}

func GenerateCycle(startYear int) {
	// One solution using existing GetZodiacInfo function
	//for i := range 60 {
	//	year := startYear + i
	//	stem, branch, element, yinYang, animal := GetZodiacInfo(year)
	//	fmt.Printf("%d - %s %s - %s %s - %s\n", year, stem, branch, yinYang, element, animal)
	//}

	// Another solution using direct calculations
	year := startYear
	stemIndex := mod(year - 1984, 10)
	branchIndex := mod(year - 1984, 12)
	fmt.Println()
	for {
		stem := Stems[stemIndex]
		branch := Branches[branchIndex]
		element := GetElement(stemIndex)
		yinYang := GetYinYang(stemIndex)
		animal := Animals[branchIndex]

		fmt.Printf("%d - %s %s - %s %s - %s\n", year, stem, branch, yinYang, element, animal)

		year++
		stemIndex = (stemIndex + 1) % 10
		branchIndex = (branchIndex + 1) % 12

		if year == startYear + 60 {
			break
		}
	}
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
			GenerateCycle(startYear)
		case 2:
			fmt.Print("Enter the Gregorian year to look up: ")
			var year int
			fmt.Scanln(&year)
			stem, branch, element, yinYang, animal := GetZodiacInfo(year)
			fmt.Println()
			fmt.Printf("%d -> %s %s (%s %s %s)\n", year, stem, branch, yinYang, element, animal)
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