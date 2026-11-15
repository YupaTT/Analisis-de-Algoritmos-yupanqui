# C4: búsqueda lineal contra binaria

Datos: **1.000.000 de enteros ordenados y 1.000 claves que sí existen**.
Protocolo: **un calentamiento, 3 corridas y promedio**.
Cada tiempo corresponde a las 1.000 búsquedas; no incluye ordenar.
[Resultados](resultados.txt) en milisegundos; cociente = lineal / binaria:

| Corrida  | Lineal (ms) | Binaria (ms) | Lineal / binaria |
| -------- | ----------: | -----------: | ---------------: |
| 1        |  642.213896 |     0.401617 |          1599.07 |
| 2        |  689.964461 |     0.683260 |          1009.81 |
| 3        |  433.182212 |     0.299756 |          1445.12 |
| Promedio |  588.453523 |     0.461544 |          1274.97 |

En H1 no ordenaría primero; reutilizaría los productos por identificador.  
Así evitaría consultas repetidas, que ordenar por sí solo no elimina.
