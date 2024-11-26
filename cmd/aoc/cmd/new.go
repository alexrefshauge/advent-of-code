/*
Copyright © 2024 Alexander Refshauge
*/
package cmd

import (
	"fmt"
	"github.com/alexrefshauge/advent-of-code/cmd/aoc/configuration"
	"github.com/alexrefshauge/advent-of-code/cmd/aoc/configurationKeys"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new solution, will default to today",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		return newSolution()
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}

func newSolution() error {
	if !viper.IsSet(configurationKeys.SolutionDirectory) {
		fmt.Println("Please configure a solution directory before creating a new solution")
		configuration.UseAsSolutionDirectoryPrompt()
	}

	solutionsDirectory := viper.GetString(configurationKeys.SolutionDirectory)
	year, day, err := configuration.GetDate()
	if err != nil {
		return err
	}
	currentDirectory := path.Join(solutionsDirectory, fmt.Sprintf("%d", year), fmt.Sprintf("%d", day))
	os.MkdirAll(currentDirectory, os.ModePerm)
	return nil
}
