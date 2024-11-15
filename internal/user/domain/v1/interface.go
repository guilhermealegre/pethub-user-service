package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gocraft/dbr/v2"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain"
)

type IController interface {
	domain.IController
	GetUserProfile(ctx *gin.Context)
	UpdateUserProfile(ctx *gin.Context)
	GetUserMe(ctx *gin.Context)
}

type IModel interface {
	CreateUser(ctx domain.IContext, tx *dbr.Tx, userProfile *User) (int, error)
	GetUserProfile(ctx domain.IContext, idUser int) (*UserProfile, error)
	UpdateUserProfile(ctx domain.IContext, idUser int, profile *UserProfile) error
	Onboard(ctx domain.IContext, idUser int, onboard *Onboard) error
	GetUserMe(ctx domain.IContext, idUser int) (*UserMe, error)
}

type IRepository interface {
	CreateUser(ctx domain.IContext, tx *dbr.Tx, userProfile *User) (int, error)
	GetUserProfile(ctx domain.IContext, idUser int) (*UserProfile, error)
	UpdateUserProfile(ctx domain.IContext, tx dbr.SessionRunner, idUser int, profile *UserProfile) error
}

type IStreaming interface{}
