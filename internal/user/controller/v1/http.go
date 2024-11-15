package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/context"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain"
	v1Routes "github.com/guilhermealegre/pethub-user-service/api/v1/http"
	"github.com/guilhermealegre/pethub-user-service/api/v1/http/envelope/request"
	"github.com/guilhermealegre/pethub-user-service/api/v1/http/envelope/response"
	v1 "github.com/guilhermealegre/pethub-user-service/internal/user/domain/v1"
)

type Controller struct {
	*domain.DefaultController
	model v1.IModel
}

func NewController(app domain.IApp, model v1.IModel) v1.IController {
	return &Controller{
		DefaultController: domain.NewDefaultController(app),
		model:             model,
	}
}

func (c *Controller) Register() {
	engine := c.App().Http().Router()
	v1Routes.OnboardUser.SetRoute(engine, c.Onboard)
	v1Routes.GetUserProfile.SetRoute(engine, c.GetUserProfile)
	v1Routes.UpdateUserProfile.SetRoute(engine, c.UpdateUserProfile)
	v1Routes.GetUserMe.SetRoute(engine, c.GetUserMe)
}

func (c *Controller) Onboard(gCtx *gin.Context) {
	ctx := context.NewContext(gCtx)
	var req request.OnboardRequest

	if err := ctx.ShouldBindJSON(&req.Body); err != nil {
		c.Json(ctx, nil, err)
		return
	}

	if err := c.App().Validator().Validate(ctx, req); err != nil {
		c.Json(ctx, nil, err)
		return
	}
	onboard := &v1.Onboard{}
	onboard.FromAPIToDomain(&req)
	err := c.model.Onboard(ctx, ctx.GetUser().Id, onboard)

	c.Json(ctx, response.SuccessResponse{Success: err == nil}, err)

}

func (c *Controller) GetUserProfile(gCtx *gin.Context) {
	ctx := context.NewContext(gCtx)

	obj, err := c.model.GetUserProfile(ctx, ctx.GetUser().Id)

	c.Json(ctx, obj.FromDomainToAPI(), err)

	return
}

func (c *Controller) UpdateUserProfile(gCtx *gin.Context) {
	ctx := context.NewContext(gCtx)
	req := request.UpdateUserProfileRequest{
		IdUser: ctx.GetUser().Id,
	}

	if err := ctx.ShouldBindJSON(&req.Body); err != nil {
		c.Json(ctx, nil, err)
		return
	}

	if err := c.App().Validator().Validate(ctx, &req); err != nil {
		c.Json(ctx, nil, err)
		return
	}

	err := c.model.UpdateUserProfile(ctx, ctx.GetUser().Id, (&v1.UserProfile{}).FromAPIToDomain(&req))

	c.Json(ctx, response.SuccessResponse{Success: err == nil}, err)
}

func (c *Controller) GetUserMe(gCtx *gin.Context) {
	ctx := context.NewContext(gCtx)

	userMe, err := c.model.GetUserMe(ctx, ctx.GetUser().Id)

	c.Json(ctx, userMe.FromDomainToAPI(), err)

}
