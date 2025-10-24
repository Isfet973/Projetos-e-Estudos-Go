package main

import "fmt"

func main() {
	var num1 float64
	var num2 float64
	var ops string

	fmt.Println("Qual operação você deseja fazer: '+', '-', '*' ou '/'?")
	fmt.Scan(&ops)

	fmt.Println("Insira o primeiro número: ")
	fmt.Scan(&num1)
	
	fmt.Println("Insira o segundo número: ")
	fmt.Scan(&num2)

	switch ops {
	case "+":
		fmt.Println("O resultado da soma deu:", num1 + num2)
	case "-":
		fmt.Println("O resultado da subtração deu:", num1 - num2)
	case "*":
		fmt.Println("O resultado da multiplicação deu:", num1 * num2)
	case "%":
		fmt.Println("O resultado da divisão deu:", num1 / num2)
	}
	
}