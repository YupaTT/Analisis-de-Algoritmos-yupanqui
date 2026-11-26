# Grafos, BFS y DFS

Grafo conexo generado con semilla fija: cada nodo se une al siguiente y a 3 nodos al azar (unas 4 aristas por nodo).
Los dos recorridos parten del nodo 0 y el programa verifica que visiten **todos** los nodos.
Protocolo: **1 calentamiento, 3 corridas y promedio**. Las 3 corridas están en [resultados.txt](resultados.txt).

- **BFS** (anchura): usa una **cola**; visita por niveles, primero los vecinos más cercanos.
- **DFS** (profundidad): usa una **pila**; va lo más hondo posible y luego retrocede.

Los dos visitan cada nodo y revisan cada arista una vez: **O(V + E)**. Al duplicar el grafo deberían tardar unas 2 veces más.

|   Nodos | Aristas   | BFS (ms) | Veces | DFS (ms) | Veces |
| ------: | --------- | -------: | ----: | -------: | ----: |
| 100.000 | 400.000   |    16,84 |     — |    41,60 |     — |
| 200.000 | 800.000   |    74,12 | ×4,40 |   127,19 | ×3,06 |
| 400.000 | 1.600.000 |   133,68 | ×1,80 |   206,99 | ×1,63 |
| 800.000 | 3.200.000 |   367,88 | ×2,75 |   639,14 | ×3,09 |

## Conclusión

- En promedio crecen ×3,0 (BFS) y ×2,6 (DFS) al duplicar, algo más que el ×2 esperado. Con grafos grandes los vecinos al azar quedan lejos en la memoria y cada acceso es más lento; además hay ruido (en 200.000 el BFS va de 61 a 91 ms).
- Aun así crecen como una recta, no como n²: con 8 veces más nodos tardan unas 15 a 22 veces más, no 64.
- DFS tardó más que BFS porque mete vecinos repetidos en la pila y los descarta después.
