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
func (m *Model) GetUserProfile(ctx dCtx.IContext, idUser int) (userProfile *v1.UserProfile, err error) {
	userProfile, err = m.repository.GetUserProfile(ctx, idUser)
	if err != nil {
		return nil, err
	}

	return userProfile, nil
}

// UpdateUserProfile update user profile
func (m *Model) UpdateUserProfile(ctx dCtx.IContext, idUser int, profile *v1.UserProfile) error {

	if err := m.repository.UpdateUserProfile(ctx, nil, idUser, profile); err != nil {
		return m.app.Logger().DBLog(err)
	}

	return nil
}

// GetUserMe get user me
func (m *Model) GetUserMe(ctx dCtx.IContext, idUser int) (*v1.UserMe, error) {
	userProfile, err := m.repository.GetUserProfile(ctx, idUser)
	if err != nil {
		return nil, err
	}

	return &v1.UserMe{
		IdUser:      userProfile.IdUser,
		OnboardSet:  userProfile.OnboardSet,
		PasswordSet: userProfile.PasswordSet,
	}, nil
}
