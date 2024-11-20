package v1

import (
	"context"
	"github.com/google/uuid"
	dCtx "github.com/guilhermealegre/go-clean-arch-infrastructure-lib/context"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain"
	"github.com/guilhermealegre/pethub-user-service/api/v1/grpc/user_service_user"
	v1 "github.com/guilhermealegre/pethub-user-service/internal/user/domain/v1"
)

type GrpcController struct {
	*domain.DefaultController
	user_service_user.UnimplementedUserServer
	model v1.IModel
}

func NewGrpcController(app domain.IApp, model v1.IModel) domain.IController {
	return &GrpcController{
		DefaultController: domain.NewDefaultController(app),
		model:             model,
	}
}

func (c *GrpcController) Register() {
	server, _ := c.App().Grpc().GetServer()
	user_service_user.RegisterUserServer(server, c)
}

func (c *GrpcController) GetUser(ctx context.Context, request *user_service_user.GetUserRequest) (*user_service_user.GetUserResponse, error) {
	/*
		newCtx := &dCtx.Context{}
		resp, err := c.model.GetUser((&dCtx.Context{}).FromGrpc(ctx), request.Id)
		if err != nil {
			return nil, err
		}

		return resp.ToApi(), err

	*/
	return nil, nil
}

func (c *GrpcController) CreateUser(ctx context.Context, request *user_service_user.CreateUserRequest) (*user_service_user.CreateUserResponse, error) {
	newCtx := (&dCtx.Context{}).FromGrpc(ctx)

	uuidUser, err := uuid.FromBytes(request.UUID)
	if err != nil {
		return nil, err
	}

	idUser, err := c.model.CreateUser(newCtx, uuidUser)
	if err != nil {
		return nil, err
	}

	return &user_service_user.CreateUserResponse{Id: int32(idUser)}, err
}
