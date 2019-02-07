package models

type Note struct {
	Id    string
	Title string
	Data  string
}

func (note *Note) IsEmpty() bool {
	return note.Id == ""
}
