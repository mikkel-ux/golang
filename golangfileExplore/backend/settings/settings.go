package settings

import (
	"encoding/json"
	files "golangfileExplore/backend/files"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
)

type Settings struct {
	PinnedDirs    []string `json:"pinnedDirs"`    // full paths
	LastOpenedDir string   `json:"lastOpenedDir"` // sisdte åbnede dir
	ViewMode      string   `json:"viewMode"`      // list eller grid
	SortBy        string   `json:"sortBy"`        // name, date, size, type
}

var AppSettings Settings

const AppName = "GolangFileExplore"

func GetConfigFilePath() (string, error) {
	configDir := filepath.Join(xdg.ConfigHome, AppName)
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return "", err
	}
	return filepath.Join(configDir, "settings.json"), nil
}

func SaveSettings(settings *Settings) error {
	filepath, err := GetConfigFilePath()
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(settings, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filepath, data, 0644)
}

func CheckAndCreateSettings() (*Settings, error) {
	filePath, err := GetConfigFilePath()
	if err != nil {
		return nil, err
	}

	var settings Settings
	_, err = os.Stat(filePath)
	if os.IsNotExist(err) {
		dirs := files.GetDefoultDir()

		settings = Settings{
			PinnedDirs:    dirs,
			LastOpenedDir: dirs[0],
			ViewMode:      "list",
			SortBy:        "name",
		}
		err = SaveSettings(&settings)
		if err != nil {
			return nil, err
		}
	} else {
		data, err := LoadSettings()
		if err != nil {
			return nil, err
		}
		settings = *data
	}
	return &settings, nil
}

func LoadSettings() (*Settings, error) {
	filePath, err := GetConfigFilePath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var settings Settings
	err = json.Unmarshal(data, &settings)
	if err != nil {
		return nil, err
	}
	return &settings, nil
}

func GetSettings() *Settings {
	return &AppSettings
}

/* data, err := settings.CheckAndCreateSettings()
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

settings.AppSettings.PinnedDirs = append(settings.AppSettings.PinnedDirs, files.DefoultDirs[files.Pictures])
*/
/* err = settings.SaveSettings(&settings.AppSettings)
if err != nil {
	println("Error saving settings:", err.Error())
	return
} */
