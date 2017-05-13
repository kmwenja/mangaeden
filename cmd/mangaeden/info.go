package main

import (
	"fmt"
	"os"

	"github.com/kmwenja/mangaeden"
	"github.com/spf13/cobra"
)

func InfoCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "info [manga id]",
		Short: "Print all the relevant info of the manga id provided",
		Run: func(ccmd *cobra.Command, args []string) {
			if len(args) < 1 {
				ccmd.HelpFunc()(ccmd, args)
				os.Exit(1)
			}

			info(args)
		},
	}

	return cmd
}

func info(args []string) {
	id := args[0]

	c := mangaeden.New(nil)
	mi, err := c.Manga(id)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("ID: %s\n", id)
	fmt.Printf("Title: %s\n", mi.Title)
	fmt.Printf("Artist: %s\n", mi.Artist)
	fmt.Printf("Author: %s\n", mi.Author)
	fmt.Printf("Description: %s\n", mi.Description())
	fmt.Printf("Categories: %s\n", mi.Categories())
	fmt.Printf("No of chapters: %d\n", len(mi.Chapters))
	fmt.Printf("Status: %s\n", mi.Status())
	fmt.Printf("Last Chapter Date: %s\n", mi.LastChapterDate)
	fmt.Printf("Language: %s\n", mi.Language())
	fmt.Printf("Hits: %d\n", mi.Hits)
}
