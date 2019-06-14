package main

import(

	"fmt"

)


func main(){

	//cria variavel
	var valor = "teste de ponteiro"
	
	//p é um ponteiro que aponta para valor
	p := &valor

	*p = "teste"

	//o * indica que é um ponteiro
	fmt.Println(valor)

	var num = 27
	var ponteiro *int
	ponteiro = &num
	fmt.Println("Utilizando ponteiros")
	fmt.Printf("Conteudo da variavel valor: %d\n", num)
	fmt.Printf("Endereço da variavel valor: %x\n", &num)
	fmt.Printf("Endereço da variavel ponteiro: %x\n", &ponteiro)
	fmt.Printf("Conteudo que a variavel ponteiro aponta: %d\n", *ponteiro)







}