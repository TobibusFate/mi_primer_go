package main

import (
	"fmt"
)

func main() {

	var lista []int

	for i := 0; i < 10; i++ {

		fmt.Print("\033[H\033[2J")
		/*Solicito puntuacion */
		fmt.Printf("Dame el puntaje nº %d: \n", (i + 1))

		/*Escaneo */
		var temp int
		fmt.Scan(&temp)

		/*Chekeo que sea aceptable y vuelvo a solicitar si es necesario */
		for j := true; j; j = (temp < 1 || 5 < temp) {
			if temp < 1 || 5 < temp {

				fmt.Print("\033[H\033[2J")
				fmt.Printf("Los valores deben ser entre 1 y 5")
				fmt.Printf("\nDame nuevamente el puntaje nº %d: \n", (i + 1))
				fmt.Scan(&temp)
			}
		}

		/*Agrego temp en la lista*/
		lista = append(lista, temp)

	}

	fmt.Print("\033[H\033[2J")
	/*Cuento apariciones de cada valor*/
	ap1 := repeticion(lista, 1)
	ap2 := repeticion(lista, 2)
	ap3 := repeticion(lista, 3)
	ap4 := repeticion(lista, 4)
	ap5 := repeticion(lista, 5)

	/*Muestro apariciones de cada valor*/
	fmt.Printf("\n- 1 aparece: %d veces", ap1)
	fmt.Printf("\n- 2 aparece: %d veces", ap2)
	fmt.Printf("\n- 3 aparece: %d veces", ap3)
	fmt.Printf("\n- 4 aparece: %d veces", ap4)
	fmt.Printf("\n- 5 aparece: %d veces", ap5)

	/*Calculo cantidad de bajos y de altos*/
	min, max := contador(ap1, ap2, ap3, ap4, ap5)

	/*Muestro mensaje depende si fueron mas bajos o mas altos*/
	if min < max {
		fmt.Println("\n\n¡Buen resultado!")
	} else {
		fmt.Println("\n\nResultado mejorable")
	}

}

/* Dada una lista y un valor, cuenta las apariciones del valor en la lista*/
func repeticion(list []int, objetivo int) int {

	count := 0
	for i := 0; i < len(list); i++ {
		if list[i] == objetivo {
			count++
		}
	}

	return count
}

/* Toma los valores de cada aparicion en orden y retorna la cantidad de numers bajor y la de numeros altos*/
func contador(lista ...int) (int, int) {
	negativos := lista[0] + lista[1]
	positivos := lista[3] + lista[4]

	return negativos, positivos
}
