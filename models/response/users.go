package response

type NotesCreate struct {
	Id string `json:"id"`
}

type NotesUpdate struct{}

type NotesDelete struct{}

type NotesGet struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Data  string `json:"data"`
}
