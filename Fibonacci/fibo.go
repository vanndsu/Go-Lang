package main

import "fmt"

func main(){

var termo = 1 
var resultado = 1
var i int

   

    /*if(termo == 1){
        printf("0");
    }

    if( termo == 2){
         printf("1");/
    }*/
       for termo != 0 {
         fmt.Println("Digite um termo que queira saber o valor\n");
         fmt.Scan(&termo);

        t2 := 0
        t3 := 1
		

        for i = 2; i <= termo; i++{

            resultado = t2 + t3
           
            t2 = t3
            t3 = resultado
           
        }

       fmt.Println("\n", resultado);

       }
	}
