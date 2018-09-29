package usecases

import (
	"Init/models/request"
	"Init/models/response"
	"errors"
)

func (controller Controller) UsersGet(req request.UsersGet) (response.UsersGet, error) {

	user, err := controller.db.GetUserById(req.Id)

	if err != nil {
		return response.UsersGet{}, errors.New("internal error")
	}

	if user.Id == 0 {
		return response.UsersGet{}, errors.New("invalid user id")
	}

	res := response.UsersGet{
		Id:   user.Id,
		Name: user.Name,
	}

	return res, nil

}
