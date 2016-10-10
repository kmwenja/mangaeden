package main

import (
	"fmt"
	"os"

	"github.com/kmwenja/mangaeden"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [manga id]\n", os.Args[0])
		return
	}

	c := mangaeden.New()
	mi, err := c.GetInfo(os.Args[1])
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Info: %v\n", mi)
	fmt.Printf("Image: %s\n", mi.ImageUrl())
}
