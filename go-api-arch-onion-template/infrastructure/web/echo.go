package web

import (
	"context"
	"errors"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"go-api-arch-onion-template/pkg/logger"
	"go-api-arch-onion-template/userinterface/web/echo/handler"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewEchoRouter(db *gorm.DB) *echo.Echo {
	router := echo.New()
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("./userinterface/web/html/*")),
	}
	router.Renderer = renderer
	router.GET("/", handler.Index)
	router.GET("/health", handler.Health)
	return router
}

func StartEchoServer(host, port string, db *gorm.DB) {
	router := NewEchoRouter(db)

	go func() {
		if err := router.Start(fmt.Sprintf("%s:%s", host, port)); err != nil && !errors.Is(err, http.ErrServerClosed) {
			router.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")
	defer logger.Sync()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := router.Shutdown(ctx); err != nil {
		logger.Error(fmt.Sprintf("Server Shutdown: %s", err.Error()))
	}
	<-ctx.Done()
}
