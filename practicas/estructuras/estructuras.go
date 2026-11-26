package estructuras

import (
	"fmt"
	"math/rand"
	"time"
)

type nodo struct {
	clave    int
	izq, der *nodo
}

func insertarBST(raiz *nodo, clave int) *nodo {
	if raiz == nil {
		return &nodo{clave: clave}
	}
	if clave < raiz.clave {
		raiz.izq = insertarBST(raiz.izq, clave)
	} else {
		raiz.der = insertarBST(raiz.der, clave)
	}
	return raiz
}

func buscarBST(raiz *nodo, clave int) bool {
	for raiz != nil {
		if clave == raiz.clave {
			return true
		}
		if clave < raiz.clave {
			raiz = raiz.izq
		} else {
			raiz = raiz.der
		}
	}
	return false
}

func buscarLista(lista []int, clave int) bool {
	for _, v := range lista {
		if v == clave {
			return true
		}
	}
	return false
}

// sacarMayorLista recorre toda la lista para hallar el mayor y lo quita: O(n).
func sacarMayorLista(lista []int) []int {
	pos := 0
	for i, v := range lista {
		if v > lista[pos] {
			pos = i
		}
	}
	lista[pos] = lista[len(lista)-1]
	return lista[:len(lista)-1]
}

// sacarMayorHeap quita la raíz del montículo de máximos y lo repara: O(log n).
func sacarMayorHeap(h []int) []int {
	n := len(h) - 1
	h[0] = h[n]
	h = h[:n]
	i := 0
	for {
		mayor, izq, der := i, 2*i+1, 2*i+2
		if izq < n && h[izq] > h[mayor] {
			mayor = izq
		}
		if der < n && h[der] > h[mayor] {
			mayor = der
		}
		if mayor == i {
			return h
		}
		h[i], h[mayor] = h[mayor], h[i]
		i = mayor
	}
}

func armarHeap(datos []int) []int {
	h := append([]int(nil), datos...)
	for i := 1; i < len(h); i++ {
		for j := i; j > 0 && h[(j-1)/2] < h[j]; j = (j - 1) / 2 {
			h[j], h[(j-1)/2] = h[(j-1)/2], h[j]
		}
	}
	return h
}

func pilaMeterSacar(n int) {
	var pila []int
	for i := 0; i < n; i++ {
		pila = append(pila, i)
	}
	for len(pila) > 0 {
		pila = pila[:len(pila)-1]
	}
}

func colaMeterSacar(n int) {
	var cola []int
	for i := 0; i < n; i++ {
		cola = append(cola, i)
	}
	for frente := 0; frente < len(cola); frente++ {
		_ = cola[frente]
	}
}

// medir: 1 calentamiento y 3 corridas. preparar arma datos nuevos fuera del tiempo.
func medir(preparar func(), operar func()) [3]float64 {
	var t [3]float64
	for corrida := 0; corrida <= 3; corrida++ {
		preparar()
		inicio := time.Now()
		operar()
		if corrida > 0 {
			t[corrida-1] = time.Since(inicio).Seconds() * 1000
		}
	}
	return t
}

func fila(nombre string, n int, t [3]float64, anterior map[string]float64) {
	prom := (t[0] + t[1] + t[2]) / 3
	veces := "—"
	if anterior[nombre] > 0 {
		veces = fmt.Sprintf("x%.2f", prom/anterior[nombre])
	}
	fmt.Printf("%-8s %-8d %10.3f %10.3f %10.3f %10.3f %8s\n", nombre, n, t[0], t[1], t[2], prom, veces)
	anterior[nombre] = prom
}

func encabezado(titulo string) {
	fmt.Printf("\n%s\n", titulo)
	fmt.Printf("%-8s %-8s %10s %10s %10s %10s %8s\n", "Versión", "n", "Corrida 1", "Corrida 2", "Corrida 3", "Promedio", "Veces")
}

func EjecutarEstructuras() {
	const consultas = 1_000
	tamanos := []int{10_000, 20_000, 40_000, 80_000}
	fmt.Println("Estructuras de datos (1 calentamiento, 3 corridas, promedio en ms)")

	encabezado("1) Meter y sacar n elementos: pila y cola, O(1) cada operación")
	anterior := map[string]float64{}
	for _, n := range []int{1_000_000, 2_000_000, 4_000_000, 8_000_000} {
		fila("Pila", n, medir(func() {}, func() { pilaMeterSacar(n) }), anterior)
		fila("Cola", n, medir(func() {}, func() { colaMeterSacar(n) }), anterior)
	}

	encabezado(fmt.Sprintf("2) Buscar %d claves: lista O(n), BST O(log n), hash O(1)", consultas))
	anterior = map[string]float64{}
	for _, n := range tamanos {
		r := rand.New(rand.NewSource(int64(n)))
		lista := r.Perm(n)
		var raiz *nodo
		hash := make(map[int]bool, n)
		for _, v := range lista {
			raiz = insertarBST(raiz, v)
			hash[v] = true
		}
		claves := make([]int, consultas)
		for i := range claves {
			claves[i] = r.Intn(n)
			if !buscarLista(lista, claves[i]) || !buscarBST(raiz, claves[i]) || !hash[claves[i]] {
				panic("las estructuras no coinciden")
			}
		}
		fila("Lista", n, medir(func() {}, func() {
			for _, c := range claves {
				buscarLista(lista, c)
			}
		}), anterior)
		fila("BST", n, medir(func() {}, func() {
			for _, c := range claves {
				buscarBST(raiz, c)
			}
		}), anterior)
		fila("Hash", n, medir(func() {}, func() {
			for _, c := range claves {
				_ = hash[c]
			}
		}), anterior)
	}

	encabezado(fmt.Sprintf("3) Sacar el mayor %d veces (top-k): lista O(n), heap O(log n)", consultas))
	anterior = map[string]float64{}
	for _, n := range tamanos {
		datos := rand.New(rand.NewSource(int64(n))).Perm(n)
		var lista, heap []int
		fila("Lista", n, medir(func() { lista = append([]int(nil), datos...) }, func() {
			for i := 0; i < consultas; i++ {
				lista = sacarMayorLista(lista)
			}
		}), anterior)
		fila("Heap", n, medir(func() { heap = armarHeap(datos) }, func() {
			for i := 0; i < consultas; i++ {
				heap = sacarMayorHeap(heap)
			}
		}), anterior)
		if heap[0] != n-1-consultas || len(lista) != len(heap) {
			panic("lista y heap no sacaron los mismos mayores")
		}
	}
}
