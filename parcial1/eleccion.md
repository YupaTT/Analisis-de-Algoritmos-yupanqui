Gonzalo Yupanqui Torrez

# P1.3 — Elección para el almacén

Se tienen **n = 250.000 códigos** y **q = 8.000 consultas por día**.
Para comparar uso el peor caso de las búsquedas y un modelo aproximado de
operaciones, con coeficiente 1 para `n`, `n log₂ n` y `n²`.
`L = log₂(250.000) ≈ 17,93156857`; una búsqueda binaria necesita como máximo
18 iteraciones. Las cuentas usan L sin redondear y redondean solo el total.

## Costos diarios

| Opción                                     |                       Preparación |                     Consultas | Total aproximado de operaciones |
| ------------------------------------------ | --------------------------------: | ----------------------------: | ------------------------------: |
| 1. Recorrer en cada consulta               |                                 0 |               8.000 × 250.000 |               **2.000.000.000** |
| 2. Ordenar en n log n y buscar con binaria | 250.000 × 17,93156857 ≈ 4.482.892 | 8.000 × 17,93156857 ≈ 143.453 |                   **4.626.345** |
| 3. Ordenar en n² y buscar con binaria      |         250.000² = 62.500.000.000 | 8.000 × 17,93156857 ≈ 143.453 |              **62.500.143.453** |

Las formas son, respectivamente, `O(qn)`, `O(n log n + q log n)` y
`O(n² + q log n)`. En la opción 1, una consulta ausente recorre los 250.000
códigos. En las opciones 2 y 3, el ordenamiento se paga una vez cada mañana.

## Elección y memoria

Elijo la **opción 2**, usando **Heapsort sobre el mismo listado**, que garantiza
O(n log n) en el peor caso y requiere O(1) de memoria auxiliar si se implementa
iterativamente. Después hago las búsquedas binarias sin copiar el listado.
Reordenar el listado está permitido por estas opciones; duplicarlo no.
Así se respeta la memoria limitada y se evita el costo cuadrático inicial.

En [P1.2](medicion.md), con 80.000 datos, comparar pares promedió
**4418,495012 ms** y la cura basada en ordenar **19,686146 ms**: unas 224 veces
menos tiempo. La diferencia y el crecimiento observado apoyan preferir
n log n frente a n² cuando aumenta n. Es evidencia sobre esas dos funciones,
no una medición de Heapsort ni de las 8.000 consultas; no convierto esos tiempos
en una predicción exacta para el almacén. La copia usada en P1.2 responde a la
prohibición de modificar aquella entrada y no se usa en esta elección.

## ¿Desde cuántas consultas cambiaría la decisión?

Comparo la opción 1 con la 2, incluyendo el ordenamiento inicial:

```text
C1(q) = q × n
C2(q) = n × L + q × L

C2(q) < C1(q)
n × L < q × (n − L)
q > n × L / (n − L)
q > 4.482.892,142331 / 249.982,068431
q > 17,932855
```

En este modelo elegiría recorrer para **1 a 17 consultas** y cambiaría a
ordenar en n log n más búsqueda binaria **desde 18 consultas**:

| Consultas |  Recorrer | Ordenar en n log n + binaria |
| --------: | --------: | ---------------------------: |
|        17 | 4.250.000 |                  ≈ 4.483.197 |
|        18 | 4.500.000 |                  ≈ 4.483.215 |

Con las 8.000 consultas del problema, la opción 2 es claramente favorable.
La opción 3 no pasa a ser mejor que la 2 al aumentar las consultas: ambas
pagan lo mismo por buscar y la 3 paga mucho más por ordenar.
El umbral de 18 pertenece a esta estimación con constantes unitarias;
el umbral real depende de las constantes de cada implementación y requeriría
medir el ordenamiento y las consultas del almacén.
