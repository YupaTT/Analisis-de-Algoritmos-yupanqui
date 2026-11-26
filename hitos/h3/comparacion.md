# H3: comparación medida

**Problema:** para cada una de las 3 sucursales, encontrar la racha de meses seguidos en que más crecieron las ventas respecto del mes anterior (ver [estrategia.md](estrategia.md)).

- **Fuerza bruta:** prueba todos los rangos de meses `desde..hasta` y suma cada uno desde cero.
- **Programación dinámica:** recorre los meses una vez; la mejor racha que termina este mes sale de la del mes anterior.

**Mismo resultado:** antes de medir, el programa comprueba que las dos den la misma mejor racha en las 3 sucursales. Si no, se detiene con un error.

**Tamaños:** historial de **6, 12, 24 y 48 meses** (medio año, 1, 2 y 4 años). Cada tiempo es lo que tardan **1.000 consultas** del reporte de las 3 sucursales.

**Protocolo:** 1 calentamiento, 3 corridas y promedio. Las ventas son simuladas (semilla fija) y cada mes suben o bajan hasta 50 unidades. Se corre con `go run .` dentro de `hitos/h3/`. «Veces más» divide el promedio de cada fila entre el de la fila anterior.

- **Fuerza bruta: O(n³).** Prueba los n² rangos y suma cada uno desde cero: al duplicar los meses debería tardar unas 8 veces más.
- **Dinámica: O(n).** Recorre los meses una vez: al duplicar debería tardar unas 2 veces más.

Promedios de las 3 corridas (cada corrida está en [resultados.txt](resultados.txt)):

| Meses | Fuerza bruta (ms) | Veces más | Dinámica (ms) | Veces más |
| ----: | ----------------: | --------: | ------------: | --------: |
|     6 |             0,704 |         — |         0,048 |         — |
|    12 |             3,390 |      4,82 |         0,104 |      2,15 |
|    24 |            22,191 |      6,55 |         0,156 |      1,50 |
|    48 |           186,860 |      8,42 |         0,730 |      4,70 |

## Conclusión

- Al duplicar los meses, la fuerza bruta tarda entre 4,8 y 8,4 veces más, y se acerca al ×8 de O(n³) cuanto más grande es n (con pocos meses pesa más el costo fijo de cada llamada).
- La dinámica tarda unas 2 veces más, porque recorre cada mes una sola vez. Con 48 meses tardó 0,73 ms contra 187 ms de la fuerza bruta (unas 256 veces menos), y dio el mismo resultado.
- Hay ruido: el ×4,70 de la dinámica en 48 meses sale de tiempos menores a 1 ms; con valores tan chicos cualquier interrupción del sistema cambia el cociente.
