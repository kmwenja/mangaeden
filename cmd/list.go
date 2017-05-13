package main

import (
	"fmt"

	"github.com/kmwenja/mangaeden"
	"github.com/spf13/cobra"
)

func ListCmd() *cobra.Command {
	var page, size int
	var language string

	var cmd = &cobra.Command{
		Use:   "list",
		Short: "Print out a list of all manga",
		Run: func(ccmd *cobra.Command, args []string) {
			list(page, size, language)
		},
	}

	cmd.Flags().IntVarP(&page, "page", "p", -1, "page to retrieve from")
	cmd.Flags().IntVarP(&size, "size", "s", -1, "page size, must be between 25 and 1500")
	cmd.Flags().StringVarP(&language, "language", "l", "english", "API language: choose between `english` and `italian`")

	return cmd
}

func list(page int, size int, language string) {
	var lang int
	switch language {
	case "english":
		lang = mangaeden.LANG_ENG
	case "italian":
		lang = mangaeden.LANG_ITA
	default:
		fmt.Printf("Unknown language: %s\n", language)
		return
	}

	c := mangaeden.New(nil)
	if page != -1 {
		var mp mangaeden.MangaPage
		var err error
		if size != -1 {
			mp, err = c.ListPage(lang, page, size)
		} else {
			mp, err = c.ListPage(lang, page, 0)
		}

		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		for _, m := range mp.MangaList {
			fmt.Printf("%v\n", m)
		}
		fmt.Printf("Current Page: %d\n", mp.Page)
		fmt.Printf("Current Page Size: %d\n", len(mp.MangaList))
		fmt.Printf("Total Pages: %d\n", mp.Pages())
		fmt.Printf("Total: %d\n", mp.Total)
		return
	}

	ms, err := c.ListAll(mangaeden.LANG_ENG)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for _, m := range ms {
		fmt.Printf("%v\n", m)
	}
	fmt.Printf("Total: %d\n", len(ms))
}
