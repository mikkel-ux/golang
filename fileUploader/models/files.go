package models

type File struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Extension string `json:"extension"`
	FileType  string `json:"fileType"`
}
