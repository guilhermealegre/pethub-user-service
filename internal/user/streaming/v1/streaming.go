package v1

import (
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain"
	domainUser "github.com/guilhermealegre/pethub-user-service/internal/user/domain/v1"
)

type Streaming struct {
	app domain.IApp
}

func NewStreaming(app domain.IApp) domainUser.IStreaming {
	return &Streaming{
		app: app,
	}
}
