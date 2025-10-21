package cmd

import (
	"fmt"
	"os"

	"github.com/JoaoPedr0Maciel/charm/internal/http"
	"github.com/spf13/cobra"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

// SetVersion define a versão do CLI
func SetVersion(v, c, d string) {
	version = v
	commit = c
	date = d
}

var rootCmd = &cobra.Command{
	Use:   "charm",
	Short: "Charm is a tool for making HTTP requests",
	Long:  `Charm is a tool for managing your charm projects`,
}

var getCmd = &cobra.Command{
	Use:   "get [url]",
	Short: "Make a GET request to the specified URL",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		bearer, _ := cmd.Flags().GetString("bearer")
		basic, _ := cmd.Flags().GetString("basic")
		contentType, _ := cmd.Flags().GetString("content-type")
		_, err := http.Get(url, &bearer, &basic, &contentType)
		if err != nil {
			fmt.Println("Error making GET request:", err)
			os.Exit(1)
		}
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the version of Charm",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Charm %s\n", version)
		fmt.Printf("Commit: %s\n", commit)
		fmt.Printf("Built: %s\n", date)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(versionCmd)
	getCmd.Flags().String("bearer", "", "Bearer token")
	getCmd.Flags().String("basic", "", "Basic auth username and password")
	getCmd.Flags().String("content-type", "", "Content-Type header")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
