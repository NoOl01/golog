package cmd

import (
	"fmt"
	"github.com/NoOl01/golog/internal/init_json"
	"github.com/spf13/cobra"
)

var filePath string

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Generate config file",
	Long:  "Generating the golog configuration in JSON format.",
	Run: func(cmd *cobra.Command, args []string) {
		file, err := cmd.Flags().GetString("output")
		if err != nil {
			return
		}

		err = init_json.InitJsonConfig(file)
		if err != nil {
			fmt.Printf("%s", err)
			return
		}
		fmt.Printf("file: '%s' created", file)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&filePath, "output", "o", "./golog_config.json", "Path to output config file")
}
