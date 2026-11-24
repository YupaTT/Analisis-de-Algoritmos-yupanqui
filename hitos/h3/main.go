package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	sucursales = 3
	consultas  = 1000 // veces que se pide el reporte
)

// fuerzaBruta prueba todos los rangos desde..hasta y suma cada uno desde cero.
func fuerzaBruta(rend []int) int {
	mejor := rend[0]
	for desde := range rend {
		for hasta := desde; hasta < len(rend); hasta++ {
			suma := 0
			for m := desde; m <= hasta; m++ {
				suma += rend[m]
			}
			mejor = max(mejor, suma)
		}
	}
	return mejor
}

// dinamica: la mejor racha que termina este mes es este mes solo
// o la mejor racha del mes anterior más este mes.
func dinamica(rend []int) int {
	racha, mejor := rend[0], rend[0]
	for _, r := range rend[1:] {
		racha = max(r, racha+r)
		mejor = max(mejor, racha)
	}
	return mejor
}

// medir: 1 calentamiento y 3 corridas; devuelve el promedio en ms.
func medir(resolver func([]int) int, datos [][]int) float64 {
	consultar := func() {
		for i := 0; i < consultas; i++ {
			for _, rend := range datos {
				resolver(rend)
			}
		}
	}
	consultar() // calentamiento
	total := 0.0
	for i := 0; i < 3; i++ {
		inicio := time.Now()
		consultar()
		total += time.Since(inicio).Seconds() * 1000
	}
	return total / 3
}

func main() {
	r := rand.New(rand.NewSource(42))
	fmt.Println("| Meses | Fuerza bruta (ms) | Veces más | Dinámica (ms) | Veces más |")
	fmt.Println("| ---: | ---: | ---: | ---: | ---: |")

	antBruta, antDinamica := 0.0, 0.0
	for _, meses := range []int{6, 12, 24, 48} {
		datos := make([][]int, sucursales)
		for s := range datos {
			datos[s] = make([]int, meses)
			ventas := 1000
			for m := range datos[s] {
				anterior := ventas
				ventas += r.Intn(101) - 50      // sube o baja hasta 50 unidades
				datos[s][m] = ventas - anterior // rendimiento = este mes − mes anterior
			}
			if fuerzaBruta(datos[s]) != dinamica(datos[s]) {
				panic("las dos versiones no dan lo mismo")
			}
		}

		bruta := medir(fuerzaBruta, datos)
		dina := medir(dinamica, datos)
		if antBruta == 0 {
			fmt.Printf("| %d | %.3f | — | %.3f | — |\n", meses, bruta, dina)
		} else {
			fmt.Printf("| %d | %.3f | %.2f | %.3f | %.2f |\n", meses, bruta, bruta/antBruta, dina, dina/antDinamica)
		}
		antBruta, antDinamica = bruta, dina
	}
	fmt.Println("\nLas dos versiones dieron la misma mejor racha en las 3 sucursales.")
}
