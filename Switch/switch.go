package main

import "fmt"

func main(){

	opcoes()

	var op int
	fmt.Scan(&op)
	switch op{

		case 1:
			fmt.Println("Pizza escolhida")
	
		case 2:
			fmt.Println("Latinha escolhida")

		case 3:
			fmt.Println("Sombresa escolhida")

		case 0:
			fmt.Println("Saindo...")
		
		default:
			fmt.Println("Opção inválida")
	}
	
}

func opcoes(){

	opc:="Escolha uma opção\n"+
		 "[1] - Pizza\n"+
		 "[2] - Latinha\n"+
		 "[3] - Sobremesa\n"+
		 "[0] - Sair\n"

	fmt.Println(opc)
}
