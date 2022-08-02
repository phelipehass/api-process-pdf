package config

import (
	"github.com/apex/log"
	"github.com/joho/godotenv"
	"os"
)

func LoadEnv() {
	envFileName := ".env"
	info, _ := os.Stat(envFileName)
	if info != nil {

		log.Infof("[LoadEnv] - Load environment from file %s", envFileName)

		err := godotenv.Load(envFileName)
		if err != nil {
			log.Fatalf("[LoadEnv] - Error loading .env file: %s", err.Error())
		}
	} else {
		log.Infof("[LoadEnv] - Using environment from OS")
	}
}

func ApiPort() string {
	env, exist := os.LookupEnv("API_PORT")

	if exist {
		return env
	}

	return "3000"
}

func GetInitExtraction() string {
	return getEnvironmentFromName("INIT_EXTRACTION")
}

func GetFinishExtraction() string {
	return getEnvironmentFromName("FINISH_EXTRACTION")
}

func GetTitle() string {
	return getEnvironmentFromName("TITLE")
}

func GetSubtitle() string {
	return getEnvironmentFromName("SUBTITLE")
}

func GetURLConsultDiaries() string {
	return getEnvironmentFromName("URL_CONSULT_DIARIO")
}

func GetURLBaseConsult() string {
	return getEnvironmentFromName("URL_CONSULT_DIARIO")
}

func GetCookie() string {
	return getEnvironmentFromName("COOKIE")
}

func getEnvironmentFromName(name string) (env string) {
	env = os.Getenv(name)
	if env == "" {
		log.Fatalf("[getEnvironmentFromName] - %s não encontrado", name)
	}

	return
}
