package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/kmwenja/mangaeden"
)

type ChapterImageResult struct {
	Key       string
	Index     int
	TimeTaken time.Duration
	Err       error
}

type ChapterResult struct {
	Key       string
	Index     int
	TimeTaken time.Duration
	Err       error
}

func main() {
	start := time.Now()

	var parallel, startChapter, endChapter int
	flag.IntVar(&parallel, "parallel", 10, "how many parallel tasks")
	flag.IntVar(&startChapter, "start", -1, "from this chapter")
	flag.IntVar(&endChapter, "end", -1, "upto this chapter(inclusive)")
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Printf("Usage: %s <manga id>\n", os.Args[0])
		return
	}

	id := flag.Arg(0)
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
		add := true
		if startChapter != -1 {
			if ch.Index < startChapter {
				add = false
			}
		}

		if endChapter != -1 {
			if ch.Index > endChapter {
				add = false
			}
		}

		if add {
			chapters[key] = ch
		}
	}

	sem := make(chan int, parallel)
	chResult := make(chan ChapterResult, len(chapters))

	// for each chapter
	for k, ch := range chapters {
		go func(k string, ch mangaeden.Chapter) {
			cStart := time.Now()
			// make directory of chapter using chapter.Index
			chDir := filepath.Join(dir, k)

			err = mkdir(chDir)
			if err != nil {
				perror(err)
				return
			}

			// get chapter images
			ims, err := c.Chapter(ch.ID)
			if err != nil {
				perror(err)
				return
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

			ciResult := make(chan ChapterImageResult, len(images))

			for ik, im := range images {
				go func(ik string, im mangaeden.ChapterImage) {
					defer func() { <-sem }()
					imageStart := time.Now()
					sem <- 1
					// download each image and save to chapterimage.Index
					p := filepath.Join(chDir, ik)
					err = download(c, im.Image, p)
					if err != nil {
						perror(err)
					}
					duration := time.Since(imageStart)
					ciResult <- ChapterImageResult{
						Key:       ik,
						Index:     im.Index,
						TimeTaken: duration,
						Err:       err,
					}
				}(ik, im)
			}

			var ok int
			for _, _ = range images {
				cir := <-ciResult
				if cir.Err != nil {
					continue
				}
				ok += 1
			}
			close(ciResult)

			duration := time.Since(cStart)
			chResult <- ChapterResult{
				Key:       k,
				Index:     ch.Index,
				TimeTaken: duration,
				Err:       nil,
			}
			fmt.Printf("Downloaded chapter %d, %d/%d images: %s\n", ch.Index, ok, len(images), duration)
		}(k, ch)
	}

	for _, _ = range chapters {
		<-chResult
	}
	close(chResult)

	close(sem)
	fmt.Printf("Done: %s\n", time.Since(start))
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
