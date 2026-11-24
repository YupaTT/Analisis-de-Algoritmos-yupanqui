# C4: búsqueda lineal contra binaria

Datos: **1.000.000 de enteros ordenados y 1.000 claves que sí existen**.
Protocolo: **un calentamiento, 3 corridas y promedio**.
Cada tiempo corresponde a las 1.000 búsquedas y no incluye ordenar.
[Resultados](resultados.txt) en milisegundos. El cociente es lineal / binaria:

| Corrida  | Lineal (ms) | Binaria (ms) | Lineal / binaria |
| -------- | ----------: | -----------: | ---------------: |
| 1        |  642.213896 |     0.401617 |          1599.07 |
| 2        |  689.964461 |     0.683260 |          1009.81 |
| 3        |  433.182212 |     0.299756 |          1445.12 |
| Promedio |  588.453523 |     0.461544 |          1274.97 |

En mi proceso (H1) **sí conviene ordenar primero** las líneas de venta por fecha: se ordenan una vez y cada consulta halla `desde..hasta` con dos binarias, en lugar de recorrer las 120.000 líneas.
Conviene porque los datos no cambian entre consultas y el reporte se pide muchas veces, así que el costo de ordenar se paga una sola vez y se recupera en pocas consultas.
