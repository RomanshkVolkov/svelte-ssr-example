package repository

import (
	"fmt"
	"time"

	"github.com/RomanshkVolkov/svelte-and-go-template/internal/core/lg"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dsn = GetEnv("DB_POSTGRE_STRING_CONECTION", "not-provider")
var DATABASE *gorm.DB

func DBConnection() {
	loadEnv()
	if dsn == "not-provider" {
		dsn = GetEnv("DB_POSTGRE_STRING_CONECTION", "not-provider")
	}
	var db *gorm.DB
	var err error
	// Intentar reconectar hasta 10 veces con espera de 5 segundos
	maxRetries := 10
	for i := 0; i < maxRetries; i++ {
		lg.Info("Intentando conectar a la base de datos...")
		lg.Info(dsn)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			lg.Info("✅ Conectado a la base de datos exitosamente")
			break
		}

		lg.Error(fmt.Sprintf("❌ Intento %d fallido: %s", i+1, err.Error()))
		if i == maxRetries-1 {
			lg.Error(fmt.Sprintf("❌ No se pudo conectar a la base de datos después de %d intentos", maxRetries))
			panic("failed to connect database")
		}

		// Esperar antes del siguiente intento
		time.Sleep(5 * time.Second)
	}

	DATABASE = db

	StartSeeding(db)
}
