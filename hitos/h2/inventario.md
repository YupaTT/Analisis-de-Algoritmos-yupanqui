# H2 — Inventario: productos más vendidos por categoría

El encargado elige una categoría y el sistema le muestra los 5 productos más vendidos de esa categoría.

## Datos

- **1.000 productos** repartidos en **6 categorías** (unos 167 por categoría).
- Cada producto tiene su **ID**, su **categoría** y su **cantidad vendida**.
- Los datos son simulados.

## Operaciones

| Operación                                          | ¿La hago? | Frecuencia                                                    |
| -------------------------------------------------- | --------- | ------------------------------------------------------------- |
| Buscar por clave (los productos de una categoría)  | Sí        | **La más frecuente**: una vez por consulta                    |
| Consultar los mayores (top 5 por vendidos)         | Sí        | **Muy frecuente**: una vez por consulta                       |
| Recorrer en orden (por vendidos, de mayor a menor) | Sí        | En cada consulta, solo los primeros 5                         |
| Insertar                                           | Sí        | n veces, solo al cargar los productos                         |
| Borrar                                             | No        | Durante la consulta no se borran productos                    |
| Rango / el siguiente                               | No        | No se piden rangos de vendidos ni el siguiente de un producto |

Las operaciones que más se repiten son **buscar por categoría** y **tomar los más vendidos**. Ambas ocurren en cada consulta, y para medir supongo **10.000 consultas**. Insertar ocurre una sola vez por producto.
