# H2 — Medición

**Problema:** 10.000 consultas de los 5 productos más vendidos, rotando las 6 categorías. Las dos versiones devuelven el mismo resultado, y el programa lo verifica antes de medir.

- **Ingenua:** una lista que cada consulta recorre entera, filtra y ordena (`topIngenua` en [main.go](main.go)).
- **Elegida:** un mapa categoría → lista ya ordenada (`topMapa`).

**Protocolo:** 1 calentamiento y 3 corridas por tamaño y versión. Promedio = suma / 3. Se mide solo la fase de consultas: generar los datos y construir el mapa quedan fuera, porque son un costo único.

"Veces más" divide el tiempo de cada fila entre el de la fila anterior.

|    n | Ingenua (ms) | Veces más | Elegida (ms) | Veces más |
| ---: | -----------: | --------: | -----------: | --------: |
| 1000 |      213,205 |         — |       0,1974 |         — |
| 2000 |      630,710 |      2,96 |       0,1224 |      0,62 |
| 4000 |     1665,069 |      2,64 |       0,1197 |      0,98 |
| 8000 |     3641,556 |      2,19 |       0,1257 |      1,05 |

## Conclusión

La ingenua multiplica su tiempo por algo más de 2 cada vez que n se duplica (2,2 a 3,0). Eso se parece a O(n): recorre todos los productos, y el orden de k ≈ n/6 elementos suma un poco más.
La elegida queda casi igual al crecer n (≈0,12 a 0,20 ms para 10.000 consultas), como corresponde a O(1) por consulta. Con n = 8.000 es unas 29.000 veces más rápida.
Hay ruido: los tiempos de la elegida son tan pequeños que sus cocientes varían entre corridas (de 0,62 a 1,40), y en una segunda corrida la ingenua con 8.000 dio 3.802,592 ms.
