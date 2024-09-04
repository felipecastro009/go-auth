package usecases

import (
	"encoding/json"
	"first-project/internal/domain"
	"first-project/internal/infra/database/repository"
	"net/http"
)

type CreateUserUseCase struct {
	Repository *repository.UserRepository
}

func NewCreateUserUseCase(Repository *repository.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{Repository: Repository}
}

func (c *CreateUserUseCase) Execute(w http.ResponseWriter, r *http.Request) {
	var user *domain.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	err = c.Repository.Create(user)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusCreated)
}
