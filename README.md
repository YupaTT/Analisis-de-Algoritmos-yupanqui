# Análisis de Algoritmos — Yupanqui

Expediente del caso integrador de la materia Análisis de Algoritmos.

**Mi proceso:** el reporte de ventas por rango de fechas (y, si se pide, por sucursal) que muestra los productos más vendidos. Hoy tarda 864 ms por mes (120.000 líneas) y 11.152 ms por año (1.440.000 líneas), porque cada consulta recorre todas las líneas, busca cada producto uno por uno y ordena con Inserción.

## Conclusión: cómo se mejora

Cambio los tres pasos lentos por tres algoritmos del curso:

| Paso                      | Antes                           | Ahora                                                                    |
| ------------------------- | ------------------------------- | ------------------------------------------------------------------------ |
| Hallar las fechas pedidas | Recorrer todas las líneas       | **Búsqueda binaria** (las líneas se ordenan por fecha una vez al cargar) |
| Sumar por producto        | Buscar el producto en una lista | **Mapa** producto → total                                                |
| Ordenar los más vendidos  | Inserción                       | **MergeSort** (rápido y deja los empates en el mismo orden)              |

El reporte da el mismo resultado y tarda mucho menos:

| Reporte                   |     Antes |      Después |
| ------------------------- | --------: | -----------: |
| Un mes (120.000 líneas)   |    864 ms |   9 ms (×98) |
| Un año (1.440.000 líneas) | 11.152 ms | 52 ms (×216) |

Medido con el protocolo (1 calentamiento, 3 corridas, promedio) en la misma corrida para las dos versiones. Detalle en [conclusion/](conclusion/README.md).

## Índice

| Parte                                                                 | Qué contiene                                                                                                                                                                 |
| --------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [h1/](h1/H1.md)                                                       | El proceso lento, quién lo sufre, volumen, tiempo actual y su O.                                                                                                             |
| [practicas/](practicas/README.md)                                     | Prácticas de clase: c3 búsqueda binaria, c4 lineal contra binaria, c5 torneo de ordenamientos, c6 elección del ordenamiento, estructuras medidas, BFS y DFS.                 |
| [hitos/h2/](hitos/h2/estructura.md)                                   | Estructura de datos elegida, con su [inventario](hitos/h2/inventario.md) y [medición](hitos/h2/medicion.md).                                                                 |
| [hitos/h3/](hitos/h3/estrategia.md)                                   | Crecimiento mensual de ventas por sucursal: por qué el voraz falla y la programación dinámica sirve, y su [comparación medida](hitos/h3/comparacion.md) contra fuerza bruta. |
| [parcial1/](parcial1/analisis.md) y [parcial2/](parcial2/eleccion.md) | Parciales de la materia.                                                                                                                                                     |
| [conclusion/](conclusion/README.md)                                   | El reporte del H1 mejorado (binaria + mapa + MergeSort), verificado y [medido](conclusion/medicion.md) antes y después.                                                      |
| [integradora/](integradora/analisis.md)                               | Compra de útiles con presupuesto (mochila 0/1): fuerza bruta contra programación dinámica, NP y su [medición](integradora/medicion.md).                                      |

## Dónde está cada tema medido

Todas las mediciones usan el mismo protocolo: **1 calentamiento, 3 corridas y promedio**, con n que se duplica para ver cómo crece el tiempo.

| Tema                                        | Evidencia                                                                                                                                          |
| ------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| O grande y recurrencias                     | [parcial1/analisis.md](parcial1/analisis.md) (teoría), [integradora/analisis.md](integradora/analisis.md) (T(n) = 2T(n−1) + 1)                     |
| Medición con protocolo                      | [h1/mediciones.txt](h1/mediciones.txt) y todas las tablas de abajo                                                                                 |
| Lineal y binaria                            | [practicas/c4](practicas/c4/README.md)                                                                                                             |
| Burbuja, Inserción, Merge, Quick y Heap     | [practicas/c5](practicas/c5/README.md), [parcial1/medicion.md](parcial1/medicion.md) (pares contra MergeSort)                                      |
| Pila, cola, lista, hash, BST y heap         | [practicas/estructuras](practicas/estructuras/README.md), [hitos/h2](hitos/h2/medicion.md) (hash), [parcial2](parcial2/medicion.md) (colas + mapa) |
| Voraz                                       | [hitos/h3/estrategia.md](hitos/h3/estrategia.md) y [integradora/analisis.md](integradora/analisis.md): contraejemplos donde falla                  |
| Programación dinámica y fuerza bruta        | [hitos/h3](hitos/h3/comparacion.md), [integradora](integradora/medicion.md)                                                                        |
| NP                                          | [integradora/analisis.md](integradora/analisis.md#e-es-un-problema-np)                                                                             |
| BFS y DFS                                   | [practicas/grafos](practicas/grafos/README.md)                                                                                                     |
| Mejora del proceso del H1 (antes y después) | [conclusion/](conclusion/README.md): binaria + mapa + MergeSort, [medición](conclusion/medicion.md)                                                |
