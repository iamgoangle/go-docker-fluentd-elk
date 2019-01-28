package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	logger "github.com/sirupsen/logrus"
)

type AppResponse struct {
	Project     string    `json:"projectName"`
	Environment string    `json:"environment"`
	Message     string    `json:"message"`
	Timestamp   time.Time `json:"timestamp"`
	Tag         string    `json:"tag"`
}

func init() {
	logger.SetFormatter(&logger.JSONFormatter{})
	logger.SetOutput(os.Stdout)
}

func main() {
	e := echo.New()
	// e.Use(middleware.Logger())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"method":"${method}","uri":"${uri}","status":"${status}"}`,
		// Format: `{"time":"${time_rfc3339_nano}", "tag":"golang.app", id":"${id}","remote_ip":"${remote_ip}","host":"${host}",` +
		// 	`"method":"${method}","uri":"${uri}","status":${status}, "latency":${latency},` +
		// 	`"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
		// 	`"bytes_out":${bytes_out}}` + "\n",
		Output: os.Stdout,
	}))

	e.GET("/test", func(c echo.Context) error {
		// Mock response data
		mockData := &AppResponse{
			Project:     "GOLF Project",
			Environment: "production",
			Message:     "This is test",
			Tag:         "golang.app",
			Timestamp:   time.Now().UTC(),
		}

		// Assume this code is HTTP Middleware
		var inInterface map[string]interface{}
		inrec, _ := json.Marshal(mockData)
		json.Unmarshal(inrec, &inInterface)
		logger.WithFields(inInterface).Info("GOLF TEST JA")

		return c.JSON(http.StatusOK, mockData)
	})

	e.Logger.Fatal(e.Start(":1324"))
}
