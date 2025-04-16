package cmd

import (
	"fmt"
	"gtodo/internal"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var initJSON = &cobra.Command{
	Use:   "init",
	Short: "Generate a JSON file for task storage",
	Long:  "Generate a JSON file, which will be used to store tasks. If the file already exists, it notifies the user.",

	Run: generateJSON,
}

func init() {
	rootCmd.AddCommand(initJSON)
}

func generateJSON(cmd *cobra.Command, args []string) {
	path := internal.GetJSONPath()

	// Check if the file already exists
	if _, err := os.Stat(path); err == nil {
		fmt.Printf("The file already exists at the specified path: %s\n", path)
		return
	} else if !os.IsNotExist(err) {
		log.Fatalf("Error checking file existence: %v", err)
	}

	// Create the file if it does not exist
	file, err := os.Create(path)
	if err != nil {
		log.Fatalf("Failed to create the file: %v", err)
	}
	defer file.Close()

	fmt.Println("Successfully created the file at the specified path.")
}
