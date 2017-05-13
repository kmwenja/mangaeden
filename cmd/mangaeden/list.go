package main

import (
	"fmt"

	"github.com/kmwenja/mangaeden"
	"github.com/spf13/cobra"
)

func ListCmd() *cobra.Command {
	var language string

	var cmd = &cobra.Command{
		Use:   "list",
		Short: `Print out a list of all manga with the following format: "ID: Title [Categories] Status"`,
		Run: func(ccmd *cobra.Command, args []string) {
			list(language)
		},
	}

	cmd.Flags().StringVarP(&language, "language", "l", "english", "API language: choose between `english` and `italian`")

	return cmd
}

func list(language string) {
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
	ms, err := c.ListAll(lang)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for _, m := range ms {
		fmt.Printf("%s: %s [%s] %s\n", m.ID, m.Title, m.Categories(), m.Status())
	}
	fmt.Printf("\nTotal: %d\n", len(ms))
}
