package main

import (
	"fmt"
	"go/printer"
	"net/http"
	"os"
)


func main(){

	/*nome,idade:= devolveNomeIdade()
		//fmt.Printf("Meu nome é: %s e tenho: %d anos\n",nome,idade)
		fmt.Println("Meu nome é:",nome,"e tenho ",idade,"anos")*/
	
	
	
	exibemenu()
	
	switch num() != 0 {

		case 1:
			iniciaMonitoramento()
		case 2:
			println("Logs")
		case 0:
			println("Saindo")
				os.Exit(0)
		default: 
			println("escolha uma opção válida")




	}
	

}

func exibemenu(){

	menu:= "+--------------------------------+\n"+
		   "|"+"[1]- INICIAR MONITORAMENTO"+"      |\n"+
		   "|"+"[2]- EXIBIR LOGS          "+"      |\n"+
		   "|"+"[0]- SAIR DO PROGRAMA     "+"      |\n"+
		   "+--------------------------------+"

	fmt.Println(menu)
}
func num () int{

	var num int
		fmt.Scan(&num)
	
	return num
}
func iniciaMonitoramento(){

	fmt.Println("Monitorando...")
	site:= "https://random-status-code.herokuapp.com"

	resp,_:= http.Get(site)
	
	if resp.StatusCode == 200{

		println("Site: ",site,"foi corregado com sucesso!")
	}else{
		println("O site",site,"está com problemas. Status code: ", resp.StatusCode)
	}
		

}
/*func devolveNomeIdade() (string,int){

	nome:="Vandim"
	idade:= 22

	return nome,idade
}*/
