package main

import (
	"github.com/jgroeneveld/configurate"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

var Config struct {
	OpenAIKey string `env:"OPENAI_API_KEY"`
}

func main() {
	configureLogger()
	log.Info().Msg("Starting server...")

	loadConfig()

	r := configureRouter()
	r.POST("/translate", translateTextHandler)

	err := r.Run("localhost:8080")
	if err != nil {
		log.Err(err).Msg("Error starting server")
	}
}

func configureLogger() {
	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	// human output for development
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Error loading .env file")
	}

	err = configurate.LoadAll(&Config, configurate.NewEnvLoader(), configurate.NewDefaultsLoader(), configurate.NewRequiredLoader())
	if err != nil {
		log.Fatal().Err(err).Msg("Error loading config")
	}
}

func configureRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(GinLogger())
	r.Use(gin.Recovery())
	return r
}

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		log.Info().
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Int("status", c.Writer.Status()).
			Dur("duration", time.Since(start)).
			Str("ip", c.ClientIP()).
			Int("size", c.Writer.Size()).
			Msg("translateTextRequest handled")
	}
}
