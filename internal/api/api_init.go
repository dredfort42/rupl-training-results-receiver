package api

import (
	"os"
	db "training_results_receiver/internal/db"

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
	server.Host = cfg.Config["training.results.receiver.host"]
	if server.Host == "" {
		panic("training.results.receiver.host is not set")
	}

	server.Port = cfg.Config["training.results.receiver.port"]
	if server.Port == "" {
		panic("training.results.receiver.port is not set")
	}

	server.CorsStatus = cfg.Config["training.results.receiver.cors"]
	if server.CorsStatus == "" {
		loger.Warning("training.results.receiver.cors is not set | CORS is disabled")
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

	// Apply the middleware to the routes you want to protect
	authorized := router.Group("/", AuthMiddleware())
	{
		authorized.POST("/api/v1/training/result", TrainingResultCreate)
		// authorized.POST("/api/v1/profile/user", UserCreate)
		// authorized.GET("/api/v1/profile/user", UserGet)
		// authorized.PATCH("/api/v1/profile/user", UserUpdate)
		// authorized.DELETE("/api/v1/profile/user", UserDelete)
		// authorized.POST("/api/v1/profile/user/email", UserChangeEmail)
		// authorized.POST("/api/v1/profile/devices", DeviceCreate)
		// authorized.GET("/api/v1/profile/devices", DevicesGet)
		// authorized.PUT("/api/v1/profile/devices", DeviceUpdate)
		// authorized.DELETE("/api/v1/profile/devices", DeviceDelete)

		// authorized.POST("/api/v1/training/result", func() {
		// 	fmt.Ptintln("authorized.POST")
		// })
	}

	// // Unprotected route
	// router.GET("/unprotected", UnprotectedEndpoint)

	url := server.Host + ":" + server.Port
	loger.Success("Service successfully started", url)
	router.Run(url)

	db.DB.Database.Close()
}
