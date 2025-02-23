package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/rogeriofontes/cert-generator/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Config contém todas as configurações da aplicação
type Config struct {
	DatabaseDSN    string
	SMTPUser       string
	SMTPPass       string
	OutputDir      string
	BackgroundPath string
}

// LoadEnv carrega variáveis de ambiente do arquivo .env
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("⚠️  Aviso: Não foi possível carregar o arquivo .env, utilizando variáveis de ambiente do sistema.")
	}
}

// LoadConfig carrega as configurações do sistema
func LoadConfig() *Config {
	LoadEnv()
	return &Config{
		DatabaseDSN: fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable client_encoding=UTF8",
			getEnv("DB_HOST", "localhost"),
			getEnv("DB_USER", "postgres"),
			getEnv("DB_PASSWORD", "postgres"),
			getEnv("DB_NAME", "certificados"),
			getEnv("DB_PORT", "5432"),
		),
		SMTPUser:       getEnv("SMTP_USER", ""),
		SMTPPass:       getEnv("SMTP_PASS", ""),
		OutputDir:      getEnv("OUTPUT_DIR", "output"),
		BackgroundPath: getEnv("BACKGROUND_PATH", "backgrounds/certificado.jpg"),
	}
}

// InitDB inicializa a conexão com o banco de dados
func InitDB() {
	cfg := LoadConfig()

	db, err := gorm.Open(postgres.Open(cfg.DatabaseDSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Erro ao conectar ao banco de dados: %v", err)
	}

	// 🔹 Faz a migração automática das tabelas
	db.AutoMigrate(&domain.Community{}, &domain.Event{}, &domain.Participant{}, &domain.User{})

	// 🔹 Garante que o banco use UTF-8 corretamente
	db.Exec("SET client_encoding = 'UTF8'")
	db.Exec("SET NAMES 'UTF8'")
	db.Exec("SET standard_conforming_strings = on")
	db.Exec("SET bytea_output = 'hex'")

	log.Println("✅ Conectado ao banco de dados com sucesso!")
	DB = db
}

// getEnv retorna o valor de uma variável de ambiente ou um valor padrão se não existir
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
