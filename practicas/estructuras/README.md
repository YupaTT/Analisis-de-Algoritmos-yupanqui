# Estructuras de datos medidas

Se mide cada estructura en la operación que la caracteriza, con n que se duplica en cada fila.
Protocolo: **1 calentamiento, 3 corridas y promedio**. Armar las estructuras queda fuera del tiempo; solo se mide la operación.
Las 3 corridas de cada fila están en [resultados.txt](resultados.txt). Promedios en ms.

## 1) Pila y cola: meter n elementos y sacarlos todos

Cada operación es O(1), así que el total es O(n): al duplicar n debería tardar unas 2 veces más.

|         n | Pila (ms) | Veces | Cola (ms) | Veces |
| --------: | --------: | ----: | --------: | ----: |
| 1.000.000 |     55,55 |     — |     56,75 |     — |
| 2.000.000 |    131,25 | ×2,36 |    116,99 | ×2,06 |
| 4.000.000 |    226,26 | ×1,72 |    178,96 | ×1,53 |
| 8.000.000 |    317,79 | ×1,40 |    461,39 | ×2,58 |

La pila saca del final (último en entrar, primero en salir) y la cola del frente (primero en entrar, primero en salir). Las dos crecen cerca de ×2.

## 2) Buscar 1.000 claves: lista, BST y hash

|      n | Lista O(n) | Veces | BST O(log n) | Veces | Hash O(1) | Veces |
| -----: | ---------: | ----: | -----------: | ----: | --------: | ----: |
| 10.000 |       5,63 |     — |        0,352 |     — |     0,029 |     — |
| 20.000 |       8,83 | ×1,57 |        0,355 | ×1,01 |     0,033 | ×1,12 |
| 40.000 |      17,22 | ×1,95 |        0,332 | ×0,93 |     0,033 | ×1,02 |
| 80.000 |      30,41 | ×1,77 |        0,350 | ×1,05 |     0,017 | ×0,52 |

- **Lista:** recorre hasta encontrar la clave; crece cerca de ×2.
- **BST:** baja por el árbol, unas log₂ n ≈ 14 a 17 comparaciones; casi no cambia. Las claves se insertan en orden aleatorio para que el árbol no quede como una lista.
- **Hash:** va directo a la clave; no cambia con n.

## 3) Sacar el mayor 1.000 veces (top-k): lista y heap

|      n | Lista O(n) | Veces | Heap O(log n) | Veces |
| -----: | ---------: | ----: | ------------: | ----: |
| 10.000 |      14,16 |     — |         0,209 |     — |
| 20.000 |      32,46 | ×2,29 |         0,177 | ×0,85 |
| 40.000 |      54,47 | ×1,68 |         0,243 | ×1,37 |
| 80.000 |     105,94 | ×1,94 |         0,199 | ×0,82 |

La lista recorre todo para hallar el mayor; el heap lo tiene en la raíz y solo repara una rama. Con 80.000 el heap es unas **530 veces más rápido**. El programa verifica que los dos saquen los mismos mayores.

## Conclusión

- Lo medido coincide con la O de cada estructura: la lista crece ×2, BST, hash y heap casi no crecen.
- Hay ruido en los tiempos pequeños (menos de 0,5 ms): sus cocientes van de ×0,5 a ×1,4 sin que la estructura cambie de comportamiento.
- Para el H1: hash para buscar producto o sucursal, BST para rangos de fechas y heap para los k más vendidos.
