package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Linea struct {
	Fecha      string
	SucursalID int
	ProductoID int
	Producto   string
	Cantidad   int
	Subtotal   float64
	Pos        int
}

type Total struct {
	ProductoID int
	Producto   string
	Cantidad   int
	Monto      float64
	Primera    int
}

// reporteActual: copia sin cambios de h1/reporte_actual.go.
func reporteActual(lineas []Linea, sucursalID int, desde, hasta string) []Total {
	var filtradas []Linea
	for _, l := range lineas {
		if (sucursalID == 0 || l.SucursalID == sucursalID) && l.Fecha >= desde && l.Fecha <= hasta {
			filtradas = append(filtradas, l)
		}
	}
	var totales []Total
	for _, l := range filtradas {
		i := 0
		for i < len(totales) && totales[i].ProductoID != l.ProductoID {
			i++
		}
		if i == len(totales) {
			totales = append(totales, Total{ProductoID: l.ProductoID, Producto: l.Producto})
		}
		totales[i].Cantidad += l.Cantidad
		totales[i].Monto += l.Subtotal
	}
	for i := 1; i < len(totales); i++ {
		t := totales[i]
		j := i - 1
		for j >= 0 && totales[j].Cantidad < t.Cantidad {
			totales[j+1] = totales[j]
			j--
		}
		totales[j+1] = t
	}
	return totales
}

// mergeSort estable: antes(a, b) dice si a debe ir antes que b.
func mergeSort[T any](datos []T, antes func(a, b T) bool) {
	if len(datos) < 2 {
		return
	}
	izq := append([]T(nil), datos[:len(datos)/2]...)
	der := append([]T(nil), datos[len(datos)/2:]...)
	mergeSort(izq, antes)
	mergeSort(der, antes)
	i, j := 0, 0
	for k := range datos {
		if j == len(der) || (i < len(izq) && !antes(der[j], izq[i])) {
			datos[k] = izq[i]
			i++
		} else {
			datos[k] = der[j]
			j++
		}
	}
}

// buscarFecha: búsqueda binaria de la primera línea cuya fecha cumple cond.
func buscarFecha(lineas []Linea, cond func(fecha string) bool) int {
	izq, der := 0, len(lineas)
	for izq < der {
		medio := (izq + der) / 2
		if cond(lineas[medio].Fecha) {
			der = medio
		} else {
			izq = medio + 1
		}
	}
	return izq
}

// reporteMejorado: binaria para el rango, mapa para sumar y MergeSort para ordenar.
// porFecha son las líneas ya ordenadas por fecha (se ordenan una vez al cargar).
func reporteMejorado(porFecha []Linea, sucursalID int, desde, hasta string) []Total {
	inicio := buscarFecha(porFecha, func(f string) bool { return f >= desde })
	fin := buscarFecha(porFecha, func(f string) bool { return f > hasta })

	posicion := map[int]int{}
	var totales []Total
	for _, l := range porFecha[inicio:fin] {
		if sucursalID != 0 && l.SucursalID != sucursalID {
			continue
		}
		i, existe := posicion[l.ProductoID]
		if !existe {
			i = len(totales)
			posicion[l.ProductoID] = i
			totales = append(totales, Total{ProductoID: l.ProductoID, Producto: l.Producto, Primera: l.Pos})
		}
		totales[i].Cantidad += l.Cantidad
		totales[i].Monto += l.Subtotal
		totales[i].Primera = min(totales[i].Primera, l.Pos)
	}
	// Empates: primero el producto que apareció antes, igual que la inserción de hoy.
	mergeSort(totales, func(a, b Total) bool {
		if a.Cantidad != b.Cantidad {
			return a.Cantidad > b.Cantidad
		}
		return a.Primera < b.Primera
	})
	return totales
}

func iguales(a, b []Total) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].ProductoID != b[i].ProductoID || a[i].Cantidad != b[i].Cantidad || a[i].Monto != b[i].Monto {
			return false
		}
	}
	return true
}

