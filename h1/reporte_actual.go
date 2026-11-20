// Gonzalo Yupanqui Torrez
//
// H1 — Simulación del proceso ACTUAL (lento) del reporte de ventas por rango de fechas
// y, opcionalmente, por sucursal.
// Reproduce la lógica tal como está hoy, sin optimizar, para medir su costo.
//
//	go run h1/reporte_actual.go h1/ventas_muestra.csv     -> reportes con la muestra
//	go run h1/reporte_actual.go -n 120000                 -> mide con datos generados
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Linea struct {
	Fecha      string // AAAA-MM-DD: como texto se compara igual que la fecha
	VentaID    int
	SucursalID int
	ProductoID int
	Producto   string
	Cantidad   int
	Subtotal   float64
}

type Total struct {
	ProductoID int
	Producto   string
	Cantidad   int
	Monto      float64
}

// reporte es la lógica actual: tres pasos, ninguno usa estructura.
// sucursalID = 0 significa todas las sucursales.
func reporte(lineas []Linea, sucursalID int, desde, hasta string) []Total {
	// 1. Filtrar: recorre TODAS las líneas para quedarse con el rango pedido. O(n)
	var filtradas []Linea
	for _, l := range lineas {
		if (sucursalID == 0 || l.SucursalID == sucursalID) && l.Fecha >= desde && l.Fecha <= hasta {
			filtradas = append(filtradas, l)
		}
	}

	// 2. Acumular: por cada línea busca su producto uno por uno. O(m·p)
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

	// 3. Ordenar de más a menos vendido con inserción. O(p²)
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

func leerCSV(ruta string) []Linea {
	f, err := os.Open(ruta)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer f.Close()
	filas, err := csv.NewReader(f).ReadAll()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	var lineas []Linea
	for _, c := range filas[1:] { // fecha,ventaId,sucursalId,productoId,producto,cantidad,precioUnitario,subtotal
		venta, _ := strconv.Atoi(c[1])
		suc, _ := strconv.Atoi(c[2])
		prod, _ := strconv.Atoi(c[3])
		cant, _ := strconv.Atoi(c[5])
		sub, _ := strconv.ParseFloat(c[7], 64)
		lineas = append(lineas, Linea{c[0], venta, suc, prod, c[4], cant, sub})
	}
	return lineas
}

// generar crea n líneas de un mes con 8 sucursales y 8.000 productos (semilla fija).
func generar(n int) []Linea {
	r := rand.New(rand.NewSource(42))
	lineas := make([]Linea, n)
	for i := range lineas {
		p := r.Intn(8000) + 1
		cant := r.Intn(5) + 1
		lineas[i] = Linea{
			Fecha:      fmt.Sprintf("2026-10-%02d", r.Intn(31)+1),
			VentaID:    i/3 + 1,
			SucursalID: r.Intn(8) + 1,
			ProductoID: p,
			Producto:   "Producto " + strconv.Itoa(p),
			Cantidad:   cant,
			Subtotal:   float64(cant * (p%200 + 5)),
		}
	}
	return lineas
}

func main() {
	n := flag.Int("n", 0, "líneas generadas para medir (0 = leer CSV)")
	flag.Parse()

	if *n == 0 {
		ruta := "h1/ventas_muestra.csv"
		if flag.NArg() > 0 {
			ruta = flag.Arg(0)
		}
		lineas := leerCSV(ruta)
		for _, suc := range []int{0, 1} {
			fmt.Printf("Del 2026-11-03 al 2026-11-05, sucursal %d (0 = todas):\n", suc)
			for i, t := range reporte(lineas, suc, "2026-11-03", "2026-11-05") {
				fmt.Printf("%2d. %-14s %4d u.  Bs %9.2f\n", i+1, t.Producto, t.Cantidad, t.Monto)
			}
		}
		return
	}

	// Protocolo del curso: un calentamiento, 3 corridas, promedio.
	lineas := generar(*n)
	reporte(lineas, 0, "2026-10-01", "2026-10-31")
	var suma time.Duration
	for i := 1; i <= 3; i++ {
		ini := time.Now()
		reporte(lineas, 0, "2026-10-01", "2026-10-31")
		d := time.Since(ini)
		suma += d
		fmt.Printf("corrida %d: %v\n", i, d)
	}
	fmt.Printf("promedio con %d líneas: %v\n", *n, suma/3)
}
