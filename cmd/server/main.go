package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/dwarvesf/fortress-discord/pkg/config"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
)

func main() {
	cfg := config.LoadConfig(config.DefaultConfigLoaders())
	log := logger.NewLogrusLogger()

	// init discord session

	// init healthcheck routes
	router := setupRouter()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.ApiServer.Port),
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err, "failed to listen and serve")
		}
	}()
}

func setupRouter() *gin.Engine {
	r := gin.New()
	r.GET("/healthz", func(c *gin.Context) {
		c.Header("Content-Type", "text/plain")
		c.String(http.StatusOK, "OK")
	})
	return r
}
