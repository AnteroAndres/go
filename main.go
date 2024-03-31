package main

import (
	"fmt"
	"runtime"

	"github.com/antero/go/ejercicios"
	"github.com/antero/go/variables"
)

func main(){
	estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)
	
	if os := runtime.GOOS; os == "linux" || os == "Os X."{
		fmt.Println("Esto no es Windows", os)
	} else {
		fmt.Println("Esto es Windows")
	}

	switch os:= runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("Esto es Darwin")	
	default:
		fmt.Printf("%s \n", os)
	}
	numero, texto := ejercicios.ConvNumerico("500")
	fmt.Println(numero)
	fmt.Println(texto)
}