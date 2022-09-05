package main


import (
    "fmt"
	"os"
)


func main(){
	
	//imp()
	//calc()	
	var resultado float32
	//var op int	
	
	opcs()

	switch num() {		

		case 1:
			fmt.Println("Operação SOMA escolhida!")
				imp()
					resultado = soma()
						fmt.Printf("Resultado: %.1f\n",resultado)
		case 2: 
			fmt.Println("Operação SUBTRAÇÃO escolhida!")
				imp()
					resultado = sub()
						fmt.Printf("Resultado: %.1f\n",resultado)
		case 3:
			fmt.Println("Operação MULTIPLICAÇÃO escolhida!")
				imp()
					resultado = mult()
						fmt.Printf("Resultado: %.1f\n",resultado)
			
		case 4:
			fmt.Println("Operação DIVISÃO escolhida!")	
				imp()
					resultado = div()
						fmt.Printf("Resultado: %.1f\n",resultado)
		case 0: 
			fmt.Println("Programa encerrado.")
				os.Exit(0)

		default:	
			fmt.Println("Não conheço este comando.")
				os.Exit(-1)

	}	
	
}

func soma()float32{

	var n1,n2 float32
	fmt.Scan(&n1)
		println("+")
			fmt.Scan(&n2)
	

	return n1+n2	
}
func sub()float32{

	var n1,n2 float32
	fmt.Scan(&n1)
		println("-")
			fmt.Scan(&n2)
	

	return n1-n2	
}
func mult()float32{

	var n1,n2 float32
	fmt.Scan(&n1)
		println("×")
			fmt.Scan(&n2)
	

	return n1*n2	
}
func div()float32{

	var n1,n2 float32
	
	fmt.Scan(&n1)
		println("÷")
			fmt.Scan(&n2)

	return n1/n2	
}


func imp(){

	imp:="\nDigite um número"

	fmt.Println(imp)
}
/*func impe(){

	imp:="Digite o segundo número"

	fmt.Println(imp)
}*/

func num () int{

	
	var num int
	fmt.Scan(&num)

	return num
}

func opcs(){

	opc:="Escolha uma operação\n"+
		 "[1]SOMA\n"+
		 "[2]SUBTRAÇÃO\n"+
		 "[3]MULTIPLICAÇÃO\n"+
		 "[4]DIVISÃO\n"+
		 "[0]SAIR DO PROGRAMA"
	
	fmt.Println(opc)

}
