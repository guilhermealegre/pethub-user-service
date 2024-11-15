package v1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/go-openapi/runtime/middleware"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain"
	v1Routes "github.com/guilhermealegre/pethub-user-service/api/v1/http"

	"github.com/gin-gonic/gin"
)

const v1 = "v1"

type Controller struct {
	domain.IController
	app domain.IApp
}

func NewController(app domain.IApp) domain.IController {
	return &Controller{
		app: app,
	}
}

func (c *Controller) Register() {
	v1Int, _ := strconv.Atoi(strings.TrimPrefix(v1, "v"))

	v1Router := c.app.Http().Router().Group("v1")
	v1Router.StaticFile(c.StaticFile(v1Int))
	v1Routes.SwaggerUserSwagger.SetRoute(c.app.Http().Router(), c.Swagger(v1Int))
	v1Routes.SwaggerUserDocs.SetRoute(c.app.Http().Router(), c.Docs(v1Int))
}

/*
	 swagger:route GET /user/docs swagger docs

	 Get swagger docs

		Produces:
		- text/html

		Responses:
		  200:
		  400: ErrorResponse
*/
func (c *Controller) Docs(version int) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		gin.WrapH(middleware.Redoc(
			middleware.RedocOpts{
				Path:    fmt.Sprintf("v%d/p/documentation/user/docs", version),
				SpecURL: fmt.Sprintf("/v%d/p/documentation/user/swagger.json", version),
			}, nil))(ctx)
	}
}

/*
	 swagger:route GET /user/swagger swagger swagger

	 Get swagger

		Produces:
		- text/html

		Responses:
		  200:
		  400: ErrorResponse
*/
func (c *Controller) Swagger(version int) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		gin.WrapH(middleware.SwaggerUI(
			middleware.SwaggerUIOpts{
				Path:    fmt.Sprintf("v%d/p/documentation/user/swagger", version),
				SpecURL: fmt.Sprintf("/v%d/p/documentation/user/swagger.json", version),
			}, nil))(ctx)
	}
}

func (c *Controller) StaticFile(version int) (relativePath, filePath string) {
	return "/p/documentation/user/swagger.json", fmt.Sprintf("internal/swagger/docs/v%d/swagger.json", version)
}
