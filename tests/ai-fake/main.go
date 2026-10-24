// Servicio determinista para las pruebas E2E. Implementa la frontera compatible con OpenAI que
// usa Catalina Support y responde según el encargo, sin red ni estado compartido.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type request struct {
	System   string    `json:"system"`
	Messages []message `json:"messages"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, map[string]string{"status": "ok"})
	})
	mux.HandleFunc("POST /v1/chat/completions", complete)
	mux.HandleFunc("POST /v1/messages", complete)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func complete(w http.ResponseWriter, r *http.Request) {
	var in request
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, 1<<20)).Decode(&in); err != nil {
		http.Error(w, `{"error":"invalid request"}`, http.StatusBadRequest)
		return
	}

	system, user := in.System, ""
	for _, item := range in.Messages {
		if item.Role == "system" {
			system = item.Content
		}
		if item.Role == "user" {
			user = item.Content
		}
	}

	text := responseFor(system, user)
	content, _ := json.Marshal(map[string]string{"text": text})
	if strings.HasSuffix(r.URL.Path, "/messages") {
		writeJSON(w, map[string]any{"content": []map[string]string{{"text": string(content)}}})
		return
	}
	writeJSON(w, map[string]any{"choices": []map[string]any{{"message": map[string]string{"content": string(content)}}}})
}

func responseFor(system, user string) string {
	lower := strings.ToLower(system)
	if strings.Contains(lower, "reescribe el mensaje") || strings.Contains(lower, "rewrite the supplied") {
		return "Hemos revisado la información indicada y necesitamos confirmar el mensaje exacto que aparece para continuar con el diagnóstico. Por favor, envíanos ese detalle y avísanos si el problema vuelve a presentarse."
	}
	if strings.Contains(lower, "qué fue lo último") || strings.Contains(lower, "what happened last") {
		actor := marker(user, "TARGET_LAST_ACTOR")
		event := marker(user, "TARGET_LAST_EVENT")
		comment := marker(user, "TARGET_LAST_COMMENT")
		verb := map[string]string{
			"cerrado": "cerró el ticket", "resuelto": "resolvió el ticket",
			"reabierto": "reabrió el ticket", "asignado": "actualizó la asignación",
			"creado": "abrió el ticket", "estado": "cambió el estado del ticket",
			"comentario": "añadió un comentario", "editado": "actualizó el ticket",
			"prioridad": "cambió la prioridad del ticket", "escalado": "escaló el ticket",
			"devuelto": "devolvió el ticket", "observador_anadido": "añadió un observador",
			"observador": "quitó un observador",
		}[strings.ToLower(event)]
		if actor == "" {
			actor = "El equipo"
		}
		if verb == "" {
			verb = "actualizó el ticket"
		}
		if comment != "" {
			return fmt.Sprintf("%s %s: %s", actor, verb, comment)
		}
		return fmt.Sprintf("%s %s.", actor, verb)
	}
	if reason := marker(user, "TARGET_REASON"); reason != "" {
		return reason
	}
	return "La persona solicitante necesita ayuda para resolver el problema descrito en el ticket."
}

func marker(text, name string) string {
	prefix := "[" + name + "]"
	for _, line := range strings.Split(text, "\n") {
		if strings.HasPrefix(strings.TrimSpace(line), prefix) {
			return strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(line), prefix))
		}
	}
	return ""
}

func writeJSON(w http.ResponseWriter, value any) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(value)
}
