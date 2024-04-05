/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/kshitijbahul/go-cli/todo"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Use this command to get the list of all the todos",
	Long: `This command is used to get all the todos that you are lookign for. 
	It also has a Priority fields and and a created at and Updated field`,
	Run: func(cmd *cobra.Command, args []string) {
		list, err := todo.ReadItems("./todos.json")
		if err != nil {
			fmt.Printf("%v", err)
		}
		fmt.Println("Your Todos are:", list)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
