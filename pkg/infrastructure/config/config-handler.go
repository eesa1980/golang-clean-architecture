package confighandler

import (
	"clean-architecture/pkg/application/common/structs"
	"github.com/joho/godotenv"
	"os"
)

func New() structs.ApplicationConfig {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	appName := os.Getenv("APP_NAME")
	baseUrl := os.Getenv("BASE_URL")
	port := os.Getenv("PORT")

	return structs.ApplicationConfig{
		AppName: appName,
		BaseUrl: baseUrl,
		Port:    port,
	}
}
