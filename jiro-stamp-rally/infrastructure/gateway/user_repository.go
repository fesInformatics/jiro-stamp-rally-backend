package repository

import (
	database "github.com/fesInformatics/jiro-stamp-rally/infrastructure/gateway/database"

	"github.com/google/uuid"

	_repository "github.com/fesInformatics/jiro-stamp-rally/usecase/repository"
)

type UserRepository struct {
	client database.DbClient
}

func (r *UserRepository) Save(mailAddress string, Password string) error {
	valuesMap := map[string]any{"id": uuid.New().String(), "mail_address": mailAddress, "password": Password}
	err := r.client.InsertSQL("users", valuesMap)
	if err != nil {
		return err
	}
	return err
}

func NewUserRespository(client database.DbClient) _repository.UserRepository {
	return &UserRepository{client: client}
}
