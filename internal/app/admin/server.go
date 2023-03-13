package admin

import (
	"fmt"
	"net/http"

	_ "github.com/Speakerkfm/social_network_otus/internal/app/admin/statik"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rakyll/statik/fs"
)

func New() (*http.Server, error) {
	engine := gin.Default()

	engine.Use(cors.Default())

	statikFS, err := fs.New()
	if err != nil {
		return nil, fmt.Errorf("fail to create statik fs: %w", err)
	}
	engine.StaticFS("/docs", statikFS)

	return &http.Server{
		Handler: engine,
		Addr:    "0.0.0.0:84",
	}, nil
}
