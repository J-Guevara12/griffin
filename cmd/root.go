/*
Copyright © 2025 Juan Esteban Guevara juguevara@unal.edu.co

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "griffin",
	Short: "A task manager developed to Grind without finish",
	Long: `A task manager developed to Grind without finish.`,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	/*Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Printing hi message: ", viper.Get("hi"))
    },*/
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
    if err := RootCmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
      }
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.griffin.yaml)")
    RootCmd.AddCommand()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".griffin.yaml"
		viper.AddConfigPath(home)
        // It can also be in the CWD
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".griffin.yaml")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
    if err := viper.ReadInConfig(); err != nil {
        if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		    fmt.Fprintln(os.Stderr, "Config file not found, please run 'griffin configure'")
            os.Exit(1)
        } else {
		    fmt.Fprintln(os.Stderr, "There was an error parsing the config file")
            os.Exit(1)
        }
    }
}
