package constants

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

type Config struct {
	Port          string
	Env           string
	ProjectID     string
	GcsBucketName string
	DbHost        string
	DbUser        string
	DbPassword    string
	DbName        string
	DbPort        string
}

var projectDirName = "quizard-backend"

func init() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func New() *Config {
	return &Config{
		DbHost:     getEnv("DB_HOST", ""),
		DbUser:     getEnv("DB_USER", ""),
		DbPassword: getEnv("DB_PASSWORD", ""),
		DbName:     getEnv("DB_NAME", ""),
		DbPort:     getEnv("DB_PORT", ""),
		Port:       getEnv("PORT", ""),
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
