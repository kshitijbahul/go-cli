/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/kshitijbahul/go-cli/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: addRun,
}

var priority int

func addRun(cmd *cobra.Command, args []string) {
	fmt.Println("add called")
	items, err := todo.ReadItems(viper.GetString("dataFile"))
	if err != nil {
		fmt.Printf("%v", err)
	}
	for _, x := range args {
		fmt.Println(x)
		item := todo.Item{Text: x}
		item.SetPriority(priority)
		items = append(items, item)
	}
	fmt.Printf("%#v\n", items)
	todo.SaveItems(viper.GetString("dataFile"), items)
}

// The package can have multiple inits to initialize multiple things
// The Order of init call is not guaranteed
func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority: 1,2,3")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
