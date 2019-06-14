package main

import(

	"fmt"

)

type obj struct{
	value string
}

func (o *obj) update(){
	fmt.Printf("endereço da variavel na função1: %x\n", &o.value)
	o.value = "novo valor 1"

}

func (o obj) update2(){
	o.value="novo valor2"
	fmt.Printf("endereço da variavel na função2: %x\n", &o.value)
	fmt.Printf("valor dentro da funcao 2: %s\n", o.value)
}

func main(){

	o := obj{
		value: "valor original",
	}

	fmt.Println("1", o.value)
	o.update()
	fmt.Println("2", o.value)
	fmt.Printf("endereço da variavel na função main: %x\n", &o.value)
	o.update2()
	fmt.Println("3", o.value)
	
}