package common

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type (
	appError struct {
		Error      string `json:"error"`
		Message    string `json:"message"`
		HttpStatus int    `json:"status"`
	}
	errorResource struct {
		Data appError `json:"data"`
	}
	configuration struct {
		Server      string
		MongoDBHost string
		DBUser      string
		DBPwd       string
		Database    string
		LogLevel    int
	}
)

func DisplayAppError(c echo.Context, handleError error, message string, code int) {
	errObj := appError{
		Error:      handlerError.Error(),
		Message:    message,
		HttpStatus: code,
	}

	Error.Printf("[AppError]: %s\n", handlerError)
	if j, err := json.Marshal(errorResource{Data: errObj}); err == nil {
		w.Write(j)
	}
}

// AppConfig holds the configuration values from config.json file
var AppConfig configuration

// Initialize AppConfig
func initConfig() {
	loadAppConfig()
}

// Reads config.json and decode into AppConfig
func loadAppConfig() {
	file, err := os.Open("common/config.json")
	defer file.Close()

	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}

	decoder := json.NewDecoder(file)
	AppConfig = configuration{}
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("[loadAppConfig]: %s\n", err)
	}
}
