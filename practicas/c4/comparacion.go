package c4

import (
	"fmt"
	"math/rand"
	"slices"
	"time"
)

func busquedaLineal(datos []int, objetivo int) int {
	for i, v := range datos {
		if v == objetivo {
			return i
		}
	}
	return -1
}

// medir hace 1 calentamiento y 3 corridas; devuelve los 3 tiempos en ms.
func medir(buscar func()) [3]float64 {
	var tiempos [3]float64
	buscar() // calentamiento
	for i := range tiempos {
		inicio := time.Now()
		buscar()
		tiempos[i] = time.Since(inicio).Seconds() * 1000
	}
	return tiempos
}

func EjecutarComparacion() {
	datos := make([]int, 1_000_000)
	for i := range datos {
		datos[i] = i // ordenados
	}
	claves := make([]int, 1_000)
	for i := range claves {
		claves[i] = rand.Intn(len(datos)) // todas existen
	}

	lineal := medir(func() {
		for _, c := range claves {
			busquedaLineal(datos, c)
		}
	})
	binaria := medir(func() {
		for _, c := range claves {
			slices.BinarySearch(datos, c)
		}
	})

	fmt.Println("Comparación medida: lineal contra binaria")
	fmt.Println("Corrida   Lineal (ms)   Binaria (ms)   Lineal / binaria")
	sumaL, sumaB := 0.0, 0.0
	for i := range lineal {
		fmt.Printf("%d         %.6f    %.6f       %.2fx\n", i+1, lineal[i], binaria[i], lineal[i]/binaria[i])
		sumaL += lineal[i]
		sumaB += binaria[i]
	}
	fmt.Printf("Promedio  %.6f    %.6f       %.2fx\n", sumaL/3, sumaB/3, sumaL/sumaB)
}
