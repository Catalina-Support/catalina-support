#!/usr/bin/env python3
"""Banco de pruebas de modelos locales para Catalina Support.

Envía a cada modelo exactamente el mensaje de sistema y de usuario que arma la aplicación
(`encargos.go` + `pedirGlobalUnaVez`), con la misma temperatura, tope de piezas y
`response_format`, y evalúa cada respuesta contra el texto que recibió el modelo.

No habla con la base de datos ni toca el motor activo: sólo manda peticiones HTTP a un
`llama-server` de pruebas.
"""

from __future__ import annotations

import argparse
import json
import re
import time
import unicodedata
import urllib.error
import urllib.request
from pathlib import Path


SYSTEM = {
    ("summary", "es"): (
        "Eres quien resume tickets en una mesa de ayuda. Escribe DE QUÉ VA EL TICKET en una sola "
        "frase, en presente y sin rodeos, como se lo contarías a un compañero que va a atenderlo: "
        "qué le pasa a quien lo abrió y qué necesita. Para un ticket principal, usa su asunto y la "
        "descripción TARGET_REASON. Para un ticket interno, resume su asunto y el motivo de escalado "
        "TARGET_REASON. Ignora las acciones posteriores y la actividad del ticket enlazado."
    ),
    ("last_action", "es"): (
        "Eres quien resume tickets en una mesa de ayuda. Escribe QUÉ FUE LO ÚLTIMO QUE PASÓ en "
        "este ticket, en una sola frase, en pasado y sin rodeos. El hecho obligatorio está separado "
        "como TARGET_LAST_EVENT, TARGET_LAST_ACTOR, TARGET_LAST_STATE y TARGET_LAST_COMMENT. Conserva "
        "el nombre del actor y la acción; usa el comentario vinculado sólo para explicar el porqué. "
        "El ticket enlazado aporta contexto, pero nunca sustituye al objetivo."
    ),
    ("summary", "en"): (
        "You summarize help desk tickets. State WHAT THE TICKET IS ABOUT in one direct present-tense "
        "sentence: what is happening to the requester and what they need. For a main ticket, focus on "
        "its subject and TARGET_REASON description. For an internal ticket, summarize its subject and "
        "TARGET_REASON escalation reason. Ignore later actions and the linked ticket's activity."
    ),
    ("last_action", "en"): (
        "You summarize help desk tickets. State WHAT HAPPENED LAST in one direct past-tense sentence. "
        "The mandatory fact is identified by TARGET_LAST_EVENT, TARGET_LAST_ACTOR, TARGET_LAST_STATE "
        "and TARGET_LAST_COMMENT. Preserve the actor name and action; use the linked comment only to "
        "explain why. The linked ticket is context, never the target."
    ),
}

ES_WORDS = {
    "el", "la", "los", "las", "de", "del", "que", "y", "en", "un", "una", "para", "por", "con",
    "se", "al", "es", "ha", "su", "sus", "no", "lo", "como", "más", "pero", "sin", "ya", "está",
}
EN_WORDS = {
    "the", "a", "an", "of", "to", "in", "and", "for", "with", "is", "was", "has", "on", "that",
    "it", "this", "from", "after", "before", "when", "user", "requester", "ticket", "support",
}
SENTENCE_START = {
    "el", "la", "los", "las", "un", "una", "se", "al", "no", "en", "the", "a", "an", "this",
    "that", "it", "support", "soporte", "desarrollo", "dev", "user", "requester", "finally",
    "entonces", "después", "tras", "cuando",
}


def normalize(value: str) -> str:
    return "".join(
        ch for ch in unicodedata.normalize("NFD", value.lower()) if unicodedata.category(ch) != "Mn"
    )


def language_of(text: str) -> str:
    words = re.findall(r"[a-záéíóúñü]+", text.lower())
    es = sum(1 for word in words if word in ES_WORDS)
    en = sum(1 for word in words if word in EN_WORDS)
    if es > en:
        return "es"
    if en > es:
        return "en"
    return "desconocido"


def first_json(content: str):
    try:
        return json.loads(content)
    except json.JSONDecodeError:
        pass

    start = content.find("{")
    end = content.rfind("}")
    if start == -1 or end <= start:
        return None
    try:
        return json.loads(content[start : end + 1])
    except json.JSONDecodeError:
        return None


def grounded_tokens(text: str, source: str) -> list[str]:
    """Nombres propios y números de la respuesta que no están en el texto entregado."""
    allowed = normalize(source)
    candidatos = re.findall(r"\b[A-ZÁÉÍÓÚÑ][a-záéíóúñü]{2,}\b", text)
    candidatos += re.findall(r"\b\d{2,}\b", text)
    fuera: list[str] = []
    for token in candidatos:
        if normalize(token) in SENTENCE_START:
            continue
        if normalize(token) not in allowed:
            fuera.append(token)
    return sorted(set(fuera))


