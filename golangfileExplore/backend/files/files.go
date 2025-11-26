package files

import (
	"os"

	"github.com/adrg/xdg"
)

type Files struct{}

/* var defoultDirs = []map[string]string{
	{"name": "Pictures", "path": xdg.UserDirs.Pictures},
	{"name": "Documents", "path": xdg.UserDirs.Documents},
	{"name": "Downloads", "path": xdg.UserDirs.Download},
	{"name": "Music", "path": xdg.UserDirs.Music},
	{"name": "Videos", "path": xdg.UserDirs.Videos},
} */

type defoultDIr string

const (
	Pictures  defoultDIr = "Pictures"
	Documents defoultDIr = "Documents"
	Downloads defoultDIr = "Downloads"
	Music     defoultDIr = "Music"
	Videos    defoultDIr = "Videos"
)

var DefoultDirs = []struct {
	value defoultDIr
	path  string
}{
	{Pictures, xdg.UserDirs.Pictures},
	{Documents, xdg.UserDirs.Documents},
	{Downloads, xdg.UserDirs.Download},
	{Music, xdg.UserDirs.Music},
	{Videos, xdg.UserDirs.Videos},
}

func NewFiles() *Files {
	return &Files{}
}

func ListFiles() ([]os.DirEntry, error) {
	test := defoultDIr("Pictures")
	println(test)
	dir := xdg.UserDirs.Pictures
	data, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	return data, nil
}
