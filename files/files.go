package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/antero/go/ejercicios"
)
var fileName string = "./files/txt/tabla.txt"

func GrabaTabla(){
	var texto string = ejercicios.TabladeMultiplicar()
	archivo, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error al crear el archivo" + err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()
}	


func SumaTabla(){
	var texto string = ejercicios.TabladeMultiplicar()
	if !Append(fileName, texto){
		fmt.Println("Error al concadenar contenido")
	}
}

func Append(filen string, texto string) bool {
	arch, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error al abrir el archivo" + err.Error())
		return false
	}

	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error durante el writeString" + err.Error())
		return false
	}
	arch.Close()
	return true
}

func LeoArchivo(){
	archivo, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error al leer archivo" + err.Error())
		return
	}
	
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan(){
		registro := scanner.Text()
		fmt.Println("> "+registro)
	}
	archivo.Close()
}