package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/fesInformatics/jiro-stamp-rally/interface/context"
	"github.com/fesInformatics/jiro-stamp-rally/usecase/interactor"
)

type UserController struct {
	interactor interactor.UserInteractor
}

func (c UserController) Create(ctx context.Context) {
	switch ctx.GetHttpMethod() {
	case http.MethodOptions:
		ctx.JSON(http.StatusOK, nil)
	case http.MethodPost:
		var user UserCreate
		if err := json.NewDecoder(ctx.GetBody()).Decode(&user); err != nil {
			ctx.BadRequest(err)
		}

		err := c.interactor.Save(user.MailAddress, user.Password)
		if err != nil && errors.Is(err, interactor.ErrUserDuplicate) {
			ctx.BadRequest(err)
		} else if err != nil {
			ctx.InternalServerError(err)
		}

		ctx.JSON(http.StatusCreated, nil)
	default:
		ctx.MethodNotAllowed(fmt.Errorf("MethodNotAllowed"))
	}
}

func NewUserController(interactor interactor.UserInteractor) UserController {
	return UserController{interactor: interactor}
}

type UserCreate struct {
	MailAddress string `json:"mailAddress"`
	Password    string `json:"password"`
}
