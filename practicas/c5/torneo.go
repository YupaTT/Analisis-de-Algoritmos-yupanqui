package c5

import (
	"fmt"
	"math/rand"
	"slices"
	"time"
)

func burbuja(datos []int) {
	for i := 0; i < len(datos)-1; i++ {
		for j := 0; j < len(datos)-1-i; j++ {
			if datos[j] > datos[j+1] {
				datos[j], datos[j+1] = datos[j+1], datos[j]
			}
		}
	}
}

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

// hundir baja datos[i] hasta que sea mayor que sus hijos (montículo de máximos).
func hundir(datos []int, i, n int) {
	for {
		mayor := i
		izq, der := 2*i+1, 2*i+2
		if izq < n && datos[izq] > datos[mayor] {
			mayor = izq
		}
		if der < n && datos[der] > datos[mayor] {
			mayor = der
		}
		if mayor == i {
			return
		}
		datos[i], datos[mayor] = datos[mayor], datos[i]
		i = mayor
	}
}

func heapSort(datos []int) {
	n := len(datos)
	for i := n/2 - 1; i >= 0; i-- {
		hundir(datos, i, n)
	}
	for fin := n - 1; fin > 0; fin-- {
		datos[0], datos[fin] = datos[fin], datos[0]
		hundir(datos, 0, fin)
	}
}

// medir: 1 calentamiento y 3 corridas, cada una sobre una copia nueva de los datos.
func medir(ordenar func([]int), datos []int) [3]float64 {
	var tiempos [3]float64
	copia := make([]int, len(datos))
	for corrida := 0; corrida <= 3; corrida++ {
		copy(copia, datos)
		inicio := time.Now()
		ordenar(copia)
		ms := time.Since(inicio).Seconds() * 1000
		if !slices.IsSorted(copia) {
			panic("el algoritmo no dejó los datos ordenados")
		}
		if corrida > 0 {
			tiempos[corrida-1] = ms
		}
	}
	return tiempos
}

func EjecutarTorneo() {
	algoritmos := []struct {
		nombre  string
		ordenar func([]int)
	}{
		{"Burbuja", burbuja},
		{"Inserción", insercion},
		{"MergeSort", mergeSort},
		{"QuickSort", quickSort},
		{"HeapSort", heapSort},
	}

	fmt.Println("Torneo de ordenamientos (1 calentamiento, 3 corridas, promedio en ms)")
	r := rand.New(rand.NewSource(42))
	anterior := map[string]float64{}
	for _, n := range []int{5_000, 10_000, 20_000, 40_000} {
		datos := make([]int, n)
		for i := range datos {
			datos[i] = r.Intn(1000)
		}
		fmt.Printf("\nn = %d\n", n)
		fmt.Printf("%-10s %12s %12s %12s %12s %8s\n", "Algoritmo", "Corrida 1", "Corrida 2", "Corrida 3", "Promedio", "Veces")
		for _, a := range algoritmos {
			t := medir(a.ordenar, datos)
			prom := (t[0] + t[1] + t[2]) / 3
			veces := "—"
			if anterior[a.nombre] > 0 {
				veces = fmt.Sprintf("x%.2f", prom/anterior[a.nombre])
			}
			fmt.Printf("%-10s %12.3f %12.3f %12.3f %12.3f %8s\n", a.nombre, t[0], t[1], t[2], prom, veces)
			anterior[a.nombre] = prom
		}
	}
}
