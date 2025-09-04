package backend

import (
	"log"
	"os"
)

func CreateDir() error {
	checkUploadsDir, err := os.Stat("./uploads")
	if err != nil && os.IsNotExist(err) {
		checkUploadsDir = nil
	}

	if checkUploadsDir == nil {
		err = os.Mkdir("./uploads", os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create uploads directory: %v\n", err)
			return err
		}
	}

	checkUsersDir, err := os.Stat("./users")
	if err != nil && os.IsNotExist(err) {
		checkUsersDir = nil
	}

	if checkUsersDir == nil {
		err = os.Mkdir("./users", os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create users directory: %v\n", err)
			return err
		}
		f, err := os.Create("./users/users.json")
		if err != nil {
			log.Fatalf("Failed to create users.json file: %v\n", err)
			return err
		}
		defer f.Close()

		err = os.WriteFile("./users/users.json", []byte("[]"), 0644)
		if err != nil {
			log.Fatalf("Failed to initialize users.json file: %v\n", err)
			return err
		}
	}

	checkTempDir, err := os.Stat("./temp")
	if err != nil && os.IsNotExist(err) {
		checkTempDir = nil
	}
	if checkTempDir == nil {
		err = os.Mkdir("./temp", os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create temp directory: %v\n", err)
			return err
		}
	}

	return nil
}
