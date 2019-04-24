package main

import(

	"fmt"

)


func main(){

	
	nome:= "lucas magnum"
	/*
	fmt.Printf("%v\n", nome)
	

	var ponteiro *string

	ponteiro = &nome
	fmt.Println("Endereço de memória: ", &nome)
	fmt.Println("Endereço no ponteiro: ", ponteiro)

	fmt.Println("Endereço no endereço de memória: ", *ponteiro)
	*/

	var ponteiro *string
	ponteiro = &nome
	fmt.Println("Endereço de memória: ", &nome)
	fmt.Println("Endereço no ponteiro: ", ponteiro)

	*ponteiro = "Golang"
	fmt.Println("Nome: ", nome)


}