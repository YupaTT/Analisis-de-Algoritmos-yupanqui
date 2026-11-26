Gonzalo Yupanqui Torrez

# Conclusión — El reporte del H1, antes y después

El proceso es el del [H1](../h1/H1.md): los productos más vendidos en un rango de fechas, sobre unas 120.000 líneas por mes, 8 sucursales y 8.000 productos.

## Qué cambié y por qué

| Paso | Antes | Ahora | Por qué |
| --- | --- | --- | --- |
| Hallar las fechas | Recorrer todas las líneas, O(n) | **Búsqueda binaria**, O(log n) | Las líneas se ordenan por fecha una vez al cargar y cada consulta salta directo al rango ([c4](../practicas/c4/README.md)). |
| Sumar por producto | Buscar en una lista, O(p) por línea | **Mapa** producto → total, O(1) | Va directo al producto ([estructuras](../practicas/estructuras/README.md)). |
| Ordenar | Inserción, O(p²) | **MergeSort**, O(p log p) | Rápido y deja los empates en el mismo orden que hoy ([c5](../practicas/c5/README.md), [c6](../practicas/c6/eleccion.md)). |

El programa verifica que las dos versiones den **el mismo reporte** (productos, cantidades, montos y orden) antes de medir.

## Antes y después

| Reporte | Antes (ms) | Después (ms) | Veces más rápido |
| --- | ---: | ---: | ---: |
| Un mes, 120.000 líneas | 864,16 | 8,82 | ×98 |
| Un año, 1.440.000 líneas | 11.151,79 | 51,74 | ×216 |

Todas las tablas en [medicion.md](medicion.md). Para correrlo: `go run .` dentro de `conclusion/`.
