# C5: torneo de ordenamientos

Se comparan **Burbuja, Inserción, MergeSort, QuickSort y HeapSort** con **5.000, 10.000, 20.000 y 40.000 elementos** (se duplica en cada fila).
Los datos son **aleatorios y con repetidos** (valores 0..999, semilla fija), como las cantidades vendidas por producto en el H1.
Protocolo: **1 calentamiento, 3 corridas y promedio**. Cada corrida ordena una copia nueva y el programa verifica que quede ordenada.
Las 3 corridas de cada algoritmo están en [resultados.txt](resultados.txt). Promedios en ms:

| Cantidad | Burbuja O(n²) | Inserción O(n²) | MergeSort O(n log n) | QuickSort O(n log n) | HeapSort O(n log n) |
| -------: | ------------: | --------------: | -------------------: | -------------------: | ------------------: |
|    5.000 |         25,09 |            7,00 |                 1,09 |                 0,58 |                0,78 |
|   10.000 |        115,54 |           28,43 |                 2,86 |                 1,44 |                2,19 |
|   20.000 |        963,84 |          171,21 |                 5,72 |                 2,91 |                6,27 |
|   40.000 |      4.958,36 |          511,20 |                19,35 |                 9,52 |                8,81 |

Veces más al duplicar n (promedio de los 3 saltos):

| Burbuja | Inserción | MergeSort | QuickSort | HeapSort |
| ------: | --------: | --------: | --------: | -------: |
|   ×6,03 |     ×4,36 |     ×2,67 |     ×2,59 |    ×2,36 |

## Conclusiones

- **Burbuja e Inserción son cuadráticas:** al duplicar los datos tardan unas 4 veces más o más. Burbuja es la peor porque hace muchos más intercambios.
- **MergeSort, QuickSort y HeapSort son n log n:** al duplicar tardan algo más de 2 veces más.
- **QuickSort fue el más rápido** en casi todos los tamaños; HeapSort lo alcanzó en 40.000.
- Con 40.000 datos, Burbuja tardó casi 5 s; QuickSort, menos de 10 ms (unas 520 veces menos).
- Hay ruido: algunos saltos se alejan del ×4 o ×2 (por ejemplo Burbuja ×8,34 o HeapSort ×1,41), pero la tendencia es clara.
