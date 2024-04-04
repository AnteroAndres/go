package main

import (
	//"github.com/antero/go/mapas"
	//"github.com/antero/go/users"
	//e "github.com/antero/go/ejer_interfaces"
	//"github.com/antero/go/modelos"
	//d "github.com/antero/go/defer_panic"
	//"github.com/antero/go/goroutines"
	//"github.com/antero/go/webserver"
	"github.com/antero/go/middleware"
)

func main(){
	// estado, texto := variables.ConviertoaTexto(1588)
	// fmt.Println(estado)
	// fmt.Println(texto)
	
	// if os := runtime.GOOS; os == "linux" || os == "Os X."{
	// 	fmt.Println("Esto no es Windows", os)
	// } else {
	// 	fmt.Println("Esto es Windows")
	// }

	// switch os:= runtime.GOOS; os {
	// case "linux":
	// 	fmt.Println("Esto es Linux")
	// case "darwin":
	// 	fmt.Println("Esto es Darwin")	
	// default:
	// 	fmt.Printf("%s \n", os)
	// }
	// numero, texto := ejercicios.ConvNumerico("500")
	// fmt.Println(numero)
	// fmt.Println(texto)

	// teclado.IngresoNumeros()
	// iteraciones.Iterar()
	
	//fmt.Println(ejercicios.TabladeMultiplicar())

	//files.GrabaTabla()

	//files.SumaTabla()
	//files.LeoArchivo()
	//funciones.Calculos()
	//funciones.LlamarClousure()
	//funciones.Exponencia(2)
	//arreglosslices.MuestroArreglo()
	//arreglosslices.Capacidad()
	//mapas.MostrarMapas()	
	//users.AltaUsuario()
	// Pedro := new(modelos.Hombre)
	// e.HumanosRespirando(Pedro)
	
	// Maria := new(modelos.Mujer)
	// e.HumanosRespirando(Maria)
	//d.EjemploPanic()
	// canal1 := make(chan bool)
	// go goroutines.MiNombreLentoo("Antero", canal1)
	// defer func(){
	// 	<- canal1

	// }()
	// fmt.Println("Estoy aqui")

	//webserver.MiWebServer()

	middleware.Middleware()
}