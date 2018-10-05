package usecases

import (
	"Init/models/request"
	"Init/models/response"

	"errors"

	"github.com/floyernick/validator-go"
)

func (controller Controller) UsersGet(req request.UsersGet) (response.UsersGet, error) {

	var res response.UsersGet

	if ok := validator.Process(req); !ok {
		return res, errors.New("invalid request params")
	}

	user, err := controller.db.GetUserById(req.Id)

	if err != nil {
		return res, errors.New("internal error")
	}

	if user.Id == 0 {
		return res, errors.New("invalid user id")
	}

	res = response.UsersGet{
		Id:   user.Id,
		Name: user.Name,
	}

	return res, nil

}
