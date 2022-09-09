package main

import (
	"context"
	"fmt"
	"master-user/middleware"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo"
	echoMid "github.com/labstack/echo/middleware"
)

func (s *Service) RoutingAndListen() {
	e := echo.New()
	e.Use(middleware.ServerHeader, middleware.Logger)
	e.Use(echoMid.Recover())
	e.Use(echoMid.CORS())
	apiV1 := e.Group("/api/v1")

	// Mounting Routing Users
	s.UserHandler.Mount(apiV1)

	// Start server
	listenerPort := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
	go func() {
		if err := e.Start(listenerPort); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
