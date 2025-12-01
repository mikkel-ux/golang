package files

import (
	"os"

	"github.com/adrg/xdg"
)

type Files struct{}

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
