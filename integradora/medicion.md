Gonzalo Yupanqui Torrez

# I3 — Medición

- **Cómo se corre:** `go run .` dentro de `integradora/` (salida completa en `resultados.txt`).
- **Protocolo:** 1 corrida de calentamiento (no cuenta) + 3 corridas medidas + promedio.
- **Carga:** la de `EjecutarCarga` portada a Go: semilla n+1000, costo 4..49, utilidad 10..99, presupuesto = un tercio del costo total. (El azar de Go no genera los mismos números que el de .NET, pero la carga es igual de grande.)
- **Tamaños:** n = 20, 22, 24, 26. Elegí 20 porque con n = 26 la ingenua tarda unos 7,3 s (entre 0,5 y 20 s); con n = 28 pasaría de 20 s (~×4).

## Ingenuo: O(2ⁿ · n)

| n   | Corrida 1 (ms) | Corrida 2 (ms) | Corrida 3 (ms) | Promedio (ms) | Veces vs fila anterior | Utilidad |
| --- | -------------- | -------------- | -------------- | ------------- | ---------------------- | -------- |
| 20  | 77,98          | 80,43          | 79,96          | 79,46         | —                      | 742      |
| 22  | 488,23         | 424,90         | 323,51         | 412,21        | ×5,19                  | 572      |
| 24  | 1.677,93       | 1.378,07       | 1.835,62       | 1.630,54      | ×3,96                  | 806      |
| 26  | 7.240,10       | 7.221,08       | 7.414,13       | 7.291,77      | ×4,47                  | 821      |

## Mejorado (programación dinámica): O(n · P)

| n   | Corrida 1 (ms) | Corrida 2 (ms) | Corrida 3 (ms) | Promedio (ms) | Veces vs fila anterior | Utilidad |
| --- | -------------- | -------------- | -------------- | ------------- | ---------------------- | -------- |
| 20  | 0,033          | 0,029          | 0,024          | 0,029         | —                      | 742      |
| 22  | 0,055          | 0,053          | 0,546          | 0,218         | ×7,52                  | 572      |
| 24  | 0,094          | 0,050          | 0,054          | 0,066         | ×0,30                  | 806      |
| 26  | 0,049          | 0,051          | 0,050          | 0,050         | ×0,76                  | 821      |

## Verificación: las dos dan la misma respuesta

El programa compara la utilidad de las dos soluciones en cada tamaño e imprime `Misma utilidad: true` (ver `resultados.txt`):

| n   | Ingenuo | Mejorado | ¿Iguales? |
| --- | ------- | -------- | --------- |
| 20  | 742     | 742      | sí        |
| 22  | 572     | 572      | sí        |
| 24  | 806     | 806      | sí        |
| 26  | 821     | 821      | sí        |

Además las dos pasan las 8 pruebas de `main` (utilidad y útiles elegidos en los 6 casos, y `EstaEnLaCompra`).

## Conclusión

- **Ingenuo:** los cocientes son ×5,19, ×3,96 y ×4,47; en promedio rondan el ×4 esperado (cada +2 útiles hay 4 veces más subconjuntos; el factor n suma un poco). Se nota ruido: en n = 22 las corridas van de 324 a 488 ms (50 % de diferencia) y en n = 24 de 1.378 a 1.836 ms; eso infla el primer cociente y achica el segundo. La fila de n = 26, con corridas muy parejas (7.221 a 7.414 ms), es la más confiable.
- **Mejorado:** tarda centésimas de ms, y sus cocientes (×7,52, ×0,30, ×0,76) no dicen nada: el ×7,52 sale de una sola corrida de 0,546 ms en n = 22 (las otras dos dieron 0,05 ms), puro ruido del sistema. Además P cambia con cada carga al azar, así que n · P no crece de forma pareja. Lo honesto es decir que se mantiene en centésimas de ms en los cuatro tamaños.
- En n = 26 el mejorado es unas **146.000 veces más rápido** (7.292 ms contra 0,050 ms), y la brecha se multiplica por ~4 con cada +2 útiles.

## Mejorado con 1.000 útiles y presupuesto 10.000

Corridas: 159,71 / 82,53 / 133,89 ms → **promedio 125,38 ms** (utilidad 35.173). Son ~10 millones de casillas (1.001 × 10.001) en cada tabla. Las corridas varían bastante (ruido). La ingenua con 1.000 útiles necesitaría 2¹⁰⁰⁰ subconjuntos: imposible.

## ¿Y con presupuesto de 1.000.000.000 Bs?

El costo O(n · P) depende del **valor** del presupuesto, no de cuántos datos hay (se llama pseudo-polinomial). Con P = 10⁹ y 1.000 útiles la tabla tendría 1.000 × 10⁹ = 10¹² celdas: unos **8 TB** de memoria (8 bytes por `int`) y muchísimo tiempo. Ni una sola fila (10⁹ celdas ≈ 8 GB) entraría cómodamente. Mi solución **no sirve** en ese caso.

Qué haría: si el presupuesto alcanza para todo (suma de costos ≤ P), compro todo directamente. Si no, cambio la tabla para indexarla por **utilidad** en vez de por Bs: «costo mínimo para lograr utilidad u», que cuesta O(n · U) con U = suma de utilidades; aquí las utilidades son ≤ 99, así que U ≤ 99.000 para 1.000 útiles, sin importar el presupuesto.
