package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

type translateTextRequest struct {
	Text           string `json:"text"`
	TargetLanguage string `json:"target_language"`
}

func translateTextHandler(c *gin.Context) {
	var request translateTextRequest
	if err := c.BindJSON(&request); err != nil {
		log.Err(err).Msg("Invalid request")
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	log.Debug().Msgf("Fetching translation: %s to %s", request.Text, request.TargetLanguage)

	err, translation := fetchTranslation(request)
	if err != nil {
		log.Err(err).Msg("ChatCompletion error")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid request " + err.Error()})
		return
	}

	log.Debug().Msg("Translation: " + translation)

	c.JSON(http.StatusOK, gin.H{"translation": translation})
}

func fetchTranslation(request translateTextRequest) (error, string) {
	prompt := buildPrompt(request)
	log.Debug().Msg("Prompt: " + prompt)

	resp, err := fetchChatCompletion(prompt)
	if err != nil {
		return err, ""
	}

	translation := resp.Choices[0].Message.Content
	return nil, translation
}

func buildPrompt(request translateTextRequest) string {
	return fmt.Sprintf("Translate the following text into %s:\n%s", request.TargetLanguage, request.Text)
}
