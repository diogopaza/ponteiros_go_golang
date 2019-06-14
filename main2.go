package main

import(

	"fmt"

)

type obj struct{
	value string
}

func update(o *obj){
	o.value = "novo valor objeto funçao update"
}

func byRef1(v *string){

	aux := "novo valor byRef1"
	v = &aux
	fmt.Printf("valor na função byRef %q\n", *v)

}

func byVal(v string){
	
	v = "novo valor byVal"
	fmt.Printf("Valor na função byVal %q\n", v )
}



func main(){

	objeto := obj{
		value: "valor original",
	}

	fmt.Println("valor inicial:", objeto.value)
	byRef1(&objeto.value)
	fmt.Printf("valor depois de chamar a função byRef1: %q\n", objeto.value)
	update(&objeto)
	fmt.Printf("Função update: %q\n", objeto.value)
	byVal(objeto.value)
	fmt.Printf("valor depois da chmamar a funçao a byVal %q\n", objeto.value)


}