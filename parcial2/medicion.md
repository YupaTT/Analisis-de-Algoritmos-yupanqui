Gonzalo Yupanqui Torrez

# P2.3 — Medición

- **Protocolo:** 1 corrida de calentamiento (no cuenta) + 3 corridas medidas + promedio.
- **Tamaños:** n = 10.000, 20.000, 40.000 (se duplican). Elegí 10.000 porque con 25.000 la ingenua tardaba ~35 s en el tamaño mayor (más de 20 s). Salida completa en `resultados.txt`.

## Ingenuo (lista simple): declarado O(n²) para la carga

| n      | Corrida 1 (ms) | Corrida 2 (ms) | Corrida 3 (ms) | Promedio (ms) | Veces respecto de la fila anterior |
| ------ | -------------- | -------------- | -------------- | ------------- | ---------------------------------- |
| 10.000 | 474,02         | 358,28         | 378,19         | 403,50        | —                                  |
| 20.000 | 1.118,10       | 800,60         | 1.114,60       | 1.011,10      | ×2,51                              |
| 40.000 | 3.815,90       | 3.710,23       | 3.574,81       | 3.700,32      | ×3,66                              |

## Mejorado (5 colas + mapa): declarado O(n) para la carga

| n      | Corrida 1 (ms) | Corrida 2 (ms) | Corrida 3 (ms) | Promedio (ms) | Veces respecto de la fila anterior |
| ------ | -------------- | -------------- | -------------- | ------------- | ---------------------------------- |
| 10.000 | 8,76           | 6,59           | 10,61          | 8,65          | —                                  |
| 20.000 | 12,23          | 15,20          | 17,97          | 15,13         | ×1,75                              |
| 40.000 | 29,91          | 27,32          | 29,84          | 29,02         | ×1,92                              |

## Conclusión

- **Ingenuo:** en el último salto crece ×3,66, cerca del ×4 de O(n²). El ×2,51 de la fila de 20.000 es ruido: sus corridas van de 800 a 1.118 ms (~30 % de diferencia) y la de 10.000 quedó alta por una corrida de 474 ms.
- **Mejorado:** crece ×1,75 y ×1,92, cerca del ×2 de O(n) (cada operación cuesta O(1)). Con tiempos de pocos ms el ruido pesa mucho (en 10.000 van de 6,6 a 10,6 ms).
- En 40.000 el mejorado es unas **127 veces más rápido**, y la diferencia crece al duplicar n, como predicen O(n²) contra O(n).
