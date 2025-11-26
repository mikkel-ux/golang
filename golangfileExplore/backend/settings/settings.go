package settings

import (
	"encoding/json"
	"os"
)

type Settings struct {
	PinnedDirs    []string `json:"pinnedDirs"`    // full paths
	LastOpenedDir string   `json:"lastOpenedDir"` // sisdte åbnede dir
	ViewMode      string   `json:"viewMode"`      // list or grid
	SortBy        string   `json:"sortBy"`        // name, date, size, type
}

var AppSettings Settings

const SettingsFileName = "golangfileExplore_settings.json"

func SaveSettings(settings *Settings) error {
	data, err := json.MarshalIndent(settings, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(SettingsFileName, data, 0644)
}

func CheckAndCreateSettings() (*Settings, error) {
	var settings Settings
	_, err := os.ReadFile(SettingsFileName)
	if err != nil {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			homeDir = "/"
		}
		settings = Settings{
			PinnedDirs:    []string{},
			LastOpenedDir: homeDir,
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
	data, err := os.ReadFile(SettingsFileName)
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
