package c5

import (
	"fmt"
	"time"
)

func insercion(datos []int) {
	for i := 1; i < len(datos); i++ {
		valor := datos[i]
		j := i - 1

		for j >= 0 && datos[j] > valor {
			datos[j+1] = datos[j]
			j--
		}

		datos[j+1] = valor
	}
}

func mergeSort(datos []int) {
	if len(datos) < 2 {
		return
	}

	medio := len(datos) / 2
	izquierda := datos[:medio]
	derecha := datos[medio:]

	mergeSort(izquierda)
	mergeSort(derecha)

	mezcla := make([]int, 0, len(datos))
	i, j := 0, 0

	for i < len(izquierda) && j < len(derecha) {
		if izquierda[i] <= derecha[j] {
			mezcla = append(mezcla, izquierda[i])
			i++
		} else {
			mezcla = append(mezcla, derecha[j])
			j++
		}
	}

	mezcla = append(mezcla, izquierda[i:]...)
	mezcla = append(mezcla, derecha[j:]...)

	copy(datos, mezcla)
}

func quickSort(datos []int) {
	if len(datos) < 2 {
		return
	}

	pivote := datos[len(datos)/2]
	i, j := 0, len(datos)-1

	for i <= j {
		for datos[i] < pivote {
			i++
		}

		for datos[j] > pivote {
			j--
		}

		if i <= j {
			datos[i], datos[j] = datos[j], datos[i]
			i++
			j--
		}
	}

	quickSort(datos[:j+1])
	quickSort(datos[i:])
}

func EjecutarTorneo() {
	fmt.Println("Torneo de ordenamientos")
	fmt.Printf("\n%10s %16s %16s %16s\n",
		"Cantidad", "Inserción (ms)", "MergeSort (ms)", "QuickSort (ms)")

	for _, cantidad := range []int{10_000, 20_000, 40_000, 80_000} {
		datos := make([]int, cantidad)
		for i := range datos {
			datos[i] = cantidad - i
		}

		copiaInsercion := make([]int, cantidad)
		copiaMerge := make([]int, cantidad)
		copiaQuick := make([]int, cantidad)
		totalInsercion, totalMerge, totalQuick := 0.0, 0.0, 0.0

		// La corrida 0 calienta; las corridas 1, 2 y 3 forman el promedio.
		for corrida := 0; corrida <= 3; corrida++ {
			copy(copiaInsercion, datos)
			copy(copiaMerge, datos)
			copy(copiaQuick, datos)

			inicio := time.Now()
			insercion(copiaInsercion)
			tiempoInsercion := time.Since(inicio).Seconds() * 1000

			inicio = time.Now()
			mergeSort(copiaMerge)
			tiempoMerge := time.Since(inicio).Seconds() * 1000

			inicio = time.Now()
			quickSort(copiaQuick)
			tiempoQuick := time.Since(inicio).Seconds() * 1000

			if corrida > 0 {
				totalInsercion += tiempoInsercion
				totalMerge += tiempoMerge
				totalQuick += tiempoQuick
			}
		}

		fmt.Printf("%10d %16.6f %16.6f %16.6f\n", cantidad,
			totalInsercion/3, totalMerge/3, totalQuick/3)
	}
}
