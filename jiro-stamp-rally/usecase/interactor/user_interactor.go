package interactor

import (
	"github.com/fesInformatics/jiro-stamp-rally/usecase/repository"
)

type UserInteractor struct {
	repository repository.UserRepository
}

func (i *UserInteractor) Save(mailAddress string, Password string) error {
	return i.repository.Save(mailAddress, Password)
}

func NewUserInteactor(repository repository.UserRepository) UserInteractor {
	return UserInteractor{
		repository: repository,
	}
}