def check_case(caso: dict, body: dict, latency_ms: int) -> dict:
    contenido = (body.get("choices") or [{}])[0].get("message", {}).get("content", "")
    if isinstance(contenido, list):
        contenido = "".join(part.get("text", "") for part in contenido if isinstance(part, dict))

    parsed = first_json(contenido)
    texto = ""
    if isinstance(parsed, dict) and isinstance(parsed.get("text"), str):
        texto = parsed["text"].strip()

    normalized_text = normalize(texto)
    facts = []
    for grupo in caso["facts"]:
        facts.append(any(normalize(alt) in normalized_text for alt in grupo))
    facts_ok = all(facts)

    no_placeholders = normalized_text not in {"", "n/a", "na"} and not re.fullmatch(
        r"[\s.…—\-_/]*", texto
    )
    language = language_of(texto)
    words = len(texto.split())
    inventados = grounded_tokens(texto, "\n".join(caso["text"]))

    contract_ok = bool(
        parsed
        and texto
        and no_placeholders
        and language == caso["language"]
        and words <= 40
    )
    content_ok = bool(facts_ok and not inventados)

    return {
        "id": caso["id"],
        "kind": caso["kind"],
        "language_expected": caso["language"],
        "language_detected": language,
        "json_valid": parsed is not None,
        "text": texto,
        "raw_content": contenido,
        "words": words,
        "facts": facts,
        "facts_ok": facts_ok,
        "invented": inventados,
        "grounding_ok": not inventados,
        "no_placeholders": no_placeholders,
        "contract_ok": contract_ok,
        "content_ok": content_ok,
        "ok": contract_ok and content_ok,
        "latency_ms": latency_ms,
    }


def call(base_url: str, model: str, caso: dict, timeout: int) -> tuple[dict, int]:
    text = "\n".join(caso["text"])
    idioma = "Spanish" if caso["language"] == "es" else "English"
    user = (
        text
        + "\n\nWrite only in "
        + idioma
        + '. Return exactly {"text":"..."} and no other keys or surrounding text.'
    )
    payload = {
        "model": model,
        "messages": [
            {"role": "system", "content": SYSTEM[(caso["kind"], caso["language"])]},
            {"role": "user", "content": user},
        ],
        "temperature": 0.2,
        "max_tokens": 320,
        "response_format": {"type": "json_object"},
    }
    request = urllib.request.Request(
        base_url.rstrip("/") + "/v1/chat/completions",
        data=json.dumps(payload).encode("utf-8"),
        headers={"Content-Type": "application/json"},
        method="POST",
    )
    start = time.perf_counter()
    with urllib.request.urlopen(request, timeout=timeout) as response:
        body = json.loads(response.read().decode("utf-8"))
    return body, int((time.perf_counter() - start) * 1000)


def main() -> int:
    parser = argparse.ArgumentParser()
    parser.add_argument("--base-url", required=True)
    parser.add_argument("--model", required=True)
    parser.add_argument("--cases", required=True)
    parser.add_argument("--out", required=True)
    parser.add_argument("--timeout", type=int, default=180)
    args = parser.parse_args()

    datos = json.loads(Path(args.cases).read_text(encoding="utf-8"))
    resultados = []
    for caso in datos["cases"]:
        try:
            body, latency = call(args.base_url, args.model, caso, args.timeout)
            resultados.append(check_case(caso, body, latency))
        except (urllib.error.URLError, TimeoutError, json.JSONDecodeError) as error:
            resultados.append(
                {
                    "id": caso["id"],
                    "kind": caso["kind"],
                    "language_expected": caso["language"],
                    "error": str(error),
                    "ok": False,
                }
            )

    salida = {"model": args.model, "cases": resultados}
    Path(args.out).parent.mkdir(parents=True, exist_ok=True)
    Path(args.out).write_text(json.dumps(salida, ensure_ascii=False, indent=2), encoding="utf-8")

    print(f"== {args.model}")
    print(f"{'caso':32} {'ok':4} {'json':4} {'idioma':7} {'palabras':8} {'latencia':8}")
    for item in resultados:
        print(
            f"{item['id']:32} {str(item.get('ok', False)):4} "
            f"{str(item.get('json_valid', False)):4} "
            f"{item.get('language_detected', '?'):7} {item.get('words', 0):<8} "
            f"{item.get('latency_ms', 0):>6} ms"
        )
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
