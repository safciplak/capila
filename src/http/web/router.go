package web

import (
	"context"
	"net/http"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmgin/v2"

	helpers "github.com/safciplak/capila/src/helpers/environment"
	middlewares "github.com/safciplak/capila/src/http/middleware"
	"github.com/safciplak/capila/src/logger"
)

// Router initializes the Gin Router and the common middleware
func Router(log logger.InterfaceLogger, environment helpers.InterfaceEnvironmentHelper) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	// If it's not explicitly disabled, for instance for local debugging purposes
	if environment.Get("DISABLE_ROUTER_LOGGING") == "" {
		router.Use(middlewares.LogMiddleware(log))
	}

	if environment.Error() != nil {
		log.Log(context.Background()).Error("Warning! " + environment.Error().Error())
	}

	// Add APM infused middlewares
	router.Use(ginzap.RecoveryWithZap(log.GetZapLogger(), true))
	router.Use(apmgin.Middleware(router))

	// Add standardized response handling in case of a Panic
	router.Use(middlewares.PanicMiddleware())

	// Ping endpoint that gets included in every microservice endpoint
	router.GET("/v1/health/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	return router
}
