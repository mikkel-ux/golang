package backend

import (
	"log"
	"os"
)

func CreateUploadsDir() error {
	checkDir, err := os.Stat("./uploads")
	if err != nil && os.IsNotExist(err) {
		checkDir = nil
	}

	if checkDir == nil {
		err = os.Mkdir("./uploads", os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create uploads directory: %v\n", err)
			return err
		}
	}
	return nil
}
