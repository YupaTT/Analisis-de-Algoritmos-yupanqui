# C6: elección del algoritmo de ordenamiento

**Proceso:** reporte de ventas del [H1](../../h1/H1.md). Al final se ordenan los productos de más a menos vendido.

**Datos:** hasta 8.000 productos, desordenados y con muchas cantidades repetidas.

**Hoy:** el Proceso ordena con Inserción, que compara cada producto con casi todos los demás; por eso, si los productos se duplican, el tiempo se multiplica por cuatro.

## Candidatos

| Algoritmo | Costo                           | ¿Respeta el orden de los empates? |
| --------- | ------------------------------- | --------------------------------- |
| MergeSort | O(n log n)                      | Sí                                |
| QuickSort | O(n log n) promedio, O(n²) peor | No                                |

## Elección: MergeSort

- **Siempre es rápido:** divide la lista por la mitad una y otra vez, y luego junta las partes ya ordenadas. Si los productos se duplican, el tiempo apenas se duplica, no se multiplica por cuatro como con Inserción.
- **No cambia el resultado:** si dos productos vendieron lo mismo, los deja en el mismo orden que hoy. QuickSort puede intercambiarlos y el reporte saldría distinto.

## Evidencia ([c5](../c5/README.md))

|      n | Inserción (ms) | MergeSort (ms) | QuickSort (ms) |
| -----: | -------------: | -------------: | -------------: |
| 10.000 |          28,43 |           2,86 |           1,44 |
| 40.000 |         511,20 |          19,35 |           9,52 |

Con 10.000 datos, MergeSort es 10 veces más rápido que Inserción, y con 40.000, 26 veces. QuickSort es apenas 1,4 ms más rápido con 10.000 (unos 10 ms con 40.000), una diferencia que no compensa cambiar el orden del reporte.
