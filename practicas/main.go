package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"practicas/c3"
	"practicas/c4"
	"practicas/c5"
)

func main() {
	entrada := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n=== Prácticas de Análisis de Algoritmos ===")
		fmt.Println("1. C3 - Búsqueda binaria")
		fmt.Println("2. C4 - Comparación medida: lineal contra binaria")
		fmt.Println("3. C5 - Torneo de ordenamientos")
		fmt.Println("0. Salir")
		fmt.Print("Elige un ejercicio: ")

		if !entrada.Scan() {
			fmt.Println()
			if err := entrada.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "No se pudo leer la opción:", err)
				os.Exit(1)
			}
			return
		}

		switch strings.TrimSpace(entrada.Text()) {
		case "1":
			fmt.Println()
			c3.EjecutarBusquedaBinaria()
		case "2":
			fmt.Println()
			c4.EjecutarComparacion()
		case "3":
			fmt.Println()
			c5.EjecutarTorneo()
		case "0":
			fmt.Println("Hasta luego.")
			return
		default:
			fmt.Println("Opción inválida. Elige una opción del menú.")
		}
	}
}
