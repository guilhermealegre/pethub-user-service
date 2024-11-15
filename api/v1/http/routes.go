package http

import (
	"net/http"

	infra "github.com/guilhermealegre/go-clean-arch-infrastructure-lib/http"
)

var (
	GroupV1  = infra.NewGroup("api").Group("v1")
	GroupV1P = GroupV1.Group("p")

	//user
	GroupV1User  = GroupV1.Group("user")
	GroupV1PUser = GroupV1P.Group("user")

	//documentation
	GroupV1PDocumentation     = GroupV1P.Group("documentation")
	GroupV1PDocumentationUser = GroupV1PDocumentation.Group("user")

	SwaggerUserDocs    = GroupV1PDocumentationUser.NewEndpoint("/docs", http.MethodGet)
	SwaggerUserSwagger = GroupV1PDocumentationUser.NewEndpoint("/swagger", http.MethodGet)

	Alive           = GroupV1.NewEndpoint("/alive", http.MethodGet)
	PublicAliveUser = GroupV1P.NewEndpoint("/alive/user", http.MethodGet)

	//user
	OnboardUser       = GroupV1User.NewEndpoint("/onboard", http.MethodPost)
	GetUserMe         = GroupV1User.NewEndpoint("/me", http.MethodGet)
	GetUserProfile    = GroupV1User.NewEndpoint("/profile", http.MethodGet)
	UpdateUserProfile = GroupV1User.NewEndpoint("/profile", http.MethodPut)
)
