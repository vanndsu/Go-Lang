package main

import "fmt"

func main(){

	var opc int
	var i int
	var c int
	

	fmt.Println("-------------Cardápio-------------")
	println("1- Pizza R$ 20,00\n2- Refri lata R$ 5,00\n3- Sobremesa R$ 8,00")

	for i  = 0 ; i < 4 ; i++{

	fmt.Scan(&opc)
	switch opc{
		
		case 1:
			println("Você escolheu pizza!")
			c = 20
			break

		
		case 2:
			println("Você refri lata!")
			c = 5
			break
		case 3:
			println("Você escolheu uma sobremesa!")
			c = 8
			break

		default: 
			println("Escolha uma opção válida: ")

	}
	println("Sua conta deu: ",c)
	
}

	


}
