package cmd

import (
	"fmt"
	"os"

	"github.com/JoaoPedr0Maciel/charm/internal/http"
	"github.com/spf13/cobra"
)

var (
	version = "dev"
)

func SetVersion(v string) {
	version = v
}

var rootCmd = &cobra.Command{
	Use:   "charm",
	Short: "Charm is a tool for making HTTP requests",
	Long:  `Charm is a tool to get beautiful and colorful HTTP requests`,
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

var postCmd = &cobra.Command{
	Use:   "post [url]",
	Short: "Make a POST request to the specified URL",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		bearer, _ := cmd.Flags().GetString("bearer")
		basic, _ := cmd.Flags().GetString("basic")
		contentType, _ := cmd.Flags().GetString("content-type")
		data, _ := cmd.Flags().GetString("data")
		_, err := http.Post(url, &bearer, &basic, &contentType, &data)
		if err != nil {
			fmt.Println("Error making POST request:", err)
			os.Exit(1)
		}
	},
}

var putCmd = &cobra.Command{
	Use:   "put [url]",
	Short: "Make a PUT request to the specified URL",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		bearer, _ := cmd.Flags().GetString("bearer")
		basic, _ := cmd.Flags().GetString("basic")
		contentType, _ := cmd.Flags().GetString("content-type")
		data, _ := cmd.Flags().GetString("data")
		_, err := http.Put(url, &bearer, &basic, &contentType, &data)
		if err != nil {
			fmt.Println("Error making PUT request:", err)
			os.Exit(1)
		}
	},
}

var patchCmd = &cobra.Command{
	Use:   "patch [url]",
	Short: "Make a PATCH request to the specified URL",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		bearer, _ := cmd.Flags().GetString("bearer")
		basic, _ := cmd.Flags().GetString("basic")
		contentType, _ := cmd.Flags().GetString("content-type")
		data, _ := cmd.Flags().GetString("data")
		_, err := http.Patch(url, &bearer, &basic, &contentType, &data)
		if err != nil {
			fmt.Println("Error making PATCH request:", err)
			os.Exit(1)
		}
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete [url]",
	Short: "Make a DELETE request to the specified URL",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		bearer, _ := cmd.Flags().GetString("bearer")
		basic, _ := cmd.Flags().GetString("basic")
		contentType, _ := cmd.Flags().GetString("content-type")
		data, _ := cmd.Flags().GetString("data")
		_, err := http.Delete(url, &bearer, &basic, &contentType, &data)
		if err != nil {
			fmt.Println("Error making DELETE request:", err)
			os.Exit(1)
		}
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the version of Charm",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("charm %s\n", version)
	},
}

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Show help for a command",
	Run: func(cmd *cobra.Command, args []string) {
		rootCmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(postCmd)
	rootCmd.AddCommand(putCmd)
	rootCmd.AddCommand(patchCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(helpCmd)

	rootCmd.Flags().String("help", "", "Show help for a command")
	// Flags para GET
	getCmd.Flags().String("bearer", "", "Bearer token")
	getCmd.Flags().String("basic", "", "Basic auth username and password")
	getCmd.Flags().String("content-type", "", "Content-Type header")

	// Flags para POST
	postCmd.Flags().String("bearer", "", "Bearer token")
	postCmd.Flags().String("basic", "", "Basic auth username and password")
	postCmd.Flags().String("content-type", "", "Content-Type header")
	postCmd.Flags().String("data", "", "Request body data (JSON)")

	// Flags para PUT
	putCmd.Flags().String("bearer", "", "Bearer token")
	putCmd.Flags().String("basic", "", "Basic auth username and password")
	putCmd.Flags().String("content-type", "", "Content-Type header")
	putCmd.Flags().String("data", "", "Request body data (JSON)")

	// Flags para PATCH
	patchCmd.Flags().String("bearer", "", "Bearer token")
	patchCmd.Flags().String("basic", "", "Basic auth username and password")
	patchCmd.Flags().String("content-type", "", "Content-Type header")
	patchCmd.Flags().String("data", "", "Request body data (JSON)")

	// Flags para DELETE
	deleteCmd.Flags().String("bearer", "", "Bearer token")
	deleteCmd.Flags().String("basic", "", "Basic auth username and password")
	deleteCmd.Flags().String("content-type", "", "Content-Type header")
	deleteCmd.Flags().String("data", "", "Request body data (JSON)")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
