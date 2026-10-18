package services

import (
	"strconv"
)

// encargo es lo que se le pide al motor: el oficio que hace y lo que se le manda de verdad.
type encargo struct {
	sistema string
	usuario string
	texto   string
}

// encargoDe arma el encargo de uno de los dos campos.
//
// **Son dos encargos distintos, y por eso son dos peticiones** (docs/modules/ai.md, decisión 7): «de
// qué va el ticket» y «qué fue lo último que pasó» no se resumen igual, y un solo texto para las dos
// cosas saldría peor en las dos.
//
// El reparto es este, y está medido contra el motor de verdad (llama.cpp con Qwen2.5-1.5B-Instruct,
// 2026-09-27):
//
//   - el **mensaje de sistema** dice el oficio: de qué va el ticket, o qué fue lo último;
//   - el **mensaje de usuario** lleva el texto del ticket y, **al final**, la exigencia del formato:
//     el objeto JSON exacto con sus dos claves y nada más, con la plantilla delante de los ojos del
//     modelo. Va al final a propósito, porque es lo último que lee antes de escribir y es lo que más
//     se le pega a un modelo pequeño;
//   - y en el cuerpo de la petición va `response_format: {"type": "json_object"}`, que es la otra
//     mitad de lo que hace que el JSON salga limpio.
//
// El texto del ticket va recortado por el principio (decisión 6).
func encargoDe(tipo Tipo, texto string, palabras int, idioma string) encargo {
	contenido := recortarJerarquizado(texto)
	return encargo{
		sistema: oficioDe(tipo, idioma),
		usuario: contenido + "\n\n" + exigenciaDelFormato(tipo, palabras, idioma),
		texto:   contenido,
	}
}

// oficioDe es lo que se le pide a cada campo, que es lo único que cambia entre los dos encargos.
func oficioDe(tipo Tipo, idioma string) string {
	if idioma == "en" {
		switch tipo {
		case Summary:
			return "You summarize help desk tickets. State WHAT THE TICKET IS ABOUT in one direct present-tense sentence: what is happening to the requester and what they need. For a main ticket, focus on its subject and TARGET_REASON description. For an internal ticket, summarize its subject and TARGET_REASON escalation reason. Ignore later actions and the linked ticket's activity."
		case UltimaAccion:
			return "You summarize help desk tickets. State WHAT HAPPENED LAST in one direct past-tense sentence. The mandatory fact is identified by TARGET_LAST_EVENT, TARGET_LAST_ACTOR, TARGET_LAST_STATE and TARGET_LAST_COMMENT. Preserve the actor name and action; use the linked comment only to explain why. The linked ticket is context, never the target."
		}
	}
	switch tipo {
	case Summary:
		return "Eres quien resume tickets en una mesa de ayuda. Escribe DE QUÉ VA EL TICKET en una sola " +
			"frase, en presente y sin rodeos, como se lo contarías a un compañero que va a atenderlo: " +
			"qué le pasa a quien lo abrió y qué necesita. Para un ticket principal, usa su asunto y la " +
			"descripción TARGET_REASON. Para un ticket interno, resume su asunto y el motivo de escalado " +
			"TARGET_REASON. Ignora las acciones posteriores y la actividad del ticket enlazado."
	case UltimaAccion:
		return "Eres quien resume tickets en una mesa de ayuda. Escribe QUÉ FUE LO ÚLTIMO QUE PASÓ en " +
			"este ticket, en una sola frase, en pasado y sin rodeos. El hecho obligatorio está separado " +
			"como TARGET_LAST_EVENT, TARGET_LAST_ACTOR, TARGET_LAST_STATE y TARGET_LAST_COMMENT. Conserva " +
			"el nombre del actor y la acción; usa el comentario vinculado sólo para explicar el porqué. " +
			"El ticket enlazado aporta contexto, pero nunca sustituye al objetivo."
	}

	return ""
}

// exigenciaDelFormato es el contrato del módulo, y va en el mensaje del usuario: **un objeto JSON con
// las dos redacciones y ningún campo más** (docs/modules/ai.md, decisión 3), con el inglés escrito a
// la vez que el español y no traducido después (sección 7).
//
// La plantilla se le enseña tal cual —`{"es": "…", "en": "…"}`— porque es lo que hace que el modelo
// conteste ese objeto y no el suyo: se midió que sin ella se inventa las claves, y con ella sale bien
// incluso con un ticket de 6 000 caracteres.
func exigenciaDelFormato(tipo Tipo, palabras int, idioma string) string {
	if idioma == "en" {
		ejemplo := "A meaningful summary sentence."
		if tipo == UltimaAccion {
			ejemplo = "The requester opened the ticket and described the reported problem."
		}
		return "Do not invent anything absent from the ticket. Preserve every proper name exactly as written. Use at most " + strconv.Itoa(palabras) +
			` words. Return exactly {"text":"` + ejemplo + `"}; replace the example sentence with the summary and never return placeholders such as ..., — or N/A.`
	}
	ejemplo := "Una frase de resumen con contenido."
	if tipo == UltimaAccion {
		ejemplo = "El usuario abrió el ticket y describió el problema reportado."
	}
	return "No inventes nada que no esté en el ticket y conserva cada nombre propio exactamente como está escrito. " +
		"No pases de " + strconv.Itoa(palabras) + " palabras. " +
		`Devuelve exactamente {"text":"` + ejemplo + `"}; sustituye la frase de ejemplo por el resumen y nunca devuelvas marcadores como ..., — o N/A.`
}

func reforzarEncargo(encargo encargo, idioma string) encargo {
	refuerzo := "Your previous answer was invalid. Answer with a complete sentence containing letters and meaningful ticket information. Never answer with only punctuation, ..., — or N/A."
	if idioma == "es" {
		refuerzo = "Tu respuesta anterior no fue válida. Contesta con una frase completa que tenga letras e información real del ticket. Nunca respondas sólo con puntuación, ..., — o N/A."
	}
	encargo.usuario += "\n\n" + refuerzo
	return encargo
}
