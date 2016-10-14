package main

import (
	"fmt"
	"os"

	"github.com/kmwenja/mangaeden"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [chapter id]\n", os.Args[0])
		return
	}

	c := mangaeden.New(nil)
	is, err := c.Chapter(os.Args[1])
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for _, i := range is {
		fmt.Printf("%v\n", i)
		fmt.Printf("%s\n", i.ImageUrl())
	}
}
