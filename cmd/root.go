package cmd

import (
	"fmt"
	"os"

	client "github.com/JoaoPedr0Maciel/charm/internal/http"
	"github.com/JoaoPedr0Maciel/charm/internal/structs"
	"github.com/JoaoPedr0Maciel/charm/internal/updater"
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

var httpMethods = []structs.HTTPMethod{
	{Name: "get", HasBody: false, HTTPFuncNoBody: client.Get},
	{Name: "post", HasBody: true, HTTPFunc: client.Post},
	{Name: "put", HasBody: true, HTTPFunc: client.Put},
	{Name: "patch", HasBody: true, HTTPFunc: client.Patch},
	{Name: "delete", HasBody: true, HTTPFunc: client.Delete},
}

func createHTTPCommand(method structs.HTTPMethod) *cobra.Command {
	cmd := &cobra.Command{
		Use:   fmt.Sprintf("%s [url]", method.Name),
		Short: fmt.Sprintf("Make a %s request to the specified URL", method.Name),
		Args:  cobra.ExactArgs(1),
		RunE:  makeHTTPRequestFunc(method),
	}

	addCommonFlags(cmd)

	if method.HasBody {
		cmd.Flags().String("data", "", "Request body data (JSON)")
		cmd.Flags().StringP("data-raw", "d", "", "Request body data (raw)")
	}

	return cmd
}

func makeHTTPRequestFunc(method structs.HTTPMethod) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		url := args[0]

		bearer, _ := cmd.Flags().GetString("bearer")
		basic, _ := cmd.Flags().GetString("basic")
		contentType, _ := cmd.Flags().GetString("content-type")

		var err error
		if method.HasBody {
			data, _ := cmd.Flags().GetString("data-raw")
			if data == "" {
				data, _ = cmd.Flags().GetString("data")
			}
			_, err = method.HTTPFunc(url, bearer, basic, contentType, data)
		} else {
			_, err = method.HTTPFuncNoBody(url, bearer, basic, contentType)
		}

		if err != nil {
			return fmt.Errorf("%s request failed: %w", method.Name, err)
		}

		return nil
	}
}

func addCommonFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("bearer", "b", "", "Bearer token for authentication")
	cmd.Flags().String("basic", "", "Basic auth in format 'username:password'")
	cmd.Flags().StringP("content-type", "H", "", "Content-Type header")
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the version of Charm",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("charm %s\n", version)
	},
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update Charm to the latest version",
	Long:  `Check for updates and automatically download and install the latest version of Charm from GitHub releases`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return updater.Update(version)
	},
}

func init() {
	for _, method := range httpMethods {
		rootCmd.AddCommand(createHTTPCommand(method))
	}

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(updateCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
