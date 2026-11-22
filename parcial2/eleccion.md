# Gonzalo Yupanqui Torrez

# P2.1 — Elección de estructuras de datos

## 1. Estructuras elegidas

| Operación        | Cantidad diaria | Estructura utilizada           | Costo         |
| ---------------- | --------------: | ------------------------------ | ------------- |
| Registrar        |         100.000 | Colas FIFO por urgencia + mapa | O(1) promedio |
| AtenderSiguiente |         100.000 | Cinco colas FIFO               | O(1)          |
| EstaEsperando    |          30.000 | Mapa (map)                     | O(1) promedio |
| Esperando        |        Variable | Contador                       | O(1)          |

Se utilizan cinco colas FIFO, una por cada nivel de urgencia (1 a 5). Los pacientes se agregan al final de su cola y se atienden desde el principio, respetando el orden de llegada cuando tienen la misma urgencia.

También se utiliza un mapa para consultar si un paciente sigue esperando y un contador para conocer la cantidad total de pacientes pendientes.

## 2. ¿Por qué no utilizar una lista simple?

Con una lista simple, para atender al paciente más urgente es necesario recorrer los elementos para encontrarlo. Además, consultar si un paciente está esperando requiere otra búsqueda lineal.

Con 100.000 pacientes:

- Registrar 100.000 pacientes: O(n) total.
- Atender 100.000 pacientes: O(n²) en el peor caso.
- Consultar 30.000 pacientes: O(30.000 × 100.000) en el peor caso.

En cambio, con las estructuras elegidas:

- Registrar 100.000 pacientes: O(100.000) promedio.
- Atender 100.000 pacientes: O(100.000).
- Consultar 30.000 pacientes: O(30.000) promedio.

La solución mejorada evita recorrer todos los pacientes para encontrar al siguiente y para comprobar si alguien sigue esperando.

## 3. ¿Cuándo cambiaría de estructura?

Si el sistema tuviera una cantidad muy grande o variable de niveles de urgencia, en lugar de solamente cinco, consideraría utilizar una cola de prioridad (Heap), porque permitiría administrar las prioridades sin depender de un número fijo de colas.
