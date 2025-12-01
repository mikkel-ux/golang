package main

import "golangfileExplore/backend/settings"

func main() {
	println(" ")
	/* dirs := files.GetDefoultDir()
	for _, dir := range dirs {
		println(dir)
	} */
	data, err := settings.CheckAndCreateSettings()
	if err != nil {
		println("Error checking/creating settings:", err.Error())
		return
	} else {
		println(data.PinnedDirs[0])
		println(" ")
		for _, dir := range data.PinnedDirs {
			println(dir)
		}
	}
}
