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

// medir: 1 calentamiento y 3 corridas; devuelve los 3 tiempos en ms.
func medir(resolver func([]int) int, datos [][]int) [3]float64 {
	consultar := func() {
		for i := 0; i < consultas; i++ {
			for _, rend := range datos {
				resolver(rend)
			}
		}
	}
	consultar() // calentamiento
	var t [3]float64
	for i := range t {
		inicio := time.Now()
		consultar()
		t[i] = time.Since(inicio).Seconds() * 1000
	}
	return t
}

func tabla(titulo string, meses []int, tiempos [][3]float64) {
	fmt.Printf("\n%s\n\n", titulo)
	fmt.Println("| Meses | Corrida 1 (ms) | Corrida 2 (ms) | Corrida 3 (ms) | Promedio (ms) | Veces más |")
	fmt.Println("| ---: | ---: | ---: | ---: | ---: | ---: |")
	anterior := 0.0
	for i, t := range tiempos {
		prom := (t[0] + t[1] + t[2]) / 3
		veces := "—"
		if anterior > 0 {
			veces = fmt.Sprintf("%.2f", prom/anterior)
		}
		fmt.Printf("| %d | %.3f | %.3f | %.3f | %.3f | %s |\n", meses[i], t[0], t[1], t[2], prom, veces)
		anterior = prom
	}
}

func main() {
	r := rand.New(rand.NewSource(42))
	todos := []int{6, 12, 24, 48}
	var bruta, dina [][3]float64
	for _, meses := range todos {
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

		bruta = append(bruta, medir(fuerzaBruta, datos))
		dina = append(dina, medir(dinamica, datos))
	}
	tabla("Fuerza bruta", todos, bruta)
	tabla("Programación dinámica", todos, dina)
	fmt.Println("\nLas dos versiones dieron la misma mejor racha en las 3 sucursales.")
}
