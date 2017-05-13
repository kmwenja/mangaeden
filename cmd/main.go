package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var RootCmd = &cobra.Command{
		Use:   "mangaeden",
		Short: "Mangaeden is a CLI for the Mangaeden API",
		Long: `Mangeden is a simple lookup and download CLI
			   for the Mangaeden API`,
	}

	RootCmd.AddCommand(ListCmd())
	RootCmd.AddCommand(InfoCmd())
	RootCmd.AddCommand(DownloadCmd())

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
