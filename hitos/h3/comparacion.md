# H3: comparación medida

**Problema:** para cada una de las 3 sucursales, encontrar la racha de meses seguidos en que más crecieron las ventas respecto del mes anterior (ver [estrategia.md](estrategia.md)).

- **Fuerza bruta:** prueba todos los rangos de meses `desde..hasta` y suma cada uno desde cero.
- **Programación dinámica:** recorre los meses una vez; la mejor racha que termina este mes sale de la del mes anterior.

**Mismo resultado:** antes de medir, el programa comprueba que las dos den la misma mejor racha en las 3 sucursales. Si no, se detiene con un error.

**Tamaños:** historial de **6, 12, 24 y 48 meses** (medio año, 1, 2 y 4 años). Cada tiempo es lo que tardan **1.000 consultas** del reporte de las 3 sucursales.

**Protocolo:** 1 calentamiento, 3 corridas y promedio. Las ventas son simuladas (semilla fija) y cada mes suben o bajan hasta 50 unidades. «Veces más» divide el tiempo de cada fila entre el de la fila anterior.

| Meses | Fuerza bruta (ms) | Veces más | Dinámica (ms) | Veces más |
| ----: | ----------------: | --------: | ------------: | --------: |
|     6 |             0,165 |         — |         0,021 |         — |
|    12 |             0,962 |      5,82 |         0,043 |      2,03 |
|    24 |             3,938 |      4,09 |         0,077 |      1,80 |
|    48 |            27,096 |      6,88 |         0,167 |      2,16 |

## Conclusión

- Al duplicar los meses, la fuerza bruta tarda entre 4 y 7 veces más, porque prueba todos los rangos y suma cada uno desde cero.
- La dinámica tarda unas 2 veces más, porque recorre cada mes una sola vez. Con 48 meses tardó 0,17 ms contra 27 ms de la fuerza bruta (unas 160 veces menos), y dio el mismo resultado.
- Hay ruido: los cocientes de la fuerza bruta cambian entre corridas. En una segunda corrida dieron 4,09, 6,37 y 5,98. Los de la dinámica se mantuvieron cerca de 2 (1,81, 1,99 y 2,15).
