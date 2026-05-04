package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	httpRoutes "github.com/RomanshkVolkov/svelte-and-go-template/internal/adapters/http"
	"github.com/RomanshkVolkov/svelte-and-go-template/internal/adapters/middleware"
	"github.com/RomanshkVolkov/svelte-and-go-template/internal/core/repository"
)

// Config estructura de configuración de la aplicación
type Config struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

func main() {
	// Configuración del servidor
	config := &Config{
		Port:         repository.GetEnv("PORT", "8080"),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	repository.DBConnection()

	// Inicializar rutas con inyección de dependencias
	routes := httpRoutes.InitRoutes(repository.DATABASE)

	// Aplicar middlewares globales
	handler := middleware.Chain(
		routes,
		middleware.Logger,
		middleware.CORS,
		middleware.Recovery,
	)

	// Configurar servidor HTTP
	server := &http.Server{
		Addr:         ":" + config.Port,
		Handler:      handler,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
		IdleTimeout:  config.IdleTimeout,
	}

	// Canal para manejar señales del sistema
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Goroutine para iniciar el servidor
	go func() {
		log.Printf("🚀 Server starting on port %s", config.Port)
		log.Printf("📍 Health check: http://localhost:%s/health", config.Port)
		log.Printf("📖 API documentation: http://localhost:%s/api/routes", config.Port)
		log.Printf("👥 Users endpoint: http://localhost:%s/api/v1/users", config.Port)

		printBanner()

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ Server failed to start: %v", err)
		}

	}()

	// Esperar señal de interrupción
	<-sigChan

	log.Println("🔄 Shutting down server...")

	// Contexto con timeout para el graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Intentar graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("❌ Server forced to shutdown: %v", err)
	}

	log.Println("✅ Server gracefully stopped")
}

// printBanner imprime un banner de inicio (opcional)
func printBanner() {
	banner := `
╭─────────────────────────────────────────────────────╮
│                                                     │
│  🔥 Svelte + Go Template API Server                 │
│                                                     │
│  📚 Clean Architecture with Dependency Injection    │
│  🏗️  Repository Pattern                              │
│  🔄 Graceful Shutdown                               │
│  🛡️  Middleware Support                              │
│                                                     │
╰─────────────────────────────────────────────────────╯
`
	fmt.Print(banner)
}
