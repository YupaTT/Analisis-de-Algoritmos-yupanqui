# H2 — Estructura elegida

**Elijo un mapa `categoría → lista de productos`, con cada lista ordenada de mayor a menor vendidos una sola vez, al cargar.**

Por qué:

- La operación más frecuente es buscar por categoría. Un mapa la resuelve en **O(1) en promedio**, sin recorrer los productos de otras categorías.
- La segunda operación es tomar los 5 más vendidos. Si la lista ya está ordenada, alcanza con tomar los primeros 5: **O(1)**.
- El orden se paga **una vez por categoría** al cargar, y no en cada una de las 10.000 consultas.
- No se borran productos ni se piden rangos, así que no necesito un árbol. Con solo 6 categorías, un árbol no aporta nada frente al mapa.

La estructura ingenua es una lista con todos los productos. Cada consulta la recorre entera, filtra la categoría y ordena lo encontrado: **O(n + k log k)** por consulta, donde k ≈ n/6.

| Operación                             | Cuántas veces la hago     | Costo O (estructura elegida) |
| ------------------------------------- | ------------------------- | ---------------------------- |
| Buscar los productos de una categoría | 10.000 (una por consulta) | O(1) en promedio             |
| Tomar los 5 más vendidos              | 10.000 (una por consulta) | O(1)                         |
| Insertar un producto al cargar        | n = 1.000                 | O(1) amortizado              |
| Ordenar cada categoría al cargar      | 6 (una vez cada una)      | O(k log k), k ≈ n/6          |

Costo: el mapa guarda las listas aparte y necesita memoria adicional. Si cambian las ventas, la lista de esa categoría debe reordenarse.
