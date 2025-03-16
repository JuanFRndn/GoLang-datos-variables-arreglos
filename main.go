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

	//mapas
	var mapa1 = map[string]string{
		"Colombia":  "Bogota",
		"Venezuela": "Caracas",
	}
	fmt.Println(mapa1)
	fmt.Println(mapa1["Colombia"])

	mapa2 := map[string]string{
		"R": "Red",
		"G": "Green",
		"B": "Blue",
	}
	fmt.Println(mapa2)

	mapa3 := map[string][]string{
		"1234567890": {"Pepito", "Perez", "Bogota", "Colombia"},
		"1324567890": {"Mauro", "Cuartas", "Bogota", "Colombia"},
		"1243567890": {"Juanito", "Alimaña", "Cali", "Colombia"},
		"1235467890": {"Pepito", "Castaño", "Medellin", "Colombia"},
		"1234657890": {"Carlos", "Perez", "Bogota", "Colombia"},
	}
	fmt.Println(mapa3)
	fmt.Println("el usuario 1243567890 tiene la información ", mapa3["1243567890"])

	//agregarle un elemento a un mapa
	mapa1["Argentina"] = "Buenos aires"
	fmt.Println(mapa1)

	//eliminar elementos de un mapa
	delete(mapa2, "R")
	fmt.Println(mapa2)

	fmt.Println(mapa2["R"])
}
