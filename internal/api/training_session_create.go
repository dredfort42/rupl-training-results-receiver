package api

import (
	"net/http"
	"strconv"
	"strings"
	db "training_sessions_receiver/internal/db"
	s "training_sessions_receiver/internal/structs"

	loger "github.com/dredfort42/tools/logprinter"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

// TrainingSessionCreate creates a training result
func TrainingSessionCreate(c *gin.Context) {
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

	var trainingResult s.TrainingSession
	if err := c.ShouldBindJSON(&trainingResult); err != nil {
		loger.Debug("Error binding JSON", err.Error())
		errorResponse.Error = "invalid_request"
		errorResponse.ErrorDescription = "Error binding JSON | " + err.Error()
		c.IndentedJSON(http.StatusBadRequest, errorResponse)
		return
	}

	if trainingResult.Session.Email != email {
		loger.Debug("Invalid email in session data")
		errorResponse.Error = "invalid_request"
		errorResponse.ErrorDescription = "Invalid email in session data"
		c.IndentedJSON(http.StatusBadRequest, errorResponse)
		return
	}

	if trainingResult.Session.StartTime == 0 || trainingResult.Session.EndTime == 0 || trainingResult.Session.StartTime > trainingResult.Session.EndTime {
		loger.Debug("Invalid session time")
		errorResponse.Error = "invalid_request"
		errorResponse.ErrorDescription = "Invalid session time"
		c.IndentedJSON(http.StatusBadRequest, errorResponse)
		return
	}

	dbTrainingResult, err := parseResult(trainingResult)
	if err != nil {
		loger.Debug("Error parsing session data", err.Error())
		errorResponse.Error = "invalid_request"
		errorResponse.ErrorDescription = "Error parsing session data | " + err.Error()
		c.IndentedJSON(http.StatusBadRequest, errorResponse)
		return
	}

	if err = db.TrainingSessionCreate(dbTrainingResult); err != nil {
		errorResponse.Error = "server_error"
		errorResponse.ErrorDescription = "Error creating session | " + err.Error()
		c.IndentedJSON(http.StatusInternalServerError, errorResponse)
		return
	}

	loger.Debug("Session created successfully for an ID: ", dbTrainingResult.SessionUUID)

	c.JSON(http.StatusCreated, gin.H{"message": "Training session created successfully", "session_uuid": dbTrainingResult.SessionUUID})
}

// int64ToString is a function to convert int64 to string
func int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

// float64ToString is a function to convert float64 to string
func float64ToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

// parseRouteData is a function to parse a route data
func parseRouteData(routeData []s.SessionRouteData) (dbRouteData pq.StringArray) {
	dbRouteData = make([]string, 0, len(routeData))

	for _, rd := range routeData {
		dbRouteData = append(dbRouteData, "("+
			int64ToString(rd.Timestamp)+", "+
			float64ToString(rd.Latitude)+", "+
			float64ToString(rd.Longitude)+", "+
			float64ToString(rd.HorizontalAccuracy)+", "+
			float64ToString(rd.Altitude)+", "+
			float64ToString(rd.VerticalAccuracy)+", "+
			float64ToString(rd.Speed)+", "+
			float64ToString(rd.SpeedAccuracy)+", "+
			float64ToString(rd.Course)+", "+
			float64ToString(rd.CourseAccuracy)+
			")")
	}

	return
}

// parseSessionDataInt is a function to parse a session data for int
func parseSessionDataInt(data []s.SessionQuantityData) (dbData pq.StringArray) {
	dbData = make([]string, 0, len(data))

	for _, d := range data {
		dbData = append(dbData, "("+
			int64ToString(d.Timestamp)+", "+
			strings.Split(strings.Split(d.Quantity, " ")[0], ".")[0]+
			")")
	}

	return
}

// parseSessionDataFloat32 is a function to parse a session data for int
func parseSessionDataFloat32(data []s.SessionQuantityData) (dbData pq.StringArray) {
	dbData = make([]string, 0, len(data))

	for _, d := range data {
		dbData = append(dbData, "("+
			int64ToString(d.Timestamp)+", "+
			strings.Split(d.Quantity, " ")[0]+
			")")
	}

	return
}

// parseResult is a function to parse a session
func parseResult(jsonData s.TrainingSession) (dbSession s.DBTrainingSession, err error) {
	dbSession = s.DBTrainingSession{
		SessionUUID:      uuid.New().String(),
		SessionStartTime: jsonData.Session.StartTime,
		SessionEndTime:   jsonData.Session.EndTime,
		Email:            jsonData.Session.Email,
		DeviceName:       jsonData.Session.DeviceName}

	dbSession.RouteData = parseRouteData(jsonData.RouteData)
	dbSession.StepCount = parseSessionDataInt(jsonData.StepCount)
	dbSession.RunningPower = parseSessionDataInt(jsonData.RunningPower)
	dbSession.VerticalOscillation = parseSessionDataFloat32(jsonData.VerticalOscillation)
	dbSession.EnergyBurned = parseSessionDataFloat32(jsonData.EnergyBurned)
	dbSession.HeartRate = parseSessionDataFloat32(jsonData.HeartRate)
	dbSession.StrideLength = parseSessionDataFloat32(jsonData.StrideLength)
	dbSession.GroundContactTime = parseSessionDataInt(jsonData.GroundContactTime)
	dbSession.Speed = parseSessionDataFloat32(jsonData.Speed)
	dbSession.Distance = parseSessionDataFloat32(jsonData.Distance)

	vo2max := float64(0)
	if len(jsonData.VO2Max) > 0 {
		vo2max, err = strconv.ParseFloat(strings.Split(jsonData.VO2Max[0].Quantity, " ")[0], 32)
		if err != nil {
			return
		}
	}
	dbSession.VO2MaxMLPerMinPerKg = float32(vo2max)

	return
}
