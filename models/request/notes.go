package request

type NotesCreate struct {
	Title string `json:"title" validate:"required,min=1"`
	Data  string `json:"data" validate:"required,min=1"`
}

type NotesUpdate struct {
	Id    string  `json:"id" validate:"required,min=36,max=36"`
	Title *string `json:"title" validate:"omitempty,min=1"`
	Data  *string `json:"data" validate:"omitempty,min=1"`
}

type NotesDelete struct {
	Id string `json:"id" validate:"required,min=36,max=36"`
}

type NotesGet struct {
	Id string `json:"id" validate:"required,min=36,max=36"`
}
