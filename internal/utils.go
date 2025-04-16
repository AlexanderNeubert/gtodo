package internal

import (
	"log"
	"os"
	"path/filepath"
)

func GetJSONPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Unable to determine user home directory:", err)
	}

	path := filepath.Join(homeDir, "Documents/.todos.json")

	return path
}
