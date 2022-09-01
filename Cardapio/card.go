package main

import "fmt"

func main(){

	

	var op,opc int
	var i int
	total:= 0.00

	opcoes()

	//Switch 1
	fmt.Scan(&opc)
	switch opc{
		case 1: 

	cardapio()
	for  i=0; op != 4; i++ {
	//Switch 2
	fmt.Scan(&op)
	switch op {			
		
	case 1:

		fmt.Println("Você escolheu pizza!")
		cardapio()
		total = total + 20


	case 2:
		fmt.Println("Você escolheu refri lata!")
		cardapio()
		total = total + 5
		

				
	case 3:
		fmt.Println("Você escolheu uma sobremesa!")
		cardapio()
		total = total + 8
	

	case 4:
		fmt.Println("\nSua conta deu: ",total)
		println("\nPrograma encerrado...\n")	
										
	
	default:
		fmt.Println("Escolha uma opção válida.")
	}
	}
	case 2: 
		dedo()

	case 3: 
		fmt.Println("Programa encerrado...")
	
	}
}

func cardapio(){

	card := "----Cardápio----\n"+
		  "[1] - Pizza R$ 20,00\n"+
		  "[2] - Lata R$ 5,00\n"+
		  "[3] - Sobremesa R$ 8,00\n"+
		  "[4] - Conta\n"

	fmt.Println(card)

}
func opcoes(){

	opcs := "----Escolha uma opção----\n"+
		"[1] - Cardapio\n"+
		"[2] - Dedo\n"+
		"[3] - Sair\n"

	fmt.Println(opcs)

}
func dedo(){

	dedo:= " ┏┓\n"+
    	   " ┃┃\n"+
    	   " ┃┃\n"+
		  "┏┫┣┳┓\n"+
		  "┣┻┫┃┃\n"+
          "┃┏┻┻┫\n"
	fmt.Println(dedo)
}
