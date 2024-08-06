package api

import (
	"os"
	db "training_sessions_receiver/internal/db"

	cfg "github.com/dredfort42/tools/configreader"
	loger "github.com/dredfort42/tools/logprinter"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Host               string
	Port               string
	CorsStatus         string
	AuthServerURL      string
	IdentifyPathUser   string
	IdentifyPathDevice string
	ChangePathEmail    string
	DeletePathUser     string
}

var server Server

// ApiInit starts the trainig results receiver service
func ApiInit() {
	server.Host = cfg.Config["training.sessions.receiver.host"]
	if server.Host == "" {
		panic("training.sessions.receiver.host is not set")
	}

	server.Port = cfg.Config["training.sessions.receiver.port"]
	if server.Port == "" {
		panic("training.sessions.receiver.port is not set")
	}

	server.CorsStatus = cfg.Config["training.sessions.receiver.cors"]
	if server.CorsStatus == "" {
		loger.Warning("training.sessions.receiver.cors is not set | CORS is disabled")
		server.CorsStatus = "false"
	}

	server.AuthServerURL = cfg.Config["auth.url"]
	if server.AuthServerURL == "" {
		panic("auth.url is not set")
	}

	server.IdentifyPathUser = cfg.Config["auth.path.identify.user"]
	if server.IdentifyPathUser == "" {
		panic("auth.path.identify.user is not set")
	}

	server.IdentifyPathDevice = cfg.Config["auth.path.identify.device"]
	if server.IdentifyPathDevice == "" {
		panic("auth.path.identify.device is not set")
	}

	if os.Getenv("DEBUG") != "true" && os.Getenv("DEBUG") != "1" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.Default()

	if server.CorsStatus == "true" || server.CorsStatus == "1" {
		router.Use(cors.Default())
	}

	authorized := router.Group("/", AuthMiddleware())
	{
		authorized.POST("/api/v1/training/session", TrainingSessionCreate)
		authorized.DELETE("/api/v1/training/session", TrainingSessionDelete)
	}

	url := server.Host + ":" + server.Port
	loger.Success("Service successfully started", url)
	router.Run(url)

	db.DB.Database.Close()
}
