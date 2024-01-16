package handler

import (
	"net/http"
	"path"

	"github.com/asfarsharief/money_management_backend/lib/config"
	"github.com/labstack/echo/v4"
)

// CommonHandler - Handler to handle basic apis like version, ping and api docs.
type CommonHandler struct {
	*config.ServerConfiguration
}

// Health godoc
// @Summary Health server
// @Description API endpoints that will check whether server is up and running.
// @Accept  json
// @Produce  json
// @Success 200
// @Router /api/v1/health [get]
// Health - This function will ping the echo server
func (s *CommonHandler) Health(context echo.Context) error {
	return context.JSON(http.StatusOK, map[string]interface{}{"health": "ok"})
}

// Swagger godoc
// @Summary Swagger for good documentation
// @Description APIs Documentation endpoint.
// @Accept  json
// @Produce  json
// @Success 200
// @Router /api/v1/swagger [get]
// Swagger method will display the swagger documentation
func (s *CommonHandler) Swagger(context echo.Context) error {
	return context.Render(http.StatusOK, "swagger.html", struct{ PublicPath string }{
		path.Join(s.APIPath, "public"),
	})
}
