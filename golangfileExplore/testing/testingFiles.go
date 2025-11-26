package main

import (
	settings "golangfileExplore/backend/settings"

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

	data, err := settings.CheckAndCreateSettings()
	if err != nil {
		println("Error checking/creating settings:", err.Error())
		return
	} else {
		settings.AppSettings = *data
	}
	println("Settings loaded:")
	for i, dir := range settings.AppSettings.PinnedDirs {
		println("Pinned Dir", i, ":", dir)
	}
	println("Last Opened Dir:", settings.AppSettings.LastOpenedDir)
	println("View Mode:", settings.AppSettings.ViewMode)
	println("Sort By:", settings.AppSettings.SortBy)

	settings.AppSettings.PinnedDirs = append(settings.AppSettings.PinnedDirs, DefoultDirs[Pictures])

	err = settings.SaveSettings(&settings.AppSettings)
	if err != nil {
		println("Error saving settings:", err.Error())
		return
	}
}
