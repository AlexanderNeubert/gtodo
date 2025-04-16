package cmd

import (
	"fmt"
	"gtodo/internal"
	"gtodo/todo"
	"os"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a todo item",
	Long:  "Delete a todo item by its ID.",

	Run: deleteTask,
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().IntP("id", "i", 0, "The ID of the todo item to delete")
	deleteCmd.MarkFlagRequired("id")
}

func deleteTask(cmd *cobra.Command, args []string) {
	id, err := cmd.Flags().GetInt("id")
	if err != nil {
		fmt.Printf("Error getting id flag: %v\n", err)
		return
	}

	todos := &todo.Todos{}

	// Load existing todos
	filePath := internal.GetJSONPath()
	if _, err := os.Stat(filePath); err == nil {
		if err := todos.Load(filePath); err != nil {
			fmt.Printf("Error loading todos: %v\n", err)
			return
		}
	} else {
		fmt.Println("No todos found to delete.")
		return
	}

	// Delete the todo
	if err := todos.Delete(id); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Save the updated list
	if err := todos.Store(filePath); err != nil {
		fmt.Printf("Error saving todos: %v\n", err)
		return
	}

	todos.Print(2, "")
	fmt.Println("Todo item deleted successfully.")
}
