# C5: torneo de ordenamientos

Se comparan Inserción, MergeSort y QuickSort con **10.000, 20.000, 40.000 y 80.000 elementos**, como en clase. La cantidad se duplica en cada fila.
Los datos son **aleatorios y con repetidos** (semilla fija), como las cantidades vendidas por producto en el H1.
Protocolo: **un calentamiento, 3 corridas y promedio**. Antes de sumar cada corrida, el programa verifica que los tres algoritmos dejen los datos ordenados.
[Resultados](resultados.txt) en milisegundos:

| Cantidad | Inserción (ms) | MergeSort (ms) | QuickSort (ms) |
| -------: | -------------: | -------------: | -------------: |
|   10.000 |      47.127309 |       0.990322 |       0.168521 |
|   20.000 |     256.327853 |       3.426887 |       0.702519 |
|   40.000 |    1370.412899 |       9.145732 |       1.779423 |
|   80.000 |    3293.205414 |      10.012564 |       2.169780 |

## Conclusiones

- **Inserción es la más lenta:** cada vez que se duplican los datos, tarda unas 4 veces más.
- **MergeSort y QuickSort son rápidos:** cada vez que se duplican los datos, tardan solo unas 2 veces más.
- **QuickSort fue el más rápido**, unas 2 veces más que MergeSort.
- Con 80.000 datos, Inserción tardó casi 800 ms; QuickSort, 5 ms.
