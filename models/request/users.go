package request

type UsersGet struct {
	Id int `json:"id" validate:"required; min=1"`
}
