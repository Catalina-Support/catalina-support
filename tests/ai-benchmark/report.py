#!/usr/bin/env python3
"""Genera el informe legible del banco de pruebas a partir de results/*.json."""

from __future__ import annotations

import json
from pathlib import Path
from statistics import mean

BASE = Path(__file__).resolve().parent
RESULTS = BASE / "results"
OUT = BASE / "report.md"

ORDEN = [
    ("qwen2.5-1.5b", "Qwen2.5-1.5B · 2 hilos (actual)"),
    ("qwen2.5-1.5b-t4", "Qwen2.5-1.5B · 4 hilos"),
    ("qwen2.5-1.5b-t8", "Qwen2.5-1.5B · 8 hilos"),
    ("qwen2.5-1.5b-t16", "Qwen2.5-1.5B · 16 hilos"),
    ("qwen2.5-1.5b-t32", "Qwen2.5-1.5B · 32 hilos"),
    ("gemma3-1b", "Gemma 3 1B · 2 hilos"),
    ("qwen3-0.6b", "Qwen3-0.6B · razonamiento por defecto"),
    ("qwen3-0.6b-nothink", "Qwen3-0.6B · razonamiento desactivado"),
    ("qwen3-1.7b", "Qwen3-1.7B · razonamiento por defecto"),
    ("qwen3-1.7b-nothink", "Qwen3-1.7B · razonamiento desactivado"),
    ("qwen2.5-0.5b", "Qwen2.5-0.5B · 2 hilos"),
]

NOTAS = {
    ("qwen2.5-1.5b-t8", "en-last-internal-resolved"): "Repite el motivo en vez de la última acción (Dev One actualizó la configuración y reinició el servicio).",
    ("qwen2.5-1.5b", "en-last-internal-resolved"): "La frase es correcta; el evaluador no aceptaba «resolve» como «resolved».",
    ("gemma3-1b", "es-summary-password"): "No menciona la contraseña; dice «errores de autenticación».",
    ("gemma3-1b", "es-summary-printer"): "Inventa que se «solicitó una impresora nueva», que no está en el texto.",
    ("gemma3-1b", "en-last-internal-resolved"): "Atribuye la acción a «Dev One team»; el evaluador no acepta «resolve» como «resolved».",
    ("gemma3-1b", "en-last-user-reply"): "Resume bien; el evaluador no acepta «attempted» como «tried».",
    ("qwen3-0.6b", "es-summary-password"): "JSON truncado: se queda a mitad de frase.",
    ("qwen3-0.6b", "en-last-escalation"): "JSON truncado antes de cerrar.",
    ("qwen3-0.6b-nothink", "es-summary-password"): "Responde con una pregunta del comentario, no con un resumen.",
    ("qwen3-0.6b-nothink", "en-summary-vpn"): "Devuelve sólo el número y el estado; no resume.",
    ("qwen2.5-0.5b", "es-last-router"): "Repite el motivo («La conexión se corta»), no la última acción.",
    ("qwen2.5-0.5b", "en-last-escalation"): "Repite el motivo, no la escalación.",
    ("qwen2.5-0.5b", "en-last-internal-resolved"): "Repite el motivo, no la última acción.",
    ("qwen2.5-0.5b", "en-last-user-reply"): "Repite el motivo, no la respuesta del usuario.",
    ("qwen3-1.7b", "es-last-router"): "Sin contenido: el razonamiento agota el tope y no llega a escribir.",
    ("qwen3-1.7b", "en-last-escalation"): "Sin contenido: el razonamiento agota el tope y no llega a escribir.",
    ("qwen3-1.7b", "es-summary-internal"): "Sin contenido: el razonamiento agota el tope y no llega a escribir.",
    ("qwen3-1.7b", "en-last-internal-resolved"): "Sin contenido: el razonamiento agota el tope y no llega a escribir.",
    ("qwen3-1.7b", "en-last-user-reply"): "Sin contenido: el razonamiento agota el tope y no llega a escribir.",
}


