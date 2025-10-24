package main

import "fmt"

var peso float64
var altura float64
var imc float64

func main(){
	fmt.Println("Informe seu peso em KG: ")
	fmt.Print("> ")
	fmt.Scan(&peso)

	fmt.Println("Informe sua altura em metros: ")
	fmt.Print("> ")
	fmt.Scan(&altura)

	imc = peso / (altura * altura)
	fmt.Printf("Seu IMC é: %.2f", imc)

	if imc < 18.5 {
		fmt.Println(" - Você está abaixo do seu peso ideal.")
		fmt.Println("Classificação: Abaixo do peso")
	} else if imc >= 18.5 && imc < 24.9 {
		fmt.Println(" - Você está no peso ideal para a sua altura")
		fmt.Println("Classificação: Peso normal")
	} else if imc >= 25.0 && imc < 29.9 {
		fmt.Println(" - Você está com sobrepeso")
		fmt.Println("Classificação: Sobrepeso")
	} else if imc >= 30.0 && imc < 34.9 {
		fmt.Println(" - Você está com obesidade em grau I")
		fmt.Println("Classificação: Obesidade grau I")
	} else if imc >= 35.0 && imc < 39.9 {
		fmt.Println(" - Você está com obesidade em grau II")
		fmt.Println("Classificação: Obesidade grau II")
	} else {
		fmt.Println(" - Você está com obesidade em grau III")
		fmt.Println("Classificação: Obesidade grau III")
	}
}