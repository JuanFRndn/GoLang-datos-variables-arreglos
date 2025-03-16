package main

import "fmt"

func main() {
	fmt.Println("Gorriending")

	//Variables
	var variable1 string = "Si"
	var variable2 = "No"
	variable3 := "tal vez"
	fmt.Println("variable1 " + variable1)
	fmt.Println("variable1 " + variable2)
	fmt.Println("variable1 " + variable3)

	//comentario
	/* Comentario2*/

	var variable4 int16 = 2024
	var variable5 int8 = 127
	edad := 4231
	fmt.Println("variable1 ", variable4)
	fmt.Println("variable1 ", variable5)
	fmt.Println("variable1 ", edad)

	//Areglos
	var lista1 = [4]string{}
	lista1[1] = "J"
	lista1[0] = "1"
	fmt.Println(lista1)

	lista2 := [3]string{"Miguelito", "Joselito", "juanito"}
	fmt.Println(lista2)

	var lista3 = []string{}
	fmt.Println(lista3)
	lista3 = append(lista3, "Ungria")
	fmt.Println(lista3)
	lista3 = append(lista3, "Ungria1")
	fmt.Println(lista3)
	lista3 = append(lista3, "Ungria2")
	fmt.Println(lista3)
	lista3 = append(lista3, "Ungria3")
	fmt.Println(lista3)
	lista3 = append(lista3, "Ungria4")
	fmt.Println(lista3)

	//copiar una arreglo otro, incluye el elemento incial pero no el final (2,3]
	lista4 := lista3[2:3]
	fmt.Println(lista4)

	//copiar un arreglo a otro, a partir de una posicion
	lista4 = lista3[2:]
	fmt.Println(lista4)

	//copiar un arreglo a otro, desde el inicio hasta una posicion sin incluir la posicion
	lista4 = lista3[:3]
	fmt.Println(lista4)

}
