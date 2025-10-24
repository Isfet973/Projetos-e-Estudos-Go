package main

import "fmt"

var soma float64

func main() {
	for i := 1; i <= 4; i++ {
		var nota float64

		for {
			fmt.Printf("Digite a nota %d (0 a 10) > ", i)
			fmt.Scan(&nota)

			if nota >= 0 && nota <= 10 {
				break
			}
			fmt.Println("Número inválio!")
		}
		soma += nota
	}
	var media = soma / 4

	if media >= 7 {
		fmt.Println("O aluno foi aprovado! Nota final:", media)
	} else if media >= 5 && media <= 6.9 {
		fmt.Println("O aluno está de recuperação! Nota final:", media)
 	} else {
		fmt.Println("O aluno foi reprovado! Nota final:", media)
	}
}