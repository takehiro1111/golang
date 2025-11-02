package web

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	ginMiddleware "github.com/oapi-codegen/gin-middleware"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag"
	"gorm.io/gorm"

	"go-api-arch-onion-template/application/service"
	"go-api-arch-onion-template/infrastructure/database/repository"
	"go-api-arch-onion-template/pkg"
	"go-api-arch-onion-template/pkg/logger"
	"go-api-arch-onion-template/userinterface/web/gin/api"
	"go-api-arch-onion-template/userinterface/web/gin/handler"
	"go-api-arch-onion-template/userinterface/web/gin/middleware"
)

func setupSwagger(router *gin.Engine) (*openapi3.T, error) {
	swagger, err := api.GetSwagger()
	if err != nil {
		return nil, err
	}

	env := pkg.GetEnvDefault("APP_ENV", "development")
	if env == "development" {
		swaggerJson, _ := json.Marshal(swagger)
		var SwaggerInfo = &swag.Spec{
			InfoInstanceName: "swagger",
			SwaggerTemplate:  string(swaggerJson),
		}
		swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
	return swagger, nil
}

func NewGinRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	router.Use(middleware.CorsMiddleware(NewConfigWeb().CorsAllowOrigins))
	swagger, err := setupSwagger(router)
	if err != nil {
		logger.Warn(err.Error())
	}

	router.Use(middleware.GinZap())
	router.Use(middleware.RecoveryWithZap())

	router.LoadHTMLGlob("./userinterface/web/html/*")
	router.GET("/", handler.Index)
	router.GET("/health", handler.Health)

	apiGroup := router.Group("/api")
	{
		apiGroup.Use(middleware.TimeoutMiddleware(2 * time.Second))
		v1 := apiGroup.Group("/v1")
		{
			v1.Use(ginMiddleware.OapiRequestValidator(swagger))
			albumRepository := repository.NewAlbumRepository(db)
			albumService := service.NewAlbumService(albumRepository)
			albumHandler := handler.NewAlbumHandler(albumService)
			api.RegisterHandlers(v1, albumHandler)
		}
	}
	return router
}

func StartGinServer(host, port string, db *gorm.DB) {
	router := NewGinRouter(db)
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", host, port),
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatal(err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")
	defer logger.Sync()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Error(fmt.Sprintf("Server Shutdown: %s", err.Error()))
	}
	<-ctx.Done()
}
