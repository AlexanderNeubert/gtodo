package cmd

import (
	"fmt"
	"gtodo/internal"
	"gtodo/todo"
	"os"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a todo item",
	Long:  "Update a todo item's task, category, or completion status by ID.",
	Run:   updateTask,
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().IntP("id", "i", 0, "The ID of the todo item to update")
	updateCmd.Flags().StringP("task", "t", "", "The new content for the todo item")
	updateCmd.Flags().StringP("cat", "c", "", "The new category for the todo item")

	// Replace the integer done flag with two boolean flags
	updateCmd.Flags().BoolP("done", "d", false, "Mark the task as completed")
	updateCmd.Flags().BoolP("undone", "u", false, "Mark the task as not completed")

	updateCmd.MarkFlagRequired("id")
}

func updateTask(cmd *cobra.Command, args []string) {
	id, err := cmd.Flags().GetInt("id")
	if err != nil {
		fmt.Printf("Error getting id flag: %v\n", err)
		return
	}

	task, err := cmd.Flags().GetString("task")
	if err != nil {
		fmt.Printf("Error getting task flag: %v\n", err)
		return
	}

	category, err := cmd.Flags().GetString("cat")
	if err != nil {
		fmt.Printf("Error getting category flag: %v\n", err)
		return
	}

	// Get the done and undone flags
	done, err := cmd.Flags().GetBool("done")
	if err != nil {
		fmt.Printf("Error getting done flag: %v\n", err)
		return
	}

	undone, err := cmd.Flags().GetBool("undone")
	if err != nil {
		fmt.Printf("Error getting undone flag: %v\n", err)
		return
	}

	// Validate that done and undone aren't both specified
	if done && undone {
		fmt.Println("Error: Cannot specify both --done and --undone flags")
		return
	}

	if id <= 0 {
		fmt.Println("Error: Please provide a valid ID to update.")
		return
	}

	if task == "" && category == "" && !done && !undone {
		fmt.Println("Error: At least one of --task, --cat, --done, or --undone must be specified.")
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
		fmt.Println("No todos found to update.")
		return
	}

	// Convert boolean flags to the integer status for the Update method
	doneStatus := 2 // unchanged by default
	if done {
		doneStatus = 1 // complete
	} else if undone {
		doneStatus = 0 // incomplete
	}

	// Update the todo
	if err := todos.Update(id, task, category, doneStatus); err != nil {
		fmt.Printf("Error updating todo: %v\n", err)
		return
	}

	// Save the updated list
	if err := todos.Store(filePath); err != nil {
		fmt.Printf("Error saving todos: %v\n", err)
		return
	}

	todos.Print(2, "")
	fmt.Println("Todo item updated successfully.")
}
