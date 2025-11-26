package main

import (
	"encoding/json"
	"os"

	"github.com/adrg/xdg"
)

type defoultDIr string

const (
	Pictures  defoultDIr = "Pictures"
	Documents defoultDIr = "Documents"
	Downloads defoultDIr = "Downloads"
	Music     defoultDIr = "Music"
	Videos    defoultDIr = "Videos"
)

var DefoultDirs = map[defoultDIr]string{
	Pictures:  xdg.UserDirs.Pictures,
	Documents: xdg.UserDirs.Documents,
	Downloads: xdg.UserDirs.Download,
	Music:     xdg.UserDirs.Music,
	Videos:    xdg.UserDirs.Videos,
}

type settings struct {
	PinnedDirs    []string `json:"pinnedDirs"`    // full paths
	LastOpenedDir string   `json:"lastOpenedDir"` // sisdte åbnede dir
	ViewMode      string   `json:"viewMode"`      // list or grid
	SortBy        string   `json:"sortBy"`        // name, date, size, type
}

const SettingsFileName = "golangfileExplore_settings.json"

func saveSettings(settings *settings) error {
	data, err := json.MarshalIndent(settings, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(SettingsFileName, data, 0644)
}

func main() {
	/* names := make([]string, 0, len(DefoultDirs))
	for name := range DefoultDirs {
		names = append(names, string(name))
	}
	for _, name := range names {
		println(name)
	}
	var dirName defoultDIr = Pictures
	test := DefoultDirs[dirName]
	println(test) */

}
