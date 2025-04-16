package cmd

import (
	"fmt"
	"gtodo/internal"
	"gtodo/todo"
	"os"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List todo items",
	Long:  "List todo items, optionally filtered by completion status or category.",
	Run:   listTasks,
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().IntP("status", "s", 2, "Filter by status (0=incomplete, 1=complete, 2=all)")
	listCmd.Flags().StringP("cat", "c", "", "Filter by category")
}

func listTasks(cmd *cobra.Command, args []string) {
	status, err := cmd.Flags().GetInt("status")
	if err != nil {
		fmt.Printf("Error getting status flag: %v\n", err)
		return
	}

	category, err := cmd.Flags().GetString("cat")
	if err != nil {
		fmt.Printf("Error getting category flag: %v\n", err)
		return
	}

	// Validate status is in the correct range
	if status < 0 || status > 2 {
		fmt.Println("Error: Status must be 0 (incomplete), 1 (complete), or 2 (all).")
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
		fmt.Println("No todos found.")
		return
	}

	// Check if there are any todos
	if len(*todos) == 0 {
		fmt.Println("No todo items found.")
		return
	}

	// Display the todos
	todos.Print(status, category)
}
