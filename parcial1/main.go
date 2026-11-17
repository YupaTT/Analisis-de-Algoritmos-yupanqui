// Gonzalo Yupanqui Torrez
package main

import (
	"fmt"
	"time"
)

func HayComprobantesRepetidosPorPares(comprobantes []int) bool {
	for a := 0; a < len(comprobantes); a++ {
		for b := a + 1; b < len(comprobantes); b++ {
			if comprobantes[a] == comprobantes[b] {
				return true
			}
		}
	}
	return false
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

	mezcla := make([]int, len(datos))
	i, j, k := 0, 0, 0
	for i < len(izquierda) && j < len(derecha) {
		if izquierda[i] <= derecha[j] {
			mezcla[k] = izquierda[i]
			i++
		} else {
			mezcla[k] = derecha[j]
			j++
		}
		k++
	}
	for i < len(izquierda) {
		mezcla[k] = izquierda[i]
		i++
		k++
	}
	for j < len(derecha) {
		mezcla[k] = derecha[j]
		j++
		k++
	}
	copy(datos, mezcla)
}

func HayComprobantesRepetidosCurado(comprobantes []int) bool {
	copia := make([]int, len(comprobantes))
	copy(copia, comprobantes)
	mergeSort(copia)
	for i := 1; i < len(copia); i++ {
		if copia[i-1] == copia[i] {
			return true
		}
	}
	return false
}

func GenerarSinRepetidos(n int) []int {
	datos := make([]int, n)
	for i := 0; i < n; i++ {
		datos[i] = n - i
	}
	return datos
}

func probarCaso(nombre string, datos []int, esperado bool) {
	original := make([]int, len(datos))
	copy(original, datos)
	lenta := HayComprobantesRepetidosPorPares(datos)
	curada := HayComprobantesRepetidosCurado(datos)
	if lenta != esperado || curada != esperado {
		panic("Resultado incorrecto: " + nombre)
	}
	for i := 0; i < len(datos); i++ {
		if datos[i] != original[i] {
			panic("La entrada fue modificada: " + nombre)
		}
	}
	fmt.Println(nombre, "| esperado:", esperado, "| lenta:", lenta, "| cura:", curada)
}

func main() {
	fmt.Println("Gonzalo Yupanqui Torrez")
	probarCaso("Repetido en el medio", []int{1010, 1003, 1009, 1003, 1005}, true)
	probarCaso("Sin repetidos", []int{1010, 1003, 1009, 1004, 1005}, false)
	probarCaso("Lista vacía", []int{}, false)
	probarCaso("Un solo elemento", []int{1008}, false)
	probarCaso("Repetido en los extremos", []int{1004, 1001, 1002, 1006, 1004}, true)

	anteriorLenta := 0.0
	anteriorCurada := 0.0
	for n := 10000; n <= 80000; n = n * 2 {
		datos := GenerarSinRepetidos(n)
		totalLenta := 0.0
		totalCurada := 0.0
		fmt.Println("\nCantidad:", n)

		// Corrida 0: calentamiento. Corridas 1, 2 y 3: mediciones para el promedio.
		for corrida := 0; corrida <= 3; corrida++ {
			inicio := time.Now()
			lenta := HayComprobantesRepetidosPorPares(datos)
			msLenta := time.Since(inicio).Seconds() * 1000

			inicio = time.Now()
			curada := HayComprobantesRepetidosCurado(datos)
			msCurada := time.Since(inicio).Seconds() * 1000

			if lenta || curada {
				panic("No debería haber repetidos")
			}
			for i := 0; i < n; i++ {
				if datos[i] != n-i {
					panic("La entrada fue modificada")
				}
			}
			fmt.Printf("Corrida %d: lenta %.6f ms | cura %.6f ms\n", corrida, msLenta, msCurada)
			if corrida > 0 {
				totalLenta += msLenta
				totalCurada += msCurada
			}
		}

		promedioLenta := totalLenta / 3
		promedioCurada := totalCurada / 3
		fmt.Printf("Promedio: lenta %.6f ms | cura %.6f ms\n", promedioLenta, promedioCurada)
		if anteriorLenta > 0 {
			fmt.Printf("Respecto del tamaño anterior: lenta %.3fx | cura %.3fx\n",
				promedioLenta/anteriorLenta, promedioCurada/anteriorCurada)
		}
		anteriorLenta = promedioLenta
		anteriorCurada = promedioCurada
	}
}
