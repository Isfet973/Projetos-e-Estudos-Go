package main

import (
	"fmt"
)

func main() {
	produto := [3]string{"Produto 1 = R$ 120.00", "Produto 2 = R$ 125.75", "Produto 3 = R$ 150.99"}
	var produto_sel int
	var quantidade float64

	for i := 0; i < 3; i++ {
		fmt.Println((produto[i]))
	}

	fmt.Print("Selecione o produto desejado > ")
	fmt.Scanln(&produto_sel)
	
	
	switch produto_sel {
	case 1:
		produto := 120.00
		fmt.Print("Quantidade > ")
		fmt.Scan(&quantidade)
		if quantidade >= 2 {
			total := quantidade * produto
			fmt.Printf("O valor total do produto é de > R$ %.2f\n", total,)
			desconto := total * 0.20
			total_desconto := total - desconto
			fmt.Printf("O produto recebeu um desconto de R$ %.2f pela compra de mais de 2 itens!\n", desconto)
			fmt.Printf("O valor total do produto agora é de > R$ %.2f", total_desconto)
		} else {
			fmt.Printf("A compra deu > R$ %.2f", produto)
		}
	case 2:
		produto := 125.75
		fmt.Print("Quantidade > ")
		fmt.Scan(&quantidade)
		if quantidade >= 5 {
			total := quantidade * produto
			fmt.Printf("O valor total do produto é de > R$ %.2f\n", total,)
			desconto := total * 0.25
			total_desconto := total - desconto
			fmt.Printf("O produto recebeu um desconto de R$ %.2f pela compra de mais de 5 itens!\n", desconto)
			fmt.Printf("O valor total do produto agora é de > R$ %.2f", total_desconto)
		} else {
			fmt.Printf("A compra deu > R$ %.2f", produto)
		}
	case 3:
		produto := 150.99
		fmt.Print("Quantidade > ")
		fmt.Scan(&quantidade)
		if quantidade >= 10 {
			total := quantidade * produto
			fmt.Printf("O valor total do produto é de > R$ %.2f\n", total,)
			desconto := total * 0.40
			total_desconto := total - desconto
			fmt.Printf("O produto recebeu um desconto de R$ %.2f pela compra de mais de 10 itens!\n", desconto)
			fmt.Printf("O valor total do produto agora é de > R$ %.2f", total_desconto)
		} else {
			fmt.Printf("A compra deu > R$ %.2f", produto)
		}
	}
}	