// generar: n líneas, 8 sucursales, 8.000 productos, semilla 42 (como el H1).
// Con unAnio = false todas son de octubre; con true se reparten en los 12 meses.
func generar(n int, unAnio bool) []Linea {
	r := rand.New(rand.NewSource(42))
	lineas := make([]Linea, n)
	for i := range lineas {
		p := r.Intn(8000) + 1
		cant := r.Intn(5) + 1
		fecha := fmt.Sprintf("2026-10-%02d", r.Intn(31)+1)
		if unAnio {
			fecha = fmt.Sprintf("2026-%02d-%02d", r.Intn(12)+1, r.Intn(28)+1)
		}
		lineas[i] = Linea{fecha, r.Intn(8) + 1, p, "Producto " + strconv.Itoa(p), cant, float64(cant * (p%200 + 5)), i}
	}
	return lineas
}

// medir: 1 calentamiento y 3 corridas, en ms.
func medir(f func()) [3]float64 {
	var t [3]float64
	f()
	for i := range t {
		inicio := time.Now()
		f()
		t[i] = time.Since(inicio).Seconds() * 1000
	}
	return t
}

type caso struct {
	nombre       string
	lineas       []Linea
	desde, hasta string
}

func main() {
	fmt.Println("Gonzalo Yupanqui Torrez")
	anio := generar(1_440_000, true)
	casos := []caso{
		{"Un mes, 30.000 líneas", generar(30_000, false), "2026-10-01", "2026-10-31"},
		{"Un mes, 60.000 líneas", generar(60_000, false), "2026-10-01", "2026-10-31"},
		{"Un mes, 120.000 líneas", generar(120_000, false), "2026-10-01", "2026-10-31"},
		{"Un mes, 240.000 líneas", generar(240_000, false), "2026-10-01", "2026-10-31"},
		{"Un año, reporte de octubre", anio, "2026-10-01", "2026-10-31"},
		{"Un año, reporte del año", anio, "2026-01-01", "2026-12-31"},
	}

	var antes, despues [][3]float64
	var cargar []float64
	for _, c := range casos {
		inicio := time.Now()
		porFecha := append([]Linea(nil), c.lineas...)
		mergeSort(porFecha, func(a, b Linea) bool { return a.Fecha < b.Fecha })
		cargar = append(cargar, time.Since(inicio).Seconds()*1000)

		for _, suc := range []int{0, 1} {
			if !iguales(reporteActual(c.lineas, suc, c.desde, c.hasta), reporteMejorado(porFecha, suc, c.desde, c.hasta)) {
				panic("las dos versiones no dan el mismo reporte: " + c.nombre)
			}
		}
		antes = append(antes, medir(func() { reporteActual(c.lineas, 0, c.desde, c.hasta) }))
		despues = append(despues, medir(func() { reporteMejorado(porFecha, 0, c.desde, c.hasta) }))
	}
	fmt.Println("Mismo reporte en todos los casos (todas las sucursales y sucursal 1): sí")

	tabla("Antes", casos, antes)
	tabla("Después", casos, despues)

	fmt.Print("\n## Resumen\n\n")
	fmt.Println("| Datos | Antes (ms) | Después (ms) | Veces más rápido | Ordenar por fecha al cargar (ms) |")
	fmt.Println("| --- | ---: | ---: | ---: | ---: |")
	for i, c := range casos {
		a, d := prom(antes[i]), prom(despues[i])
		fmt.Printf("| %s | %.2f | %.2f | x%.0f | %.2f |\n", c.nombre, a, d, a/d, cargar[i])
	}
}

func prom(t [3]float64) float64 { return (t[0] + t[1] + t[2]) / 3 }

func tabla(titulo string, casos []caso, t [][3]float64) {
	fmt.Printf("\n## %s\n\n", titulo)
	fmt.Println("| Datos | Corrida 1 (ms) | Corrida 2 (ms) | Corrida 3 (ms) | Promedio (ms) | Veces vs fila anterior |")
	fmt.Println("| --- | ---: | ---: | ---: | ---: | ---: |")
	for i, c := range casos {
		veces := "—"
		if i > 0 && i < 4 {
			veces = fmt.Sprintf("x%.2f", prom(t[i])/prom(t[i-1]))
		}
		fmt.Printf("| %s | %.2f | %.2f | %.2f | %.2f | %s |\n", c.nombre, t[i][0], t[i][1], t[i][2], prom(t[i]), veces)
	}
}
