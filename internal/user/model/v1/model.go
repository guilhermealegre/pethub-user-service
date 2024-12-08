package v1

import (
	"github.com/google/uuid"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain"
	dCtx "github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain/context"
	"github.com/guilhermealegre/pethub-user-service/internal/user/domain/v1"
)

type Model struct {
	app        domain.IApp
	repository v1.IRepository
	streaming  v1.IStreaming
}

func NewModel(app domain.IApp, repository v1.IRepository, streaming v1.IStreaming) v1.IModel {
	return &Model{
		app:        app,
		repository: repository,
		streaming:  streaming,
	}
}

// Onboard user
func (m *Model) Onboard(ctx dCtx.IContext, idUser int, onboard *v1.Onboard) error {
	tx, err := m.app.Database().Write().Begin()
	if err != nil {
		return err
	}
	defer tx.RollbackUnlessCommitted()

	userProfile := &v1.UserProfile{
		FirstName: onboard.FirstName,
		LastName:  onboard.LastName,
		Avatar:    onboard.Avatar,
	}
	if err := m.repository.UpdateUserProfile(ctx, tx, idUser, userProfile); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

// CreateUser create new user
func (m *Model) CreateUser(ctx dCtx.IContext, uuidUser uuid.UUID) (idUser int, err error) {
	return m.repository.CreateUser(ctx, uuidUser)
}

// GetUserProfile get user profile
func (m *Model) GetUserProfile(ctx dCtx.IContext, userUUID uuid.UUID) (userProfile *v1.UserProfile, err error) {
	return m.repository.GetUserProfile(ctx, userUUID)
}

// UpdateUserProfile update user profile
func (m *Model) UpdateUserProfile(ctx dCtx.IContext, idUser int, profile *v1.UserProfile) error {

	return m.repository.UpdateUserProfile(ctx, nil, idUser, profile)
}

// GetUserMe get user me
func (m *Model) GetUserMe(ctx dCtx.IContext, userUUID uuid.UUID) (*v1.UserMe, error) {
	userProfile, err := m.repository.GetUserProfile(ctx, userUUID)
	if err != nil {
		return nil, err
	}

	return &v1.UserMe{
		UserUUID:   userProfile.UserUUID,
		UserId:     userProfile.UserID,
		Email:      userProfile.Email,
		FirstName:  userProfile.FirstName,
		LastName:   userProfile.LastName,
		OnboardSet: userProfile.OnboardSet,
	}, nil
}
