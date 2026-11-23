package main

import (
	"fmt"
	"slices"
	"time"
)

type Producto struct {
	ID        int
	Categoria int
	Vendidos  int
}

const (
	categorias = 6
	consultas  = 10_000
	top        = 5
)

func generar(n int) []Producto {
	lista := make([]Producto, n)
	for i := range lista {
		id := i + 1
		lista[i] = Producto{ID: id, Categoria: id%categorias + 1, Vendidos: (id * 37) % 500}
	}
	return lista
}

func masVendidoPrimero(a, b Producto) int {
	if a.Vendidos != b.Vendidos {
		return b.Vendidos - a.Vendidos
	}
	return a.ID - b.ID
}

func topIngenua(lista []Producto, categoria int) []Producto {
	var encontrados []Producto
	for _, p := range lista {
		if p.Categoria == categoria {
			encontrados = append(encontrados, p)
		}
	}
	slices.SortFunc(encontrados, masVendidoPrimero)
	return encontrados[:min(top, len(encontrados))]
}

func construirMapa(lista []Producto) map[int][]Producto {
	mapa := make(map[int][]Producto, categorias)
	for _, p := range lista {
		mapa[p.Categoria] = append(mapa[p.Categoria], p)
	}
	for _, productos := range mapa {
		slices.SortFunc(productos, masVendidoPrimero)
	}
	return mapa
}

func topMapa(mapa map[int][]Producto, categoria int) []Producto {
	productos := mapa[categoria]
	return productos[:min(top, len(productos))]
}

func medir(consultar func(categoria int) []Producto) float64 {
	total := 0.0
	for corrida := 0; corrida <= 3; corrida++ {
		inicio := time.Now()
		for q := 0; q < consultas; q++ {
			consultar(q%categorias + 1)
		}
		if corrida > 0 {
			total += time.Since(inicio).Seconds() * 1000
		}
	}
	return total / 3
}

func main() {
	fmt.Printf("%d consultas por tamaño; 1 calentamiento y promedio de 3 corridas.\n\n", consultas)
	fmt.Println("| n | Ingenua (ms) | Veces más | Elegida (ms) | Veces más |")
	fmt.Println("| ---: | ---: | ---: | ---: | ---: |")

	var antIngenua, antElegida float64
	for _, n := range []int{1_000, 2_000, 4_000, 8_000} {
		lista := generar(n)
		mapa := construirMapa(lista)
		for c := 1; c <= categorias; c++ {
			if !slices.Equal(topIngenua(lista, c), topMapa(mapa, c)) {
				panic(fmt.Sprintf("resultados distintos en n=%d, categoría %d", n, c))
			}
		}

		ingenua := medir(func(c int) []Producto { return topIngenua(lista, c) })
		elegida := medir(func(c int) []Producto { return topMapa(mapa, c) })

		vecesI, vecesE := "—", "—"
		if antIngenua > 0 {
			vecesI = fmt.Sprintf("%.2f", ingenua/antIngenua)
			vecesE = fmt.Sprintf("%.2f", elegida/antElegida)
		}
		fmt.Printf("| %d | %.3f | %s | %.4f | %s |\n", n, ingenua, vecesI, elegida, vecesE)
		antIngenua, antElegida = ingenua, elegida
	}

	lista := generar(1_000)
	fmt.Println("\nCategoría 4, 5 más vendidos (n = 1000):")
	for _, p := range topMapa(construirMapa(lista), 4) {
		fmt.Printf("  producto %d, vendidos %d\n", p.ID, p.Vendidos)
	}
}
