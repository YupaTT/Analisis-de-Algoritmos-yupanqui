Gonzalo Yupanqui Torrez

# P1.1 — Análisis · Variante B

## a) Cuatro fragmentos

| Fragmento | Costo en el peor caso | Justificación                                                                                                       |
| --------- | --------------------- | ------------------------------------------------------------------------------------------------------------------- |
| F1        | O(n)                  | Recorre los n datos una vez y en cada vuelta realiza una comparación y, como máximo, un incremento.                 |
| F2        | O(n2)                 | Por cada uno de los n valores de i recorre los n valores de j: n × n pares, con trabajo constante por par.          |
| F3        | O(log n)              | Después de t vueltas, k = 2^t; alcanza n en ⌈log₂ n⌉ vueltas.                                                       |
| F4        | O(n log n)            | En cada una de las n vueltas reinicia k = n y lo divide entre 3 hasta llegar a 1 o menos, con O(log3 n) divisiones. |

Las operaciones aritméticas y los accesos a los datos se consideran de costo
constante. En F4, la división entera también da O(log n). La base del logaritmo
no cambia la clase de complejidad. Estas cotas también son ajustadas: Θ.

## b) Dos recurrencias, con T(1) = 1

Se usa la forma `T(n) = a T(n/b) + Θ(n^d)` y se compara `a` con `b^d`.

### (i) T(n) = 4 T(n/3) + n

- `a = 4`, `b = 3`, `d = 1`.
- `b^d = 3^1 = 3`; por tanto, `a > b^d` porque `4 > 3`.
- Domina la ramificación recursiva: **T(n) = Θ(n^(log3 4)) ≈ Θ(n^1,262)**.

### (ii) T(n) = 2 T(n/2) + n²

- `a = 2`, `b = 2`, `d = 2`.
- `b^d = 2² = 4`; por tanto, `a < b^d` porque `2 < 4`.
- Domina el trabajo fuera de las llamadas: **T(n) = Θ(n²)**.

En (ii) se cumple también la condición de regularidad:
`2(n/2)² = n²/2`, menor que `c n²` con `c = 1/2 < 1`.
