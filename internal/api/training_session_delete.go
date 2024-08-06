package api

import (
	"net/http"
	db "training_sessions_receiver/internal/db"
	s "training_sessions_receiver/internal/structs"

	loger "github.com/dredfort42/tools/logprinter"
	"github.com/gin-gonic/gin"
)

// TrainingSessionDelete deletes a training result
func TrainingSessionDelete(c *gin.Context) {
	var errorResponse s.ResponseError

	tmpEmail, exists := c.Get("email")
	if !exists || tmpEmail.(string) == "" {
		loger.Debug("Missing email")
		errorResponse.Error = "invalid_request"
		errorResponse.ErrorDescription = "Missing email"
		c.IndentedJSON(http.StatusInternalServerError, errorResponse)
		return
	}

	email := tmpEmail.(string)

	sessionUUID := c.Request.URL.Query().Get("session_uuid")
	if sessionUUID == "" {
		loger.Debug("Missing session_uuid")
		errorResponse.Error = "invalid_request"
		errorResponse.ErrorDescription = "Missing session_uuid"
		c.IndentedJSON(http.StatusBadRequest, errorResponse)
		return
	}

	if result := db.TrainingSessionExistsCheckByUUID(sessionUUID, email); !result {
		loger.Debug("Training session does not exist")
		errorResponse.Error = "not_found"
		errorResponse.ErrorDescription = "Training session not exist"
		c.IndentedJSON(http.StatusNotFound, errorResponse)
		return
	}

	err := db.TrainingSessionDelete(sessionUUID, email)
	if err != nil {
		loger.Debug("Error deleting training session", err.Error())
		errorResponse.Error = "internal_error"
		errorResponse.ErrorDescription = "Error deleting training session | " + err.Error()
		c.IndentedJSON(http.StatusInternalServerError, errorResponse)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Training session deleted successfully"})
}
