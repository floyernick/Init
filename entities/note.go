package entities

type Note struct {
	Id    string
	Title string
	Data  string
}

func (note *Note) Exists() bool {
	return note.Id != ""
}
