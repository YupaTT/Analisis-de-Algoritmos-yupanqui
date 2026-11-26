# H3: estrategia

## El problema

El reporte de ventas del [H1](../../h1/H1.md) se puede sacar por **mes** y por **sucursal**. El gerente quiere saber el **rendimiento mensual** de cada sucursal:

- Cada mes tiene un **rendimiento** = ventas del mes − ventas del mes anterior. Si las ventas subieron, es positivo; si bajaron, es negativo.
- Para cada sucursal, hay que encontrar la **mejor racha**: los meses seguidos `desde..hasta` en que más crecieron las ventas (la mayor suma de rendimientos).

Hay **3 sucursales**.

El reporte del H1 (filtrar, sumar y ordenar) no es un problema de optimización; por eso tomo una pregunta del mismo caso: la mejor racha de crecimiento.

## ¿Sirve el voraz? No

El voraz elige lo que parece mejor en el momento: toma el mejor mes y solo agranda el rango si la suma sube.

Ejemplo de una sucursal:

| Mes         | Dic | Ene | Feb | Mar | Abr |
| ----------- | --: | --: | --: | --: | --: |
| Ventas      | 100 | 106 | 105 | 110 | 100 |
| Rendimiento |   — |  +6 |  −1 |  +5 | −10 |

- **Voraz:** toma enero (+6). Agregar febrero baja la suma a +5, así que se detiene. Total: **+6**.
- **Mejor opción:** enero a marzo: 6 − 1 + 5 = **+10** (las ventas pasaron de 100 a 110).

Para llegar a marzo había que aceptar primero un mes malo. El voraz no lo acepta, así que no llega al óptimo.

Tampoco deja el mismo problema más chico: después de elegir enero, lo que queda no es «la mejor racha del resto», porque la racha tiene que seguir pegada a enero.

## ¿Sirve la programación dinámica? Sí

Cumple las dos condiciones:

1. **Las mismas cuentas se repiten.** La fuerza bruta suma enero–marzo, después enero–abril, después febrero–abril… y cada vez vuelve a sumar los mismos meses desde cero, igual que el H1 recalcula todo en cada consulta. La dinámica suma cada mes **una sola vez**.
2. **La respuesta se arma con respuestas más chicas.** La mejor racha que termina este mes es una de dos:
   - **empezar de nuevo:** solo este mes;
   - **seguir la racha:** la mejor racha que terminó el mes anterior, más este mes.

   Me quedo con la mayor. La respuesta del mes anterior ya estaba calculada, así que se recorre la lista una sola vez.
