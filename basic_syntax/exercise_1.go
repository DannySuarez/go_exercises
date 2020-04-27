package main

import "fmt"

func askAboutUser() {
	var name, city string
	var years int
	var weather bool
	fmt.Println("Enter your Name, City, Years live in City, and a boolean")
	fmt.Scan(&name, &city, &years, &weather)

	fmt.Printf("Hi! My name is %s. I have lived in %s for %d years. They say the weather is amazing, which is %t", name, city, years, weather)
}

func printOddLetters() {
	mySentence := "My awesome GO code"
	fmt.Println(mySentence, "The odd letters are:")
	for index, letter := range mySentence {
		if index%2 != 0 {
			fmt.Println(string(letter)) //yaeoeg oe
		}
	}
}

func main() {
	askAboutUser()
	printOddLetters()
}
