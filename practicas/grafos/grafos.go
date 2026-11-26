package grafos

import (
	"fmt"
	"math/rand"
	"time"
)

// generar arma un grafo conexo: cada nodo se une al siguiente y a 3 nodos al azar.
func generar(n int) [][]int {
	r := rand.New(rand.NewSource(int64(n)))
	ady := make([][]int, n)
	unir := func(a, b int) {
		ady[a] = append(ady[a], b)
		ady[b] = append(ady[b], a)
	}
	for v := 0; v < n; v++ {
		if v+1 < n {
			unir(v, v+1)
		}
		for k := 0; k < 3; k++ {
			unir(v, r.Intn(n))
		}
	}
	return ady
}

// bfs recorre por niveles con una cola; devuelve cuántos nodos visitó.
func bfs(ady [][]int, inicio int) int {
	visto := make([]bool, len(ady))
	cola := []int{inicio}
	visto[inicio] = true
	for frente := 0; frente < len(cola); frente++ {
		for _, w := range ady[cola[frente]] {
			if !visto[w] {
				visto[w] = true
				cola = append(cola, w)
			}
		}
	}
	return len(cola)
}

// dfs va lo más hondo posible con una pila; devuelve cuántos nodos visitó.
func dfs(ady [][]int, inicio int) int {
	visto := make([]bool, len(ady))
	pila := []int{inicio}
	visitados := 0
	for len(pila) > 0 {
		v := pila[len(pila)-1]
		pila = pila[:len(pila)-1]
		if visto[v] {
			continue
		}
		visto[v] = true
		visitados++
		for _, w := range ady[v] {
			if !visto[w] {
				pila = append(pila, w)
			}
		}
	}
	return visitados
}

func medir(recorrer func([][]int, int) int, ady [][]int) [3]float64 {
	var t [3]float64
	for corrida := 0; corrida <= 3; corrida++ {
		inicio := time.Now()
		if recorrer(ady, 0) != len(ady) {
			panic("el recorrido no visitó todos los nodos")
		}
		if corrida > 0 {
			t[corrida-1] = time.Since(inicio).Seconds() * 1000
		}
	}
	return t
}

func EjecutarGrafos() {
	fmt.Println("Grafos: BFS y DFS, O(V + E) (1 calentamiento, 3 corridas, promedio en ms)")
	fmt.Printf("%-5s %-9s %-9s %10s %10s %10s %10s %8s\n", "", "Nodos", "Aristas", "Corrida 1", "Corrida 2", "Corrida 3", "Promedio", "Veces")
	anterior := map[string]float64{}
	for _, n := range []int{100_000, 200_000, 400_000, 800_000} {
		ady := generar(n)
		aristas := 0
		for _, vecinos := range ady {
			aristas += len(vecinos)
		}
		for _, alg := range []struct {
			nombre   string
			recorrer func([][]int, int) int
		}{{"BFS", bfs}, {"DFS", dfs}} {
			t := medir(alg.recorrer, ady)
			prom := (t[0] + t[1] + t[2]) / 3
			veces := "—"
			if anterior[alg.nombre] > 0 {
				veces = fmt.Sprintf("x%.2f", prom/anterior[alg.nombre])
			}
			fmt.Printf("%-5s %-9d %-9d %10.3f %10.3f %10.3f %10.3f %8s\n", alg.nombre, n, aristas/2, t[0], t[1], t[2], prom, veces)
			anterior[alg.nombre] = prom
		}
	}
}
