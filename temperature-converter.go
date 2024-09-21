package main

import "fmt"

func main() {

	var option int
	var temperature float64
	var converted_temperature float64

	fmt.Print("What is the current temperature in celcius: ")
	fmt.Scan(&temperature)

	fmt.Println("To what do you want to convert?")
	fmt.Println("1. Fahrenheit")
	fmt.Println("2. Kelvin")
	fmt.Println("3. Reamur")
	fmt.Scan(&option)

	switch option {
	case 1:
		converted_temperature = (temperature * (9 / 5)) + 32
	case 2:
		converted_temperature = temperature + 273
	case 3:
		converted_temperature = temperature * (4 / 5)
	}
	println("The converted temperature is: ", converted_temperature)

}
