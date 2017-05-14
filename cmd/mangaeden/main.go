package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version string

func main() {
	var RootCmd = &cobra.Command{
		Use:   "mangaeden",
		Short: "Mangaeden is a CLI for the Mangaeden API",
		Long:  `Mangeden is a simple lookup and download CLI for the Mangaeden API. Make sure to checkout www.mangaeden.com`,
	}

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print out the version of this app",
		Run: func(ccmd *cobra.Command, args []string) {
			fmt.Printf("%s\n", version)
		},
	}

	RootCmd.AddCommand(ListCmd())
	RootCmd.AddCommand(InfoCmd())
	RootCmd.AddCommand(DownloadCmd())
	RootCmd.AddCommand(versionCmd)

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
