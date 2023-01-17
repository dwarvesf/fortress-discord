package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"

	"github.com/dwarvesf/fortress-discord/pkg/adapter"
	"github.com/dwarvesf/fortress-discord/pkg/config"
	"github.com/dwarvesf/fortress-discord/pkg/discord"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
)

func main() {
	cfg := config.LoadConfig(config.DefaultConfigLoaders())
	l := logger.NewLogrusLogger()

	// init healthcheck routes
	router := setupRouter()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.ApiServer.Port),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			l.Fatal(err, "failed to listen and serve")
		}
	}()

	// init discord session
	ses, err := discordgo.New("Bot " + cfg.Discord.SecretToken)
	if err != nil {
		l.Fatal(err, "failed to create discord session")
	}

	adapter := adapter.New(cfg, l)

	discord := discord.New(ses, cfg, l, service.New(adapter, l), view.New(ses))
	session, err := discord.ListenAndServe()
	if err != nil {
		l.Fatal(err, "failed to listen and serve discord")
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	session.Close()
	srv.Shutdown(context.TODO())
}

func setupRouter() *gin.Engine {
	r := gin.New()
	r.GET("/healthz", func(c *gin.Context) {
		c.Header("Content-Type", "text/plain")
		c.String(http.StatusOK, "OK")
	})
	return r
}
