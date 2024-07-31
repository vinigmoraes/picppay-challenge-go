package dependency_injection

import (
	"gorm.io/gorm"
	"picpay-challenge-go/cmd/api/handlers"
	usecase "picpay-challenge-go/internal/usecase/users"
	database "picpay-challenge-go/pkg/database/users"
)

func InjectCreateUserHandler(db *gorm.DB) handlers.CreateUserHandler {
	return handlers.CreateUserHandler{
		UseCase: usecase.CreateUserUseCase{
			Repository: database.UserRepositoryAdapter{DB: db},
		}}
}
