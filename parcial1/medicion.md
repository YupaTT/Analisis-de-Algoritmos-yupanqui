Gonzalo Yupanqui Torrez

# P1.2 — Medición

- **Cómo se corre:** `go run parcial1/main.go` desde la raíz (salida completa en [resultados.txt](resultados.txt)).
- **Protocolo:** 1 corrida de calentamiento (no cuenta) + 3 corridas medidas + promedio.
- **Datos:** n comprobantes sin repetidos (peor caso: la lenta revisa todos los pares). n se duplica en cada fila.

## Lenta (comparar pares): O(n²)

| n      | Corrida 1 (ms) | Corrida 2 (ms) | Corrida 3 (ms) | Promedio (ms) | Veces vs fila anterior |
| ------ | -------------: | -------------: | -------------: | ------------: | ---------------------: |
| 10.000 |         101,83 |         112,48 |          83,35 |         99,22 | —                      |
| 20.000 |         314,02 |         168,70 |         207,87 |        230,19 | ×2,32                  |
| 40.000 |         790,96 |         679,11 |         812,17 |        760,75 | ×3,30                  |
| 80.000 |       2.597,70 |       2.437,97 |       3.315,58 |      2.783,75 | ×3,66                  |

## Cura (ordenar con MergeSort y mirar vecinos): O(n log n)

| n      | Corrida 1 (ms) | Corrida 2 (ms) | Corrida 3 (ms) | Promedio (ms) | Veces vs fila anterior |
| ------ | -------------: | -------------: | -------------: | ------------: | ---------------------: |
| 10.000 |           2,13 |           2,11 |           2,86 |          2,37 | —                      |
| 20.000 |           1,86 |           2,97 |           1,93 |          2,25 | ×0,95                  |
| 40.000 |           5,63 |           4,92 |           4,98 |          5,18 | ×2,30                  |
| 80.000 |           9,47 |          18,64 |          12,85 |         13,65 | ×2,64                  |

## Conclusión

- **Lenta:** al duplicar n tarda ×3,3 y ×3,7, cerca del ×4 de O(n²). El primer salto (×2,32) es ruido: las corridas de 20.000 van de 169 a 314 ms.
- **Cura:** crece entre ×2,3 y ×2,6, cerca del ×2 de O(n log n). Con 2 ms el ruido manda (el ×0,95).
- Con 80.000 datos la cura es unas **204 veces más rápida** (2.784 ms contra 13,65 ms), y las dos dan el mismo resultado en las 5 pruebas.