def main() -> int:
    lineas = [
        "# Banco de pruebas de modelos locales",
        "",
        "Este informe se genera desde `results/*.json`. Cada respuesta se contrasta con el texto que",
        "recibió el modelo: JSON, idioma, palabras, hechos esperados, nombres/números inventados y",
        "revisión humana de la frase.",
        "",
        "## Resumen",
        "",
        "| Modelo | Casos | Correctos | Latencia media | Memoria del motor |",
        "| --- | ---: | ---: | ---: | ---: |",
    ]

    memorias = {
        "qwen2.5-1.5b": "1,191 GiB",
        "qwen2.5-0.5b": "493 MiB",
        "qwen3-0.6b": "717 MiB",
        "qwen3-1.7b": "1,576 GiB",
        "gemma3-1b": "880 MiB",
    }

    datos = []
    for archivo, etiqueta in ORDEN:
        ruta = RESULTS / f"{archivo}.json"
        if not ruta.exists():
            continue
        datos.append((archivo, etiqueta, json.loads(ruta.read_text(encoding="utf-8"))))

    for archivo, etiqueta, data in datos:
        casos = data["cases"]
        correctos = sum(1 for caso in casos if caso.get("ok"))
        latencias = [caso["latency_ms"] for caso in casos if "latency_ms" in caso]
        memoria = memorias.get(archivo.split("-t")[0], "—")
        media = f"{mean(latencias):.0f} ms" if latencias else "—"
        lineas.append(f"| {etiqueta} | {len(casos)} | {correctos}/{len(casos)} | {media} | {memoria} |")

    lineas += ["", "## Cada respuesta", ""]

    casos_entrada = {
        caso["id"]: caso
        for caso in json.loads((BASE / "cases.json").read_text(encoding="utf-8"))["cases"]
    }

    lineas += [
        "## Entradas entregadas al modelo",
        "",
        "El texto siguiente es el que recibió cada modelo, además del mensaje de sistema y de la",
        "exigencia de formato. Se reproduce para poder contrastar cada respuesta con su entrada.",
        "",
    ]
    for caso in casos_entrada.values():
        lineas += [f"### {caso['id']}", "", "```text", "\n".join(caso["text"]), "```", ""]

    for archivo, etiqueta, data in datos:
        lineas += [f"### {etiqueta}", ""]
        lineas += [
            "| Caso | Lat. | Pal. | JSON | Idioma | Hechos | Inventado | Respuesta / nota |",
            "| --- | ---: | ---: | :---: | :---: | :---: | :---: | --- |",
        ]
        for caso in data["cases"]:
            if "error" in caso:
                lineas.append(
                    f"| {caso['id']} | — | — | — | — | — | — | Error: {caso['error']} |"
                )
                continue
            nota = NOTAS.get((archivo, caso["id"]), "")
            texto = caso.get("text") or "(sin texto)"
            if nota:
                texto = f"{texto} — **{nota}**"
            texto = texto.replace("|", "\\|").replace("\n", " ")
            hechos = caso.get("facts", [])
            hechos_txt = f"{sum(hechos)}/{len(hechos)}" if hechos else "—"
            inventado = ", ".join(caso.get("invented", [])) or "—"
            lineas.append(
                f"| {caso['id']} | {caso.get('latency_ms', 0)} ms | {caso.get('words', 0)} | "
                f"{'sí' if caso.get('json_valid') else 'no'} | "
                f"{caso.get('language_detected', '?')} | {hechos_txt} | {inventado} | {texto} |"
            )
        lineas.append("")

    lineas += [
        "## Lectura",
        "",
        "- **Qwen2.5-1.5B con 2 hilos** es el mejor equilibrio actual: 7/8 correctos y 3,4-4,9 s.",
        "- **Subir a 8 hilos** baja la latencia a 1,2-2,4 s sin cambiar el modelo ni el prompt. La",
        "  máquina tiene 32 núcleos y el motor usa 2; es la mejora de velocidad más clara y con menos",
        "  riesgo. 16 y 32 hilos no mejoran: el cuello deja de ser la CPU.",
        "- **Gemma 3 1B** es válida en formato y ocupa menos memoria, pero no es más rápida y comete",
        "  errores de contenido (la impresora «nueva»).",
        "- **Qwen3-0.6B** con razonamiento por defecto tarda 10-13 s y trunca el JSON; sin razonamiento",
        "  baja a 1-1,6 s pero falla casi todos los casos por contenido e idioma.",
        "- **Qwen3-1.7B** con razonamiento por defecto tarda 25-29 s y deja respuestas vacías; sin",
        "  razonamiento mejora, pero ocupa 1,576 GiB (por encima del techo de 1,5 GB) y sigue fallando",
        "  la última acción en dos casos.",
        "- **Qwen2.5-0.5B** es rapidísimo (1,2-2,3 s) y cabe en 493 MiB, pero en los casos de última",
        "  acción repite el motivo y no cuenta lo que pasó.",
        "",
        "**Conclusión:** no hay un cambio de modelo que sea a la vez más rápido y de calidad igual o",
        "mejor. La mejora recomendada es mantener Qwen2.5-1.5B y subir los hilos de 2 a 8, midiendo en",
        "desarrollo antes de tocar producción. Si se quiere explorar un modelo, el siguiente paso es",
        "Qwen3-1.7B sin razonamiento y con el techo de memoria en 2 GiB, pero necesita más pruebas.",
    ]

    OUT.write_text("\n".join(lineas) + "\n", encoding="utf-8")
    print(f"Informe: {OUT}")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
