package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kmwenja/mangaeden"
)

func main() {
	var verbose bool
	flag.BoolVar(&verbose, "verbose", false, "show a lot")
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Printf("Usage: %s [manga id]\n", os.Args[0])
		return
	}

	id := flag.Arg(0)

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
