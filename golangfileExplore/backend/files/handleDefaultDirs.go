package files

import (
	"os"

	"github.com/adrg/xdg"
)

type defoultDIr string

const (
	Home      defoultDIr = "Home"
	Pictures  defoultDIr = "Pictures"
	Documents defoultDIr = "Documents"
	Downloads defoultDIr = "Downloads"
	Music     defoultDIr = "Music"
	Videos    defoultDIr = "Videos"
)

var DefoultDirs = map[defoultDIr]string{
	Home:      getUserHomeDir(),
	Pictures:  xdg.UserDirs.Pictures,
	Documents: xdg.UserDirs.Documents,
	Downloads: xdg.UserDirs.Download,
	Music:     xdg.UserDirs.Music,
	Videos:    xdg.UserDirs.Videos,
}

func getUserHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = "/"
	}
	return homeDir
}

func GetDefoultDir() []string {
	names := make([]string, 0, len(DefoultDirs))
	for name := range DefoultDirs {
		names = append(names, string(name))
	}
	homeDir := getUserHomeDir()
	names = append([]string{homeDir}, names...)
	return names
}
