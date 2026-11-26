Gonzalo Yupanqui Torrez

# Medición — Reporte del H1, antes y después

- **Cómo se corre:** `go run .` dentro de `conclusion/` (salida en [resultados.txt](resultados.txt)).
- **Protocolo:** 1 calentamiento (no cuenta) + 3 corridas + promedio.
- **Datos:** como el H1: semilla 42, 8 sucursales, 8.000 productos. Los meses tienen todas las líneas en octubre; el año las reparte en 12 meses.
- **Tamaños:** los de [h1/mediciones.txt](../h1/mediciones.txt) (30.000 a 240.000, se duplican) y un año (1.440.000 = 12 × 120.000).

## Antes: O(n + m·p + p²)

| Datos | Corrida 1 (ms) | Corrida 2 (ms) | Corrida 3 (ms) | Promedio (ms) | Veces vs fila anterior |
| --- | ---: | ---: | ---: | ---: | ---: |
| Un mes, 30.000 líneas | 306,96 | 306,81 | 290,25 | 301,34 | — |
| Un mes, 60.000 líneas | 449,71 | 447,34 | 356,84 | 417,96 | ×1,39 |
| Un mes, 120.000 líneas | 786,89 | 1.041,96 | 763,62 | 864,16 | ×2,07 |
| Un mes, 240.000 líneas | 1.982,44 | 2.441,74 | 1.744,27 | 2.056,15 | ×2,38 |
| Un año, reporte de octubre | 1.284,14 | 1.049,01 | 810,39 | 1.047,85 | — |
| Un año, reporte del año | 10.519,02 | 12.327,48 | 10.608,88 | 11.151,79 | — |

## Después: O(log n + m + p log p)

| Datos | Corrida 1 (ms) | Corrida 2 (ms) | Corrida 3 (ms) | Promedio (ms) | Veces vs fila anterior |
| --- | ---: | ---: | ---: | ---: | ---: |
| Un mes, 30.000 líneas | 17,93 | 20,52 | 24,99 | 21,15 | — |
| Un mes, 60.000 líneas | 16,44 | 14,89 | 14,33 | 15,22 | ×0,72 |
| Un mes, 120.000 líneas | 9,05 | 8,70 | 8,70 | 8,82 | ×0,58 |
| Un mes, 240.000 líneas | 27,53 | 28,68 | 28,37 | 28,19 | ×3,20 |
| Un año, reporte de octubre | 10,62 | 10,49 | 11,67 | 10,93 | — |
| Un año, reporte del año | 57,88 | 48,89 | 48,45 | 51,74 | — |

## Resumen

| Datos | Antes (ms) | Después (ms) | Veces más rápido | Ordenar por fecha al cargar (ms, una vez) |
| --- | ---: | ---: | ---: | ---: |
| Un mes, 30.000 líneas | 301,34 | 21,15 | ×14 | 44,65 |
| Un mes, 60.000 líneas | 417,96 | 15,22 | ×27 | 183,47 |
| Un mes, 120.000 líneas | 864,16 | 8,82 | ×98 | 132,69 |
| Un mes, 240.000 líneas | 2.056,15 | 28,19 | ×73 | 436,62 |
| Un año, reporte de octubre | 1.047,85 | 10,93 | ×96 | 3.952,22 |
| Un año, reporte del año | 11.151,79 | 51,74 | ×216 | 3.477,93 |

**Mismo resultado:** antes de medir, el programa compara los dos reportes en cada caso (todas las sucursales y sucursal 1) e imprime `Mismo reporte en todos los casos: sí`.

## Conclusión

- **Antes** crece cerca de ×2 al duplicar las líneas (×1,39, ×2,07, ×2,38), como en el H1.
- **Después** se queda en decenas de ms en todos los tamaños: depende de los 8.000 productos más que de las líneas. Con un año completo es unas **216 veces más rápido**.
- Ordenar por fecha al cargar se paga una vez: con 120.000 líneas (133 ms) se recupera en la primera consulta; con un año (3.952 ms), en unas 4 consultas de un mes.
- **Ruido:** la máquina estaba cargada por otros programas, así que los tiempos salieron unas 3 veces más altos que en el H1 (864 ms contra 232 ms para 120.000 líneas). Los cocientes del después (×0,58 a ×3,20) son ruido de tiempos chicos. Como las dos versiones se midieron juntas y con los mismos datos, la comparación entre ellas sí es justa.
