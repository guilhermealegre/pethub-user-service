package v1

import (
	"github.com/gocraft/dbr/v2"
	"github.com/google/uuid"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain"
	dCtx "github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain/context"
	"github.com/guilhermealegre/pethub-user-service/internal/infrastructure/database"
	"github.com/guilhermealegre/pethub-user-service/internal/user/domain/v1"
)

type Repository struct {
	app domain.IApp
}

func NewRepository(app domain.IApp) v1.IRepository {
	return &Repository{
		app: app,
	}
}

// CreateUser create new user
func (r *Repository) CreateUser(ctx dCtx.IContext, uuidUser uuid.UUID) (idUser int, err error) {
	_, err = r.app.Database().Write().InsertInto(database.UserTableUser).
		Columns(
			"uuid",
		).
		Values(
			uuidUser,
		).
		Returning("id_users").
		Record(&idUser).
		ExecContext(ctx)

	if err != nil {
		return 0, r.app.Logger().DBLog(err)
	}

	return idUser, nil
}

// GetUserProfile get user profile
func (r *Repository) GetUserProfile(ctx dCtx.IContext, idUser int) (userProfile *v1.UserProfile, err error) {
	columns := []any{
		"COALESCE(first_name, '') as first_name",
		"COALESCE(first_name, '') as first_name",
		"COALESCE(last_name, '') as last_name",
		"onboard_set",
		"password_set",
	}

	_, err = r.app.Database().Read().Select(columns...).
		From(database.UserTableUser).
		Where("id_users  = ?", idUser).
		LoadContext(ctx, &userProfile)

	if err != nil {
		return nil, r.app.Logger().DBLog(err)
	}

	return userProfile, nil
}

// UpdateUserProfile update user profile
func (r *Repository) UpdateUserProfile(ctx dCtx.IContext, tx dbr.SessionRunner, idUser int, profile *v1.UserProfile) error {
	if tx == nil {
		tx = r.app.Database().Write()
	}

	_, err := tx.Update(database.UserTableUser).
		Set("first_name", profile.FirstName).
		Set("last_name", profile.LastName).
		Set("onboard_set", true).
		Where("id_users = ?", idUser).
		ExecContext(ctx)

	if err != nil {
		return r.app.Logger().DBLog(err)
	}

	return nil
}
