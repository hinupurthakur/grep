/*
Copyright © 2022 NAME HERE <hinupurthakur@gmail.com>

*/
package cli

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	availableSubcommands = []*cobra.Command{}
)

func grepFunctionality(fileName, regex string) ([]string, error) {
	_, err := os.Stat(fileName)
	if err != nil {
		return nil, err
	}
	f, err := os.Open(fileName)
	defer f.Close()
	if err != nil && err != io.EOF {
		return nil, err
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), regex) {
			txtlines = append(txtlines, scanner.Text())
		}
	}
	return txtlines, err
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "grep",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		matchedLines, err := grepFunctionality(args[1], args[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(matchedLines)
		return
	},

	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.democtl.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	availableSubcommands = []*cobra.Command{
		versionCmd,
	}
	rootCmd.AddCommand(availableSubcommands...)

}
