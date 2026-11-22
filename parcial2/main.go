// Gonzalo Yupanqui Torrez
package main

import (
	"fmt"
	"strconv"
	"time"
)

type Paciente struct {
	nombre   string
	urgencia int
}

// ---------- Solución ingenua: una sola lista ----------

type Ingenuo struct {
	lista []Paciente
}

func registrarIngenuo(g Ingenuo, nombre string, urgencia int) Ingenuo {
	g.lista = append(g.lista, Paciente{nombre, urgencia})
	return g
}

func atenderIngenuo(g Ingenuo) (Ingenuo, string) {
	if len(g.lista) == 0 {
		return g, "(nadie)"
	}
	mejor := 0
	for i := 1; i < len(g.lista); i++ {
		if g.lista[i].urgencia < g.lista[mejor].urgencia {
			mejor = i
		}
	}
	nombre := g.lista[mejor].nombre
	g.lista = append(g.lista[:mejor], g.lista[mejor+1:]...)
	return g, nombre
}

func estaIngenuo(g Ingenuo, nombre string) bool {
	for i := 0; i < len(g.lista); i++ {
		if g.lista[i].nombre == nombre {
			return true
		}
	}
	return false
}

func esperandoIngenuo(g Ingenuo) int {
	return len(g.lista)
}

// ---------- Solución mejorada: una cola por urgencia + un mapa ----------

type Mejorado struct {
	colas    [6][]string
	enEspera map[string]bool
	cantidad int
}

func nuevoMejorado() Mejorado {
	var g Mejorado
	g.enEspera = map[string]bool{}
	return g
}

func registrarMejorado(g Mejorado, nombre string, urgencia int) Mejorado {
	g.colas[urgencia] = append(g.colas[urgencia], nombre)
	g.enEspera[nombre] = true
	g.cantidad++
	return g
}

func atenderMejorado(g Mejorado) (Mejorado, string) {
	for u := 1; u <= 5; u++ {
		if len(g.colas[u]) > 0 {
			nombre := g.colas[u][0]
			g.colas[u] = g.colas[u][1:]
			delete(g.enEspera, nombre)
			g.cantidad--
			return g, nombre
		}
	}
	return g, "(nadie)"
}

func estaMejorado(g Mejorado, nombre string) bool {
	return g.enEspera[nombre]
}

func esperandoMejorado(g Mejorado) int {
	return g.cantidad
}

// ---------- Guion de pruebas ----------

type Paso struct {
	op       string
	arg      string
	urg      int
	esperado string
}

var guion = []Paso{
	{"registrar", "Ana", 3, ""}, {"registrar", "Beto", 1, ""}, {"registrar", "Carla", 3, ""}, {"registrar", "Diego", 2, ""},
	{"esperando", "", 0, "4"},
	{"esta", "Beto", 0, "true"}, {"esta", "Zoe", 0, "false"},
	{"atender", "", 0, "Beto"}, {"esta", "Beto", 0, "false"},
	{"atender", "", 0, "Diego"},
	{"registrar", "Elena", 1, ""},
	{"atender", "", 0, "Elena"}, {"atender", "", 0, "Ana"}, {"atender", "", 0, "Carla"},
	{"atender", "", 0, "(nadie)"}, {"esperando", "", 0, "0"},
}

func revisar(p Paso, obtenido string, bien int) int {
	if obtenido == p.esperado {
		return bien + 1
	}
	fmt.Println("   MAL en", p.op, p.arg, "esperado", p.esperado, "obtuvo", obtenido)
	return bien
}

func probarIngenuo() {
	g := Ingenuo{}
	bien, total := 0, 0
	for _, p := range guion {
		obtenido := ""
		if p.op == "registrar" {
			g = registrarIngenuo(g, p.arg, p.urg)
			continue
		} else if p.op == "atender" {
			g, obtenido = atenderIngenuo(g)
		} else if p.op == "esta" {
			obtenido = strconv.FormatBool(estaIngenuo(g, p.arg))
		} else {
			obtenido = strconv.Itoa(esperandoIngenuo(g))
		}
		total++
		bien = revisar(p, obtenido, bien)
	}
	fmt.Println("Ingenuo: ", bien, "de", total, "pasos correctos")
}

func probarMejorado() {
	g := nuevoMejorado()
	bien, total := 0, 0
	for _, p := range guion {
		obtenido := ""
		if p.op == "registrar" {
			g = registrarMejorado(g, p.arg, p.urg)
			continue
		} else if p.op == "atender" {
			g, obtenido = atenderMejorado(g)
		} else if p.op == "esta" {
			obtenido = strconv.FormatBool(estaMejorado(g, p.arg))
		} else {
			obtenido = strconv.Itoa(esperandoMejorado(g))
		}
		total++
		bien = revisar(p, obtenido, bien)
	}
	fmt.Println("Mejorado:", bien, "de", total, "pasos correctos")
}

// ---------- Carga: n pacientes llegan y después se atiende a los n ----------

func urgenciaDe(i int) int {
	return (i*7919)%5 + 1
}

func cargaIngenuo(n int) {
	g := Ingenuo{}
	for i := 0; i < n; i++ {
		g = registrarIngenuo(g, "P"+strconv.Itoa(i), urgenciaDe(i))
	}
	for i := 0; i < n; i++ {
		g, _ = atenderIngenuo(g)
	}
}

func cargaMejorado(n int) {
	g := nuevoMejorado()
	for i := 0; i < n; i++ {
		g = registrarMejorado(g, "P"+strconv.Itoa(i), urgenciaDe(i))
	}
	for i := 0; i < n; i++ {
		g, _ = atenderMejorado(g)
	}
}

// ---------- Medición: 1 calentamiento + 3 corridas + promedio ----------

func medir(cual string, n int) float64 {
	suma := 0.0
	for corrida := 0; corrida <= 3; corrida++ {
		inicio := time.Now()
		if cual == "ingenuo" {
			cargaIngenuo(n)
		} else {
			cargaMejorado(n)
		}
		ms := float64(time.Since(inicio).Microseconds()) / 1000
		if corrida == 0 {
			fmt.Printf("   n=%d calentamiento: %.2f ms\n", n, ms)
		} else {
			fmt.Printf("   n=%d corrida %d: %.2f ms\n", n, corrida, ms)
			suma += ms
		}
	}
	return suma / 3
}

func tabla(cual string) {
	fmt.Println("\nMedición", cual)
	tamanos := []int{10000, 20000, 40000}
	anterior := 0.0
	for _, n := range tamanos {
		promedio := medir(cual, n)
		if anterior == 0 {
			fmt.Printf("   n=%d promedio %.2f ms\n", n, promedio)
		} else {
			fmt.Printf("   n=%d promedio %.2f ms (x%.2f)\n", n, promedio, promedio/anterior)
		}
		anterior = promedio
	}
}

func main() {
	probarIngenuo()
	probarMejorado()
	tabla("ingenuo")
	tabla("mejorado")
}
