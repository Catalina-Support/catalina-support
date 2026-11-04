# Banco de pruebas de modelos locales

Este banco mide **modelos locales** para los resúmenes de tickets sin tocar el motor activo ni
producción. Manda a cada modelo **el mismo mensaje de sistema y de usuario que arma la aplicación**
(`backend/modules/ai/services/encargos.go` + `pedirGlobalUnaVez`), con la misma temperatura, el mismo
tope de piezas y `response_format: {"type":"json_object"}`.

## Qué mide

Cada respuesta se contrasta con el texto que recibió el modelo:

- JSON válido y clave `text`.
- Idioma correcto.
- Tope de 40 palabras.
- Sin marcadores vacíos (`...`, `—`, `N/A`).
- **Hechos esperados** del caso (`facts`), definidos en `cases.json`.
- **Nombres propios y números inventados**: cualquier nombre o cifra de la respuesta que no aparezca
  en el texto entregado se marca.

El resultado por caso queda en `results/<modelo>.json`; la salida de consola resume JSON, idioma,
palabras, latencia y veredicto.

## Uso

```powershell
# Todos los modelos (descarga ~4 GB la primera vez)
pwsh -File tests/ai-benchmark/run.ps1

# Sólo algunos, con más hilos
pwsh -File tests/ai-benchmark/run.ps1 -Only qwen2.5-1.5b -Threads 8 -Suffix '-t8'

# Qwen3 sin razonamiento (variante exploratoria; la aplicación no lo envía así)
pwsh -File tests/ai-benchmark/run.ps1 -Only qwen3-1.7b -ThinkingOff -Suffix '-nothink'
```

Requiere Docker. Los modelos viven en el volumen `catalina_support_ai_bench_models`, separado del
volumen del motor real. El servidor de pruebas escucha en `127.0.0.1:18080` y se detiene al terminar
cada modelo.

El informe legible se regenera con:

```powershell
docker run --rm -v "${PWD}/tests/ai-benchmark:/bench" -w /bench python:3.12-slim python report.py
```

y queda en `report.md`.

## Casos

`cases.json` contiene ocho casos: resumen y última acción, en español e inglés, principal e interno.
Cada caso declara los hechos que una respuesta correcta debe contener. No se evalúa el estilo: eso
queda para la revisión humana del informe.

## Modelos

- `qwen2.5-1.5b`: el modelo actual, como referencia.
- `qwen2.5-0.5b`: la alternativa mínima sin razonamiento.
- `qwen3-0.6b`: el candidato más rápido.
- `qwen3-1.7b`: el candidato equilibrado.
- `gemma3-1b`: alternativa multilingüe.

Los Qwen3 se prueban **tal como los enviaría la aplicación hoy**, sin añadir `enable_thinking=false`.
La variante `-ThinkingOff` sirve para saber si el problema es el modelo o el modo razonamiento; no es
el contrato actual.

## Resultados

Cada respuesta se guarda en `results/<modelo>.json` (ignorado por Git) y el informe consolidado en
`report.md`, con el texto exacto, la latencia, los hechos, los nombres inventados y una nota de
revisión.

La conclusión de la primera pasada está en `report.md`: el modelo actual sigue siendo el mejor
equilibrio; subir de 2 a 8 hilos reduce la latencia a la mitad sin cambiar la calidad, y ningún
modelo probado gana a la vez en velocidad y contenido.
