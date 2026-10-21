package middleware

import (
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"catalina-support/backend/shared/auth"
	"catalina-support/backend/shared/httpx"
)

type limitEntry struct {
	used  int64
	reset time.Time
}

// Limiter es un contador en memoria por proceso. La demo usa una sola réplica; reiniciar sólo
// vuelve el contador más permisivo y no afecta ningún dato.
type Limiter struct {
	mu      sync.Mutex
	entries map[string]limitEntry
}

func NewLimiter() *Limiter { return &Limiter{entries: make(map[string]limitEntry)} }

type LimitKey func(*http.Request) string
type LimitCost func(*http.Request) int64

func (l *Limiter) Middleware(prefix string, maximum int64, window time.Duration, key LimitKey, cost LimitCost) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		if maximum <= 0 {
			return next
		}
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			value := key(r)
			units := int64(1)
			if cost != nil {
				units = cost(r)
			}
			if units < 1 {
				units = 1
			}
			now := time.Now()
			l.mu.Lock()
			entry := l.entries[prefix+":"+value]
			if entry.reset.IsZero() || !now.Before(entry.reset) {
				entry = limitEntry{reset: now.Add(window)}
			}
			allowed := entry.used+units <= maximum
			if allowed {
				entry.used += units
			}
			l.entries[prefix+":"+value] = entry
			l.mu.Unlock()
			if !allowed {
				seconds := int64(time.Until(entry.reset).Seconds())
				if seconds < 1 {
					seconds = 1
				}
				w.Header().Set("Retry-After", strconv.FormatInt(seconds, 10))
				httpx.WriteError(w, http.StatusTooManyRequests, "rateLimit.exceeded")
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func IPKey(r *http.Request) string {
	if value := strings.TrimSpace(r.Header.Get("X-Real-IP")); value != "" {
		return value
	}
	if value := r.Header.Get("X-Forwarded-For"); value != "" {
		return strings.TrimSpace(strings.Split(value, ",")[0])
	}
	return r.RemoteAddr
}

func AccountKey(r *http.Request) string       { return auth.MustFromContext(r.Context()).Subject }
func ContentLengthCost(r *http.Request) int64 { return r.ContentLength }
