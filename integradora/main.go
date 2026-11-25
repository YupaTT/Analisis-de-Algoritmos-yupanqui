// Gonzalo Yupanqui Torrez
package main

import (
	"fmt"
	"math/rand"
	"slices"
	"time"
)

// ---------- MODELO ----------

type Util struct {
	Codigo   string
	CostoBs  int
	Utilidad int
}

type Comprador interface {
	Elegir([]Util, int) (int, []string)
	EstaEnLaCompra(string) bool
}

// ---------- ALGORITMO INGENUO ----------

type Ingenuo struct {
	elegidos []string
}

func (c *Ingenuo) Elegir(utiles []Util, presupuesto int) (int, []string) {
	n := len(utiles)
	mejor, mascaraMejor := 0, 0

	for mascara := 0; mascara < 1<<n; mascara++ {
		costo, utilidad := 0, 0

		for i, u := range utiles {
			if mascara&(1<<i) != 0 {
				costo += u.CostoBs
				utilidad += u.Utilidad
			}
		}

		if costo <= presupuesto && utilidad > mejor {
			mejor = utilidad
			mascaraMejor = mascara
		}
	}

	c.elegidos = []string{}

	for i, u := range utiles {
		if mascaraMejor&(1<<i) != 0 {
			c.elegidos = append(c.elegidos, u.Codigo)
		}
	}

	return mejor, c.elegidos
}

func (c *Ingenuo) EstaEnLaCompra(codigo string) bool {
	return slices.Contains(c.elegidos, codigo)
}

// ---------- ALGORITMO MEJORADO ----------

type Mejorado struct {
	enCompra map[string]bool
}

func (c *Mejorado) Elegir(utiles []Util, presupuesto int) (int, []string) {
	n := len(utiles)
	c.enCompra = make(map[string]bool)

	if presupuesto < 0 || n == 0 {
		return 0, nil
	}

	// dp[i][p] = utilidad máxima usando los primeros i útiles
	// con un presupuesto de p bolivianos.
	dp := make([][]int, n+1)
	tomar := make([][]bool, n+1)

	for i := range dp {
		dp[i] = make([]int, presupuesto+1)
		tomar[i] = make([]bool, presupuesto+1)
	}

	// Construir la tabla de programación dinámica.
	for i, u := range utiles {
		for p := 0; p <= presupuesto; p++ {
			// Opción 1: no comprar el útil.
			dp[i+1][p] = dp[i][p]

			// Opción 2: comprarlo si alcanza el presupuesto.
			if u.CostoBs <= p {
				conUtil := dp[i][p-u.CostoBs] + u.Utilidad

				if conUtil > dp[i+1][p] {
					dp[i+1][p] = conUtil
					tomar[i+1][p] = true
				}
			}
		}
	}

	// Recuperar los útiles seleccionados.
	p := presupuesto
	elegidos := []string{}

	for i := n; i > 0; i-- {
		if tomar[i][p] {
			u := utiles[i-1]

			elegidos = append(elegidos, u.Codigo)
			c.enCompra[u.Codigo] = true
			p -= u.CostoBs
		}
	}

	return dp[n][presupuesto], elegidos
}

func (c *Mejorado) EstaEnLaCompra(codigo string) bool {
	return c.enCompra[codigo]
}

// ---------- CASOS DE PRUEBA ----------

type Caso struct {
	nombre      string
	utiles      []Util
	presupuesto int
	esperada    int
	codigos     []string
}

var casos = []Caso{
	{"el voraz se equivoca",
		[]Util{{"U1", 20, 120}, {"U2", 40, 200}, {"U3", 60, 240}},
		100, 440, []string{"U2", "U3"}},

	{"cuatro útiles",
		[]Util{{"U1", 4, 6}, {"U2", 6, 8}, {"U3", 8, 10}, {"U4", 10, 12}},
		16, 20, []string{"U2", "U4"}},

	{"presupuesto cero",
		[]Util{{"U1", 5, 10}, {"U2", 7, 12}},
		0, 0, []string{}},

	{"sin útiles",
		[]Util{}, 30, 0, []string{}},

	{"nada cabe",
		[]Util{{"U1", 80, 90}, {"U2", 100, 100}},
		60, 0, []string{}},

	{"entran todos",
		[]Util{{"U1", 10, 20}, {"U2", 14, 24}, {"U3", 6, 8}},
		200, 52, []string{"U1", "U2", "U3"}},
}

// Comparar códigos sin importar su orden.
func iguales(a, b []string) bool {
	x, y := slices.Clone(a), slices.Clone(b)
	slices.Sort(x)
	slices.Sort(y)
	return slices.Equal(x, y)
}

func probar(nombre string, nuevo func() Comprador) {
	bien := 0

	for _, caso := range casos {
		c := nuevo()
		utilidad, codigos := c.Elegir(caso.utiles, caso.presupuesto)

		if utilidad == caso.esperada && iguales(codigos, caso.codigos) {
			bien++
		} else {
			fmt.Printf("MAL en %s\n", caso.nombre)
		}
	}

	c := nuevo()
	c.Elegir(casos[0].utiles, casos[0].presupuesto)

	if c.EstaEnLaCompra("U2") {
		bien++
	}

	if !c.EstaEnLaCompra("U1") {
		bien++
	}

	fmt.Printf("%s: %d de 8 pruebas correctas\n", nombre, bien)
}

// ---------- GENERAR DATOS ----------

func carga(n int) ([]Util, int) {
	azar := rand.New(rand.NewSource(int64(n + 1000)))
	utiles := make([]Util, n)
	total := 0

	for i := range utiles {
		costo := 4 + azar.Intn(46)
		utiles[i] = Util{
			Codigo:   fmt.Sprintf("K%d", i),
			CostoBs:  costo,
			Utilidad: 10 + azar.Intn(90),
		}
		total += costo
	}

	return utiles, total / 3
}

// ---------- MEDIR RENDIMIENTO ----------

func medir(nombre string, nuevo func() Comprador, utiles []Util, presupuesto int) int {
	nuevo().Elegir(utiles, presupuesto) // Calentamiento.

	var total time.Duration
	utilidad := 0

	fmt.Printf("%-9s n=%-4d corridas:", nombre, len(utiles))

	for i := 0; i < 3; i++ {
		inicio := time.Now()
		utilidad, _ = nuevo().Elegir(utiles, presupuesto)
		t := time.Since(inicio)
		total += t
		fmt.Printf(" %10.3f", float64(t.Microseconds())/1000)
	}

	promedio := total / 3

	fmt.Printf("  promedio: %10.3f ms utilidad: %d\n",
		float64(promedio.Microseconds())/1000, utilidad)

	return utilidad
}

// ---------- PROGRAMA PRINCIPAL ----------

func main() {
	ingenuo := func() Comprador { return &Ingenuo{} }
	mejorado := func() Comprador { return &Mejorado{} }

	fmt.Println("=== PRUEBAS ===")
	probar("Ingenuo", ingenuo)
	probar("Mejorado", mejorado)

	fmt.Println("\n=== COMPARACIÓN DE RENDIMIENTO ===")

	for _, n := range []int{20, 22, 24, 26} {
		utiles, presupuesto := carga(n)

		ui := medir("Ingenuo", ingenuo, utiles, presupuesto)
		um := medir("Mejorado", mejorado, utiles, presupuesto)

		fmt.Printf("Misma utilidad: %v\n\n", ui == um)
	}

	fmt.Println("=== PRUEBA DE ESCALABILIDAD ===")

	utiles, _ := carga(1000)
	medir("Mejorado", mejorado, utiles, 10000)
}
