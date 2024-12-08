package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gocraft/dbr/v2"
	"github.com/google/uuid"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain"
	dCtx "github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain/context"
)

type IController interface {
	domain.IController
	GetUserProfile(ctx *gin.Context)
	UpdateUserProfile(ctx *gin.Context)
	GetUserMe(ctx *gin.Context)
}

type IModel interface {
	CreateUser(ctx dCtx.IContext, uuid2 uuid.UUID) (int, error)
	GetUserProfile(ctx dCtx.IContext, idUser uuid.UUID) (*UserProfile, error)
	UpdateUserProfile(ctx dCtx.IContext, idUser int, profile *UserProfile) error
	Onboard(ctx dCtx.IContext, idUser int, onboard *Onboard) error
	GetUserMe(ctx dCtx.IContext, userUUID uuid.UUID) (*UserMe, error)
}

type IRepository interface {
	CreateUser(ctx dCtx.IContext, uuid uuid.UUID) (int, error)
	GetUserProfile(ctx dCtx.IContext, uuid uuid.UUID) (*UserProfile, error)
	UpdateUserProfile(ctx dCtx.IContext, tx dbr.SessionRunner, idUser int, profile *UserProfile) error
}

type IStreaming interface{}
