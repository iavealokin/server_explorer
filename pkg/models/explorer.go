package models

type Page struct {
	Files  []File
	NeedUp bool
	Level  int
	Path   string
	UpLink string
}
type File struct {
	Id       int
	Size     int64
	Level    int
	Name     string
	IsFolder bool
	IsImage  bool
	Link     string
}

type Object struct {
	Id    int
	Prev  int
	Path  string
	Level int
}
