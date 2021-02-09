package di

import (
	"github.com/Tatsuemon/anony/domain/repository"
	"github.com/Tatsuemon/anony/domain/service"
	"github.com/Tatsuemon/anony/infrastructure/datastore"
	"github.com/Tatsuemon/anony/infrastructure/web/handler"
	"github.com/Tatsuemon/anony/usecase"
	"github.com/jmoiron/sqlx"
)

// UserController is Di for test
type UserController struct {
	Handler    *handler.UserHandler
	UseCase    usecase.UserUseCase
	Service    service.UserService
	Repository repository.UserRepository
}

func SetUserController(db *sqlx.DB, t datastore.Transaction) UserController {
	repository := datastore.NewUserRepository(db)
	service := service.NewUserService(repository)
	usecase := usecase.NewUserUseCase(repository, t, service)
	handler := handler.NewUserHandler(usecase)

	return UserController{
		Handler:    handler,
		UseCase:    usecase,
		Service:    service,
		Repository: repository,
	}
}
