# C5: torneo de ordenamientos

Se comparan Inserción, MergeSort y QuickSort con **10.000, 20.000, 40.000 y 80.000 elementos**.
Protocolo: **un calentamiento, 3 corridas y promedio**.
[Resultados](resultados.txt) en milisegundos:

| Cantidad | Inserción (ms) | MergeSort (ms) | QuickSort (ms) |
| -------: | -------------: | -------------: | -------------: |
|   10.000 |      47.127309 |       0.990322 |       0.168521 |
|   20.000 |     256.327853 |       3.426887 |       0.702519 |
|   40.000 |    1370.412899 |       9.145732 |       1.779423 |
|   80.000 |    3293.205414 |      10.012564 |       2.169780 |

**Conclusión:** QuickSort fue el más rápido, seguido de MergeSort. Inserción fue
la más lenta. Estos tiempos corresponden a listas en orden inverso y dependen
del equipo.
