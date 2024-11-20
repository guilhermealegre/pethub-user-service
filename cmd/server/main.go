package main

import (
	"fmt"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/grpc"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/logger"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/tracer"
	v1UserController "github.com/guilhermealegre/pethub-user-service/internal/user/controller/v1"

	"os"

	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/validator"

	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/database"

	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/app"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/http"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/redis"

	v1AliveController "github.com/guilhermealegre/pethub-user-service/internal/alive/controller/v1"
	v1AliveModel "github.com/guilhermealegre/pethub-user-service/internal/alive/model/v1"
	v1UserModel "github.com/guilhermealegre/pethub-user-service/internal/user/model/v1"

	v1UserRepository "github.com/guilhermealegre/pethub-user-service/internal/user/repository/v1"

	v1UserStreaming "github.com/guilhermealegre/pethub-user-service/internal/user/streaming/v1"

	v1Middleware "github.com/guilhermealegre/pethub-user-service/internal/middleware/v1"
	v1SwaggerController "github.com/guilhermealegre/pethub-user-service/internal/swagger/controller/v1"
	_ "github.com/lib/pq" // postgres driver
)

func main() {
	// app initialization
	newApp := app.New(nil)
	newHttp := http.New(newApp, nil)
	newTracer := tracer.New(newApp, nil)
	newLogger := logger.New(newApp, nil)
	newValidator := validator.New(newApp).
		AddFieldValidators().
		AddStructValidators()
	newRedis := redis.New(newApp, nil)
	newDatabase := database.New(newApp, nil)
	newGrpc := grpc.New(newApp, nil)

	// streaming
	userStreaming := v1UserStreaming.NewStreaming(newApp)
	// repository
	// user
	userRepository := v1UserRepository.NewRepository(newApp)
	// models
	aliveModel := v1AliveModel.NewModel(newApp)
	// user
	userModel := v1UserModel.NewModel(newApp, userRepository, userStreaming)

	newGrpc.
		WithController(v1UserController.NewGrpcController(newApp, userModel))

	newHttp.
		//middlewares
		WithMiddleware(v1Middleware.NewAuthenticateMiddleware(newApp)).
		WithMiddleware(v1Middleware.NewPrintRequestMiddleware(newApp)).
		WithMiddleware(v1Middleware.NewPrepareContextMiddleware(newApp)).
		//controllers
		WithController(v1SwaggerController.NewController(newApp)).
		WithController(v1AliveController.NewController(newApp, aliveModel)).
		//user
		WithController(v1UserController.NewController(newApp, userModel))

	newApp.
		WithValidator(newValidator).
		WithDatabase(newDatabase).
		WithRedis(newRedis).
		WithLogger(newLogger).
		WithTracer(newTracer).
		WithGrpc(newGrpc).
		WithHttp(newHttp)

	// start app
	if err := newApp.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
