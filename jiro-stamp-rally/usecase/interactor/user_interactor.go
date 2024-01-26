package interactor

import (
	"errors"

	"github.com/fesInformatics/jiro-stamp-rally/usecase/repository"
)

var ErrUserDuplicate = errors.New("登録済のアカウントです。")

type UserInteractor struct {
	repository repository.UserRepository
}

func (i *UserInteractor) Save(mailAddress string, Password string) error {
	exists, err := i.repository.ExistsByMailAddress(mailAddress)
	if err != nil {
		return err
	}
	if exists {
		return ErrUserDuplicate
	}
	return i.repository.Save(mailAddress, Password)
}

func NewUserInteactor(repository repository.UserRepository) UserInteractor {
	return UserInteractor{
		repository: repository,
	}
}
