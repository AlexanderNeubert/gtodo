package cmd

import (
	"fmt"
	"gtodo/internal"
	"gtodo/todo"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo item",
	Long:  "Add a new todo item to the list with an optional category.",

	Run: addTask,
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringP("task", "t", "", "The content of the new todo item")
	addCmd.Flags().StringP("cat", "c", "Uncategorized", "The category of the todo item")
}

func addTask(cmd *cobra.Command, args []string) {
	task, _ := cmd.Flags().GetString("task")
	category, _ := cmd.Flags().GetString("cat")

	if task == "" {
		fmt.Println("Error: the --task flag is required for the 'add' command.")
		return
	}

	todos := &todo.Todos{}

	// Load existing todos before adding new one
	filePath := internal.GetJSONPath()
	if _, err := os.Stat(filePath); err == nil {
		if err := todos.Load(filePath); err != nil {
			log.Fatal(err)
		}
	}

	todos.Add(task, category)

	err := todos.Store(filePath)
	if err != nil {
		log.Fatal(err)
	}

	todos.Print(2, "")
	fmt.Println("Todo item added successfully.")
}
