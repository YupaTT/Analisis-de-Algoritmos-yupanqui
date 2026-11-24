# Análisis de Algoritmos — Yupanqui

Expediente del caso integrador de la materia Análisis de Algoritmos.

**Mi proceso:** el reporte de ventas por rango de fechas (y, si se pide, por sucursal) que muestra los productos más vendidos. Hoy tarda 0,23 s por mes y 2,6 s por año, porque cada consulta recorre todas las líneas, busca cada producto uno por uno y ordena con Inserción.

## Índice

| Parte | Qué contiene |
| --- | --- |
| [h1/](h1/H1.md) | El proceso lento, quién lo sufre, volumen, tiempo actual y su O. |
| [practicas/](practicas/README.md) | Prácticas de clase: c3 búsqueda binaria, c4 lineal contra binaria, c5 torneo de ordenamientos, c6 elección del ordenamiento. |
| [hitos/h2/](hitos/h2/estructura.md) | Estructura de datos elegida, con su [inventario](hitos/h2/inventario.md) y [medición](hitos/h2/medicion.md). |
| [hitos/h3/](hitos/h3/estrategia.md) | Crecimiento mensual de ventas por sucursal: por qué el voraz falla y la programación dinámica sirve, y su [comparación medida](hitos/h3/comparacion.md) contra fuerza bruta. |
| [parcial1/](parcial1/analisis.md) y [parcial2/](parcial2/eleccion.md) | Parciales de la materia. |

## Estado: qué falta para la integradora y la defensa

- [x] H1: proceso lento descrito y medido.
- [x] H2: estructura elegida y medida.
- [x] H3: estrategia argumentada y comparada.
- [ ] Integradora: juntar las mejoras en una versión optimizada del reporte (mapa producto → total, líneas ordenadas por fecha con búsqueda binaria y MergeSort) y verificar que da el mismo resultado que la versión actual.
- [ ] Medir la versión optimizada contra la actual con el protocolo (calentar, 3 corridas, promedio).
- [ ] Defensa: preparar la explicación de cada decisión con su evidencia medida.
