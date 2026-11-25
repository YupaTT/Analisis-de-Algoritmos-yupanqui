# Gonzalo Yupanqui Torrez

# I1 — Analizar y elegir

n = cantidad de utiles, P = presupuesto en Bs..

## a) Costo del ingenuo

**O(2ⁿ · n).** Prueba todas las combinaciones (2ⁿ) y en cada una suma los n utiles.

## b) T(n) = 2 T(n−1) + 1

El teorema maestro sirve cuando el problema se **divide** (n/b). Aquí solo baja en 1 (n−1), así que no se aplica.
Cada util abre 2 ramas (lo compro / no lo compro): 2 · 2 · … · 2 = **O(2ⁿ)**.

## c) Voraz y estrategia elegida

**El voraz no sirve.** Primer caso, presupuesto 100:

| util | Costo | Utilidad | Utilidad por Bs |
| ---- | ----- | -------- | --------------- |
| U1   | 20    | 120      | 6               |
| U2   | 40    | 200      | 5               |
| U3   | 60    | 240      | 4               |

- Voraz: toma U1 + U2 (60 Bs) y U3 ya no entra → **320**.
- Óptimo: U2 + U3 (100 Bs) → **440**.

**Estrategia: programación dinámica.** `dp[i][p]` = mejor utilidad con los primeros i utiles y p Bs:

`dp[i][p] = mayor( no comprarlo: dp[i−1][p] , comprarlo: dp[i−1][p − costo] + utilidad )`

**Qué se repite:** la recursión vuelve a calcular muchas veces «lo mejor con estos utiles y este dinero». La tabla lo calcula una sola vez por casilla.

**utiles elegidos:** la tabla `tomar` marca cuándo se compró cada util; se recorre de atrás hacia adelante restando el costo.

**Costo: O(n · P).**

## d) Estructura para `EstaEnLaCompra`

Un **mapa** (`map[string]bool`): cada consulta es **O(1)**. Recorrer la lista sería O(k) por consulta (k = utiles comprados).
