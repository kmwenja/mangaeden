package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/kmwenja/mangaeden"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <manga id>\n", os.Args[0])
		return
	}

	id := os.Args[1]
	c := mangaeden.New(nil)

	m, err := c.Manga(id)
	if err != nil {
		perror(err)
		return
	}

	// make directory of manga using manga.Title
	dir := m.Title
	err = mkdir(dir)
	if err != nil {
		perror(err)
		return
	}
	// download image and put in directory
	err = download(c, m.Image, filepath.Join(dir, m.Title))
	if err != nil {
		perror(err)
		return
	}

	// TODO save json in directory

	fmt.Printf("Downloading %d chapters\n", len(m.Chapters))

	// get accurate chapter history in case of repeating indices
	chapters := make(map[string]mangaeden.Chapter)
	for _, ch := range m.Chapters {
		key := fmt.Sprintf("%d", ch.Index)
		if _, present := chapters[key]; present {
			for i := 1; i < 100; i++ {
				newKey := fmt.Sprintf("%s_%d", key, i)
				if _, present := chapters[newKey]; present == false {
					key = newKey
					break
				}
			}
		}
		chapters[key] = ch
	}

	// for each chapter
	for k, ch := range chapters {
		// make directory of chapter using chapter.Index
		chDir := filepath.Join(dir, k)

		err = mkdir(chDir)
		if err != nil {
			perror(err)
			continue
		}

		// get chapter images
		ims, err := c.Chapter(ch.ID)
		if err != nil {
			perror(err)
			continue
		}

		// get accurate image history in case of repeating indices
		images := make(map[string]mangaeden.ChapterImage)
		for _, im := range ims {
			key := fmt.Sprintf("%d", im.Index)
			if _, present := images[key]; present {
				for i := 1; i < 100; i++ {
					newKey := fmt.Sprintf("%s_%d", key, i)
					if _, present := images[newKey]; present == false {
						key = newKey
						break
					}
				}
			}
			images[key] = im
		}

		for ik, im := range images {
			// download each image and save to chapterimage.Index
			p := filepath.Join(chDir, ik)
			err = download(c, im.Image, p)
			if err != nil {
				perror(err)
				continue
			}
		}

		fmt.Printf("Downloaded chapter %d, %d images\n", ch.Index, len(ims))
	}
	fmt.Printf("Done\n")
}

func perror(e error) {
	fmt.Printf("Error: %v\n", e)
}

func mkdir(p string) error {
	err := os.Mkdir(p, os.ModePerm)
	if os.IsExist(err) {
		return nil
	}

	if err != nil {
		return err
	}

	return nil
}

func download(c *mangaeden.Client, i mangaeden.Image, p string) error {
	fp := fmt.Sprintf("%s%s", p, i.Ext())
	if _, err := os.Stat(fp); err == nil {
		return nil
	}

	f, err := os.Create(fp)
	if err != nil {
		return err
	}
	defer f.Close()

	r, err := c.DownloadImage(i)
	if err != nil {
		return err
	}
	defer r.Close()

	_, err = io.Copy(f, r)
	if err != nil {
		return err
	}

	return nil
}
