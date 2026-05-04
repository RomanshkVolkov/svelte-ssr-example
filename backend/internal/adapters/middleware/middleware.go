package middleware

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"time"
)

// Middleware tipo para definir middlewares
type Middleware func(http.Handler) http.Handler

// Chain combina múltiples middlewares en uno solo
func Chain(h http.Handler, middlewares ...Middleware) http.Handler {
	// Aplicar middlewares en orden inverso para que se ejecuten en el orden correcto
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}

// Logger middleware para logging de requests
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Crear un ResponseWriter que capture el status code
		ww := &responseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		// Procesar el request
		next.ServeHTTP(ww, r)

		// Log del request
		duration := time.Since(start)
		log.Printf(
			"[%s] %s %s - %d - %v - %s",
			r.Method,
			r.RequestURI,
			r.RemoteAddr,
			ww.statusCode,
			duration,
			r.UserAgent(),
		)
	})
}

// CORS middleware para manejar Cross-Origin Resource Sharing
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Configurar headers CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Max-Age", "86400")

		// Manejar preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Procesar el request normal
		next.ServeHTTP(w, r)
	})
}

// Recovery middleware para capturar panics y convertirlos en errores HTTP
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// Log del panic
				log.Printf("PANIC: %v\n%s", err, debug.Stack())

				// Responder con error 500
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)

				errorResponse := fmt.Sprintf(`{
					"success": false,
					"message": "Internal server error",
					"error": "An unexpected error occurred"
				}`)

				w.Write([]byte(errorResponse))
			}
		}()

		next.ServeHTTP(w, r)
	})
}

// RateLimiter middleware básico para rate limiting (implementación simple)
func RateLimiter(requestsPerMinute int) Middleware {
	// En una implementación real usarías redis o memoria compartida
	// Esta es una implementación muy básica solo para demostrar el concepto
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Por ahora solo pasamos al siguiente handler
			// En una implementación real implementarías la lógica de rate limiting
			next.ServeHTTP(w, r)
		})
	}
}

// ContentType middleware para asegurar Content-Type JSON en endpoints API
func ContentTypeJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Solo aplicar a métodos que envían datos
		if r.Method == "POST" || r.Method == "PUT" || r.Method == "PATCH" {
			contentType := r.Header.Get("Content-Type")
			if contentType != "" && contentType != "application/json" {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnsupportedMediaType)
				w.Write([]byte(`{
					"success": false,
					"message": "Content-Type must be application/json",
					"error": "Unsupported media type"
				}`))
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

// Security middleware para headers de seguridad básicos
func Security(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Headers de seguridad básicos
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")

		next.ServeHTTP(w, r)
	})
}

// responseWriter wrapper para capturar el status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader captura el status code
func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
