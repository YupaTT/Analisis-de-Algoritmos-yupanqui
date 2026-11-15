package c3

import (
	"fmt"
	"math/bits"
)

// busquedaBinaria requiere datos ordenados de menor a mayor.
func busquedaBinaria(datos []int, objetivo int) (int, []int) {
	izquierda, derecha := 0, len(datos)-1
	var medios []int

	for izquierda <= derecha {
		medio := izquierda + (derecha-izquierda)/2
		medios = append(medios, datos[medio])

		if datos[medio] == objetivo {
			return medio, medios
		}
		if datos[medio] < objetivo {
			izquierda = medio + 1
		} else {
			derecha = medio - 1
		}
	}

	return -1, medios
}

func EjecutarBusquedaBinaria() {
	datos := []int{3, 8, 12, 19, 25, 31, 40, 47}
	fmt.Println("Búsqueda binaria a mano")
	fmt.Println("Datos ordenados:", datos)

	for _, objetivo := range []int{31, 5} {
		indice, medios := busquedaBinaria(datos, objetivo)
		fmt.Printf("\nBuscar %d\n", objetivo)
		fmt.Println("Medios consultados:", medios)
		fmt.Println("Pasos:", len(medios))
		if indice == -1 {
			fmt.Println("No está en los datos.")
		} else {
			fmt.Printf("Encontrado en el índice %d (posición %d).\n", indice, indice+1)
		}
	}

	fmt.Println("\nMáximo de pasos: floor(log2(n)) + 1, para n > 0.")
	for _, cantidad := range []int{1000, 1000000} {
		fmt.Printf("%d datos: %d pasos.\n", cantidad, bits.Len(uint(cantidad)))
	}
}
