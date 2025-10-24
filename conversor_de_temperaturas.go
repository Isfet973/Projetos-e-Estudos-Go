package main

import "fmt"

var sel int
var temp float64
var celsius float64
var far float64
var kelvin float64
var rankine float64

func main() {

	
	fmt.Println("Escolhe a conversão de temperatura que deseja fazer (1 ou 2):")
	fmt.Println("> 1 - Celsius para Kelvin")
	fmt.Println("> 2 - Celsius para Fahrenheit")
	fmt.Println("> 3 - Celsius para Rankine")
	fmt.Println("> 4 - Kelvin para Celsius")
	fmt.Println("> 5 - Kelvin para Fahrenheit")
	fmt.Println("> 6 - Kelvin para Rankine")
	fmt.Println("> 7 - Fahrenheit para Celsius")
	fmt.Println("> 8 - Fahrenheit para Kelvin")
	fmt.Println("> 9 - Fahrenheit para Rankine")
	fmt.Println("> 10 - Rankine para Celsius")
	fmt.Println("> 11 - Rankine para Kelvin")
	fmt.Println("> 12 - Rankine para Farenheit")

	fmt.Print("> ")
	fmt.Scanln(&sel)

	

	switch sel {

	// °C para °K
	case 1:
		fmt.Println("Coloque a temperatura que deseja converter:")
		fmt.Print("> ")
		fmt.Scanln(&celsius)

		fmt.Println("A conversão de Graus Celsius para Fahrenheit é de:")
		fmt.Printf("> %.2f°K", celsius + 273.15)

	// °C para °F
	case 2:
		fmt.Println("Coloque a temperatura que deseja converter:")
		fmt.Print("> ")
		fmt.Scanln(&celsius)

		fmt.Println("A conversão de Graus Celsius para Kelvin é de:")
		fmt.Printf("> %.2f°F", (celsius * 1.8) + 32)

	// °C para °R
	case 3:
		fmt.Println("Coloque a temperatura que deseja converter:")
		fmt.Print("> ")
		fmt.Scanln(&celsius)

		fmt.Println("A conversão de Graus Kelvin para Celsius é de:")
		fmt.Printf("> %.2f°R", (celsius + 273.15) * 9.0/5.0)
	
	// °K para °C
	case 4:
		fmt.Println("Coloque a temperatura que deseja converter:")
		fmt.Print("> ")
		fmt.Scanln(&kelvin)

		fmt.Println("A conversão de Graus Kelvin para Fharenheit é de:")
		fmt.Printf("> %.2f°C", kelvin - 273.15)

	// °K para °F
	case 5:
		fmt.Println("Coloque a temperatura que deseja converter:")
		fmt.Print("> ")
		fmt.Scanln(&kelvin)

		fmt.Println("A conversão de Graus Fahrenheit para Kelvin é de:")
		fmt.Printf("> %.2f°F", 1.8 * (kelvin - 273.15) + 32)

	// °K para °R
	case 6:
		fmt.Println("Coloque a temperatura que deseja converter:")
		fmt.Print("> ")
		fmt.Scanln(&kelvin)

		fmt.Println("A conversão de Fahrenheit para Celsius é de:")
		fmt.Printf("> %.2f°R", kelvin * 9.0/5.0)

	// °F para C°
	case 7:
		fmt.Println("Coloque a temperatura que deseja converter:")
		fmt.Print("> ")
		fmt.Scanln(&far)

		fmt.Println("A conversão de Graus Celsius para Fahrenheit é de:")
		fmt.Printf("> %.2f°C", (far - 32) / 1.8)

	// °F para °K
	case 8:
		fmt.Println("Coloque a temperatura que deseja converter:")
		fmt.Print("> ")
		fmt.Scanln(&far)

		fmt.Println("A conversão de Graus Celsius para Kelvin é de:")
		fmt.Printf("> %.2f°K", 5.0/9.0 * (far - 32) + 273.15)

	// °F para °R
	case 9:
		fmt.Println("Coloque a temperatura que deseja converter:")
		fmt.Print("> ")
		fmt.Scanln(&far)

		fmt.Println("A conversão de Graus Kelvin para Celsius é de:")
		fmt.Printf("> %.2f°R", far + 459.67)

	// °R para °C
	case 10:
		fmt.Println("Coloque a temperatura que deseja converter:")
		fmt.Print("> ")
		fmt.Scanln(&rankine)

		fmt.Println("A conversão de Graus Kelvin para Fharenheit é de:")
		fmt.Printf("> %.2f°C", (rankine - 491.67) * 5.0/9.0)

	// °R para °K
	case 11:
		fmt.Println("Coloque a temperatura que deseja converter:")
		fmt.Print("> ")
		fmt.Scanln(&rankine)

		fmt.Println("A conversão de Graus Fahrenheit para Kelvin é de:")
		fmt.Printf("> %.2f°K", rankine * 5.0/9.0)

	// °R para °F
	case 12:
		fmt.Println("Coloque a temperatura que deseja converter:")
		fmt.Print("> ")
		fmt.Scanln(&rankine)

		fmt.Println("A conversão de Fahrenheit para Celsius é de:")
		fmt.Printf("> %.2f°F", rankine - 459.67)
	}
	
}