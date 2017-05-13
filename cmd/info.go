package main

import (
	"fmt"
	"os"

	"github.com/kmwenja/mangaeden"
	"github.com/spf13/cobra"
)

func InfoCmd() *cobra.Command {
	var verbose bool

	var cmd = &cobra.Command{
		Use:   "info [manga id]",
		Short: "Print all the relevant info of the manga id provided",
		Run: func(ccmd *cobra.Command, args []string) {
			if len(args) < 1 {
				ccmd.HelpFunc()(ccmd, args)
				os.Exit(1)
			}

			info(verbose, args)
		},
	}

	cmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "show a lot more")

	return cmd
}

func info(verbose bool, args []string) {
	id := args[0]

	c := mangaeden.New(nil)
	mi, err := c.Manga(id)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if verbose {
		fmt.Printf("Info: %v\n", mi)
		fmt.Printf("Image: %s\n", mi.ImageUrl())
	} else {
		fmt.Printf("Title: %s\n", mi.Title)
		fmt.Printf("Description: %s\n", mi.Description)
		fmt.Printf("No of chapters: %d\n", len(mi.Chapters))
		fmt.Printf("Completed: %d\n", mi.Status)
	}
}
